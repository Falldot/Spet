package db

import (
	model "github.com/Survereignty/spet-rest-api/private/core/api/models"
)

type StudentRep struct {
	store *Store
}

// Создать студента
func (r *StudentRep) Create(m *model.Student) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO students (surname, middleName, name, date_b, city, street, house, flat, phone, info, numGroup, activs, gender, status, orphan, invalid, low_income, low_parents, idn, kdn, many_children, login, password, budget) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24) RETURNING id",
		m.Surname, m.MiddleName, m.Name, m.Date_b, m.City, m.Street, m.House,
		m.Flat, m.Phone, m.Info, m.NumGroup, m.Activs, m.Gender, m.Status, m.Orphan, m.Invalid,
		m.Low_income, m.Low_parents, m.Idn, m.Kdn, m.Many_children, m.Login, m.Password, m.Budget,
	).Scan(&m.Id); err != nil {
		return err
	}

	return nil
}

// Изменить студента
func (r *StudentRep) Update(m *model.Student) error {
	_, err := r.store.db.Exec("UPDATE students SET surname = $2, middleName = $3, name = $4, date_b = $5, city = $6, street = $7, house = $8, flat = $9, phone = $10, info = $11, numGroup = $12, activs = $13, gender = $14, status = $15, orphan = $16, invalid = $17, low_income = $18, low_parents = $19, idn = $20, kdn = $21, many_children = $22, login = $23, password = $24, budget = $25 WHERE id = $1",
		m.Id, m.Surname, m.MiddleName, m.Name, m.Date_b, m.City, m.Street, m.House,
		m.Flat, m.Phone, m.Info, m.NumGroup, m.Activs, m.Gender, m.Status, m.Orphan, m.Invalid,
		m.Low_income, m.Low_parents, m.Idn, m.Kdn, m.Many_children, m.Login, m.Password, m.Budget)
	if err != nil {
		return err
	}
	return nil
}

// Удалить
func (r *StudentRep) Delete(m *model.Student) error {
	_, err := r.store.db.Exec(
		"DELETE from students where id = $1",
		m.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRep) GetSelectGroup(name string) ([]*model.Student, error) {

	rows, err := r.store.db.Query("SELECT id, surname, name, middleName, numGroup, login FROM students WHERE numGroup = $1",
		name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sts := make([]*model.Student, 0)
	for rows.Next() {
		st := new(model.Student)
		err := rows.Scan(&st.Id, &st.Surname, &st.MiddleName, &st.Name, &st.NumGroup, &st.Login)
		if err != nil {
			return nil, err
		}
		sts = append(sts, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sts, nil
}

// Получить всех студентов
func (r *StudentRep) Get() ([]*model.Student, error) {

	query := "SELECT * FROM students"

	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sts := make([]*model.Student, 0)
	for rows.Next() {
		st := new(model.Student)
		err := rows.Scan(&st.Id, &st.Surname, &st.MiddleName, &st.Name, &st.Date_b, &st.City, &st.Street, &st.House,
			&st.Flat, &st.Phone, &st.Info, &st.NumGroup, &st.Activs, &st.Gender, &st.Status, &st.Orphan, &st.Invalid,
			&st.Low_income, &st.Low_parents, &st.Idn, &st.Kdn, &st.Many_children, &st.Login, &st.Password, &st.Budget)
		if err != nil {
			return nil, err
		}
		sts = append(sts, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sts, nil
}
