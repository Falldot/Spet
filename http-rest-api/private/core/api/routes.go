package api

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	model "github.com/Survereignty/spet-rest-api/private/core/api/models"
	"github.com/Survereignty/spet-rest-api/private/core/lig"
	"github.com/mholt/archiver/v3"

	"encoding/json"
	"net/http"
)

func (s *Api) Routers() {
	s.router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))))
	// s.router.PathPrefix("/").HandlerFunc(IndexHandler("templates/index.html"))

	s.router.HandleFunc("/getDocs", s.GetDocs()).Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/createDocs", s.CreateDocs()).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/login", s.Login()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/logout", s.Logout()).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/groups", s.Groups()).Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions)
	s.router.HandleFunc("/groupsDel", s.GroupsDel()).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/turnONPC", s.TurnONPC()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/turnOFFPC", s.TurnOFFPC()).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/students",
		s.Students()).Methods(
		http.MethodGet,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPut,
		http.MethodOptions,
	)

	s.router.HandleFunc("/studentsGroup", s.StudentsGroup()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/studentsDel", s.StudentsDel()).Methods(http.MethodPost, http.MethodOptions)
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func (s *Api) setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (s *Api) indexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// t, _ := template.ParseFiles("./public/pages/index.html")
		// t.ExecuteTemplate(w, "index", nil)
		http.ServeFile(w, r, "./public/static/index.html")
	}
}

func (s *Api) TurnONPC() http.HandlerFunc {

	type Data struct {
		Login    string `json:"login"`
		NumGroup string `json:"numGroup"`
		NumRoom  string `json:"numRoom"`
		NumComp  string `json:"numComp"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &Data{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {
			var groups = "/root/http-rest-api/students/groups/" + u.NumGroup + "/" + u.Login
			var shares = "/root/http-rest-api/students/share/" + u.NumRoom + "/" + u.NumComp
			var generals = "/root/http-rest-api/students/generals/" + u.NumGroup

			var to = shares + "/" + u.Login
			var to2 = shares + "/" + u.NumGroup

			os.MkdirAll(groups, os.ModePerm)
			os.MkdirAll(shares, os.ModePerm)
			os.MkdirAll(generals, os.ModePerm)

			err := os.Symlink(groups, to)
			if err != nil {
				lig.Error("Connect symlink", err)
			}
			err = os.Symlink(generals, to2)
			if err != nil {
				lig.Error("Connect symlink", err)
			}
			err = os.Chmod(to, 777)
			if err != nil {
				lig.Error("Chmod", err)
			}
			err = os.Chmod(to2, 444)
			if err != nil {
				lig.Error("Chmod", err)
			}
		}
	}
}

func (s *Api) TurnOFFPC() http.HandlerFunc {

	type Data struct {
		Login    string `json:"login"`
		NumGroup string `json:"numGroup"`
		NumRoom  string `json:"numRoom"`
		NumComp  string `json:"numComp"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &Data{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {
			var groups = "/root/http-rest-api/students/groups/" + u.NumGroup + "/" + u.Login
			var shares = "/root/http-rest-api/students/share/" + u.NumRoom + "/" + u.NumComp
			var generals = "/root/http-rest-api/students/generals/" + u.NumGroup

			var to = shares + "/" + u.Login
			var to2 = shares + "/" + u.NumGroup

			os.MkdirAll(groups, os.ModePerm)
			os.MkdirAll(shares, os.ModePerm)
			os.MkdirAll(generals, os.ModePerm)

			err := os.Remove(to)
			if err != nil {
				lig.Error("Remove symlink", err)
			}
			err = os.Remove(to2)
			if err != nil {
				lig.Error("Remove symlink", err)
			}
		}
	}
}

func (s *Api) CreateDocs() http.HandlerFunc {

	type Data struct {
		File string                 `json:"file"`
		Id   string                 `json:"id"`
		Obj  map[string]interface{} `json:"obj"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		if r.Method == http.MethodPost {

			os.RemoveAll("docs/docx/temp/*")

			u := &Data{}
			json.NewDecoder(r.Body).Decode(u)

			var path_to_docx = "docs/docx/" + u.File + ".docx"
			var path_to_docx_new = "docs/temp/" + u.File + "_new_" + u.Id + ".docx"
			var path_to_json = "docs/json/" + u.File + ".json"
			var path_to_temp = "docs/temp/" + u.File + "_temp" + u.Id
			var path_to_zip = "docs/temp/" + u.File + "_temp_" + u.Id + ".zip"
			var path_to_data = path_to_temp + "/word/document.xml"

			dataJSON, err := ioutil.ReadFile(path_to_json)
			if err != nil {
				lig.Error("ReadFile", err)
			}

			var f interface{}
			err = json.Unmarshal(dataJSON, &f)
			m := f.(map[string]interface{})

			err = Unzip(path_to_docx, path_to_temp)
			if err != nil {
				lig.Error("Unzip", err)
			}

			file, err := os.Open(path_to_data)
			if err != nil {
				lig.Error("Open", err)
			}
			defer file.Close()

			var out string
			data := make([]byte, 64)
			for {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				out += string(data[:n])
			}

			for k, v := range m {
				switch vv := v.(type) {
				case string:
					fmt.Println(u.Obj)
					for q, b := range u.Obj {
						switch vv2 := b.(type) {
						case string:
							//fmt.Println(k, "is string", vv)
							if k == q {
								out = strings.Replace(out, q, vv2, -1)
								fmt.Println(vv2)
								fmt.Println(vv)
							}
						}
					}
				}
			}

			mydata := []byte(out)

			err = ioutil.WriteFile(path_to_data, mydata, 0777)
			if err != nil {
				lig.Error("WriteFile", err)
			}

			files := []string{
				path_to_temp + "/docProps",
				path_to_temp + "/word",
				path_to_temp + "/_rels",
				path_to_temp + "/[Content_Types].xml",
			}
			err = archiver.Archive(files, path_to_zip)
			if err != nil {
				lig.Error("Archive", err)
			}

			err = os.Rename(path_to_zip, path_to_docx_new)
			if err != nil {
				lig.Error("Rename", err)
			}

			s.respond(w, r, http.StatusOK, file)
		}
	}
}
func (s *Api) GetDocs() http.HandlerFunc {

	type Data struct {
		File string `json:"file"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		if r.Method == http.MethodPost {

			u := &Data{}
			json.NewDecoder(r.Body).Decode(u)

			var path_to_json = "docs/json/" + u.File + ".json"

			dataJSON, err := ioutil.ReadFile(path_to_json)
			if err != nil {
				lig.Error("ReadFile", err)
			}

			var data interface{}

			err = json.Unmarshal([]byte(dataJSON), &data)
			if err != nil {
				lig.Error("Unmarshal", err)
			}

			s.respond(w, r, http.StatusOK, data)

		} else if r.Method == http.MethodGet {

			files, err := ioutil.ReadDir("docs/docx/")
			if err != nil {
				lig.Error("ReadDir", err)
			}

			strs := []string{}

			for _, file := range files {
				strs = append(strs, file.Name())
			}

			s.respond(w, r, http.StatusOK, strs)
		}
	}
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)
	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Api) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &model.Group{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {

			s.respond(w, r, http.StatusOK, nil)
		}
	}
}

func (s *Api) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &model.Group{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {

			s.respond(w, r, http.StatusOK, nil)
		}
	}
}

func (s *Api) StudentsGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &model.Group{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {
			um, err := s.store.Student().GetSelectGroup(u.Name)
			if err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusNotFound, err)
				return
			}
			s.respond(w, r, http.StatusOK, um)
		}
	}
}

func (s *Api) StudentsDel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &model.Student{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {
			// Удалить студента
			if err := s.store.Student().Delete(u); err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusOK, u)
		}
	}
}

func (s *Api) GroupsDel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.setHeaders(&w)

		u := &model.Group{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodPost {
			if err := s.store.Group().Delete(u); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusOK, u)
		}
	}
}

func (s *Api) Students() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		s.setHeaders(&w)

		u := &model.Student{}
		json.NewDecoder(r.Body).Decode(u)

		if r.Method == http.MethodGet {
			// Получаем студентов
			um, err := s.store.Student().Get()
			if err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusNotFound, err)
				return
			}
			s.respond(w, r, http.StatusOK, um)
		} else if r.Method == http.MethodPost {
			// Создаем студента
			if err := s.store.Student().Create(u); err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, u)
		} else if r.Method == http.MethodDelete {
			// Удалить студента
			if err := s.store.Student().Delete(u); err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusOK, u)
		} else if r.Method == http.MethodPut {
			// Изменить студента
			if err := s.store.Student().Update(u); err != nil {
				lig.Error("Route", err)
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusOK, u)
		}
	}
}

func (s *Api) Groups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		s.setHeaders(&w)

		u := &model.Group{}
		json.NewDecoder(r.Body).Decode(u)

		// Если метод "Get" то:
		if r.Method == http.MethodGet {
			gs, err := s.store.Group().Get()
			if err != nil {
				s.error(w, r, http.StatusNotFound, err)
				return
			}
			s.respond(w, r, http.StatusOK, gs)
			// Если метод "Post" то:
		} else if r.Method == http.MethodPost {
			if err := s.store.Group().Create(u); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, u)
			// Если метод "Delete" то:
		} else if r.Method == http.MethodDelete {
			if err := s.store.Group().Delete(u); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusOK, u)
		}
	}
}

func (s *Api) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Api) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
