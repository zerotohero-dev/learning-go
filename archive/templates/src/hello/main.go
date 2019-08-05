package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func (User) IsNinja() bool {
	return true
}

func readFileTwo() {
	u := User{"Volkan", "Ozcelik", "me@volkan.io", 16}
	body, err := ioutil.ReadFile("data/templates/file2.tmpl")

	fmt.Println(err)

	tmpl, err := template.New("awesometpl").Parse(string(body))

	if err != nil {
		log.Panic(err)
	}

	tmpl.Execute(os.Stdout, u)
}

func readFileThree() {
	u := User{"Volkan", "Ozcelik", "me@volkan.io", 16}

	tmpl, _ = template.ParseFiles("data/templates/file2.tmpl", "data/templates/file3.tmpl")

	err := tmpl.ExecuteTemplate(os.Stdout, "file2.tmpl", u)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	u := User{"Volkan", "Ozcelik", "me@volkan.io", 16}
	body, err := ioutil.ReadFile("data/templates/file1.tmpl")

	fmt.Println(err)

	tmpl, err := template.New("awesometpl").Parse(string(body))

	if err != nil {
		log.Panic(err)
	}

	tmpl.Execute(os.Stdout, u)

	readFileTwo()
	readFileThree()
}
