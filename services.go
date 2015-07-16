package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "html/template"
)

// TODO figure out what this data structure should look like
type Service struct {
  Name string
}

// service index
func servicesHandler(w http.ResponseWriter, r *http.Request) {
  enforcer(w, r, auth_success)
  t, _ := template.ParseFiles("views/services/index.html")
  t.Execute(w, nil)
}

// service show
func serviceHandler(w http.ResponseWriter, r *http.Request) {
  enforcer(w, r, auth_success)
  vars := mux.Vars(r)
  s := &Service{Name: vars["name"]}
  t, _ := template.ParseFiles("views/services/show.html")
  t.Execute(w, s)
}
