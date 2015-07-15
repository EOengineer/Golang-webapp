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
  rtr.HandleFunc("/services", serviceHandler)
  rtr.HandleFunc("/services/{key}/resources", resourceHandler)

  // some server feedback
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}


func serviceHandler(w http.ResponseWriter, r *http.Request) {
  // params := mux.Vars(r)
  p := "whatever"
  // key := params["key"]
  t, _ := template.ParseFiles("services.html")
  t.Execute(w, p)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
  p := "Whatever"
  t, _ := template.ParseFiles("resources.html")
  t.Execute(w, p)
}
