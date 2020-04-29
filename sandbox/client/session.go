package main

import (
	"fmt"
	"net/http"
)

// Dummy user and db
type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var dbUsers = map[string]user{}      // userID , user
var dbSessions = map[string]string{} // sessionID, userID
var u user

func createSession(w http.ResponseWriter, r *http.Request) {
	// get cookie if no cookies set cookies and assign session
	c, err := r.Cookie("session")
	if err != nil {
		createCookie(w, r)
		fmt.Println(c)
	}

	// get session
	// dummy user and session
	dbUsers["1"] = user{
		UserName:  "sat",
		FirstName: "satriyo",
		LastName:  "pamungkas",
	}

	dbSessions["asdf"] = "1"

	// getSession(c)

	fmt.Println(u)
	w.Write([]byte(u.UserName))
}

func getSession(c *http.Cookie) (user, error) {
	// if user exist, get the user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u, nil
}

func tesSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte(un))
}
