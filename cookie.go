package main

import (
	"net/http"
	"time"
)

func createCookie(w http.ResponseWriter, r *http.Request) {
	cookieName := "session"
	c := &http.Cookie{}

	if c.Value == "" {
		c.Name = cookieName
		c.Value = "qwerty"
		c.Expires = time.Now().Add(2 * time.Minute)
		c.Path = "/"
	}

	http.SetCookie(w, c)

	// process form submission eg login
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
	}

	// save user to db
	dbUsers["2"] = u

	// create session
	dbSessions[c.Value] = "2"
}
