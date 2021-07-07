package main

import (
  "fmt"
  "net/http"
  //"html/template"
)

type User struct {
  name string
  age uint16
  money int16
  avg_grades, happiness float64
}

func (u *User) getAllInfo() string{
  return fmt.Sprintf("User name is: %s. Age: %d. Money: %d.", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
  man := User {"Fedya", 20, 20, 3.8, 0.6}
  man.setNewName("Vova")
  fmt.Fprintf(w, man.getAllInfo() )
}

func about_page(w http.ResponseWriter, r *http.Request) {
  //tmpl, _ := template.ParseFiles("Портал услуг Уссурийской Федерации.html")
  //tmpl.Execute(w, 0)
}

func webRequest() {

  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)
  http.ListenAndServe(":81", nil)
  //test push
}

func main() {
  webRequest()
}
