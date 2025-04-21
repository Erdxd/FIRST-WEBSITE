package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                float64
	Avg_grades, Happines float64
	Hobbies              []string
}

func (u User) getAllInfo() string {

	return fmt.Sprintf("User name is: %s. He is %d and he has money"+
		" equal: %.2f ", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {

	bob := User{Name: "Alexsey", Age: 15, Money: 0, Avg_grades: 4.4, Happines: 0.5, Hobbies: []string{"Programming", "play computer games"}}
	//bob.setNewName("Alex")
	//mt.Fprintf(page, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, bob)

}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, " contacts page")

}
func HandleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}
func main() {

	HandleRequest()

}
