package lig

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Log in go

var (
	path string
	file *os.File
)

// Create создает файл в который будут записываться логи
func Create(p string) error {
	path = p
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("lig -> %v", err)
	}
	file = f
	log.SetOutput(f)
	log.Println("|------------- Server starting ----------------|")
	return nil
}

// Quit ...
func Quit() {
	file.Close()
}

// Get выдает все логи
func Get() {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		Error("Not read", err)
	}
	fmt.Print(string(b))
}

// Info ...
func Info(s string) {
	log.Println("[INFO]", s)
	fmt.Println("[INFO]", s)
}

func Listen(s string, port string) {
	log.Println("[START]", s, "->", port)
	fmt.Println("[START]", s, "->", port)
}

// Error выдает данные в консоль и записывает лог, возвращая значение ошибки
func Error(s string, e error) error {
	log.Println("[ERROR]", s, "->", e)
	fmt.Println("[ERROR]", s, "->", e)
	return fmt.Errorf("[ERROR]", s, "->", e)
}

// Сrash выдает все данные и логи в консоль и записывает лог
func Сrash(s string, e error) {
	fmt.Println("[CRASH]", s, "->", e)
	log.Fatal("[CRASH]", s, "->", e)
}
