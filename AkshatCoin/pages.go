package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/login.html") //Loads search template
	if err != nil {
		panic(err)
	}
	r.ParseForm() //get user inputs
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	if len(username+password) > 0 {
		reference, ok := accounts[username][0].(string)
		if ok && reference == password {
			setCookie(w, "username", username)
			http.Redirect(w, r, "/"+username, http.StatusFound)
		} else {
			http.Redirect(w, r, "/wrong_login", http.StatusFound)
		}
	}
	t.Execute(w, "a")
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}

type Page struct {
	Name  string
	Coins int
	Level string
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/main.html")
	if err != nil {
		panic(err)
	}
	uname, err := r.Cookie("username")
	if err != nil {
		panic(err)
	}
	username := uname.Value
	data := accounts[username]
	t.Execute(w, &Page{
		username,
		data[1].(int),
		data[2].(string),
	})
}
