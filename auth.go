package main

import (
  "net/http"
  "html/template"
)

const (
  AUTH_USER = "User"
  AUTH_PW = "Password1"
)

var (
  auth_success = false
)


// render login if get request, otherwise hand off to perform authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    t, _ := template.ParseFiles("views/login.html")
    t.Execute(w, nil)
  } else {
    authCheck(w, r)
  }
}

// logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
  auth_success = false
  enforcer(w, r, auth_success)
}

// eventually modify this method to call out to authentication service
func authCheck(w http.ResponseWriter, r *http.Request) {
  user := r.FormValue("username")
  pw := r.FormValue("password")
  if (pw == AUTH_PW) && (user == AUTH_USER) {
    auth_success = true
    http.Redirect(w, r, "/services", http.StatusFound)
  } else {
    enforcer(w, r, auth_success)
  }
}

// used to restrict access to any function that includes it
func enforcer(w http.ResponseWriter, r *http.Request, authed bool) {
  if authed == false {
    http.Redirect(w, r, "/login", http.StatusFound)
  }
}
