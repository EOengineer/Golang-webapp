package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

func main() {

  // initialize router
  rtr := mux.NewRouter()
  http.Handle("/", rtr)

  // routes
  rtr.HandleFunc("/login", loginHandler)
  rtr.HandleFunc("/logout", logoutHandler)
  rtr.HandleFunc("/services", servicesHandler)
  rtr.HandleFunc("/services/{name}", serviceHandler)
  rtr.HandleFunc("/services/{s_name}/resources", resourcesHandler)
  rtr.HandleFunc("/services/{s_name}/resources/new", newResourceHandler)

  // some server feedback
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}



