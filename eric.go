package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "html/template"
)


func main() {

  // initialize router
  rtr := mux.NewRouter()
  http.Handle("/", rtr)

  // routes
  rtr.HandleFunc("/login", loginHandler)
  rtr.HandleFunc("/services", servicesHandler)
  rtr.HandleFunc("/services/{name}", serviceHandler)
  rtr.HandleFunc("/services/{s_name}/resources", resourcesHandler)
  rtr.HandleFunc("/services/{s_name}/resources/new", newResourceHandler)

  // some server feedback
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

type Service struct {
  Name string
}

type Resource struct {
  Name string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    t, _ := template.ParseFiles("views/login.html")
    t.Execute(w, nil)
  } else {
    // delegate to some method that performs the auth check
   // and conditional redirect on success 
  // authCheck(w, r)
  }
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/services/index.html")
  t.Execute(w, nil)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  s := &Service{Name: vars["name"]}
  t, _ := template.ParseFiles("views/services/show.html")
  t.Execute(w, s)
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  s := &Service{Name: vars["s_name"]}
  t, _ := template.ParseFiles("views/resources/index.html")
  t.Execute(w, s)
}

func newResourceHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/resources/new.html")
  t.Execute(w, nil)
}
