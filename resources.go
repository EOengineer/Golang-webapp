package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "html/template"
)

// TODO - figure out what this data structure should look like
type Resource struct {
  Name string
}

// resource index
func resourcesHandler(w http.ResponseWriter, r *http.Request) {
  enforcer(w, r, auth_success)
  vars := mux.Vars(r)
  s := &Service{Name: vars["s_name"]}
  t, _ := template.ParseFiles("views/resources/index.html")
  t.Execute(w, s)
}

// resource new
func newResourceHandler(w http.ResponseWriter, r *http.Request) {
  enforcer(w, r, auth_success)
  t, _ := template.ParseFiles("views/resources/new.html")
  t.Execute(w, nil)
}
