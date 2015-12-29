package main

import (
  "fmt"
  "time"
  "log"
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/lib/pq"
)

const (
  db_config = "host=localhost dbname=testdb sslmode=disable"
)



type User struct {
  Uid       int64
  Username  string
  Created   time.Time
}



type Users []User



// TODO **** Share a db connection pool, get rid of all these opens/closes

func usersHandler(resp http.ResponseWriter, req *http.Request) {
  enforcer(resp, req, authenticated)

  db, err := sql.Open("postgres", db_config)
  if err != nil {
    log.Fatal(err)
  }

  rows, err := db.Query("SELECT * FROM users;")
  if err != nil {
    log.Fatal(err)
  }

  // replace with code to output this structure as html template
  for rows.Next() {
    var uid int
    var username string
    var created time.Time
    err = rows.Scan(&uid, &username, &created)
    fmt.Println("uid | username | created ")
    fmt.Printf("%3v | %8v | %6v\n", uid, username, created)

  }

  t, _ := template.ParseFiles("views/users/index.html")
  t.Execute(resp, nil)
}


func userHandler(resp http.ResponseWriter, req *http.Request) {
  enforcer(w, r, authenticated)

  vars := mux.Vars(r)
  // get the id
  uid := vars["uid"]

  db, err := sql.Open("postgres", db_config)
  if err != nil {
    log.Fatal(err)
  }

  // interpolate and run the query
  rows, err := db.Query("SELECT * FROM users WHERE uid = $1", uid)
  if err != nil {
    log.Fatal(err)
  }
}

// some code to insert the data into the view

// func serviceHandler(w http.ResponseWriter, r *http.Request) {
//   enforcer(w, r, authenticated)
//   vars := mux.Vars(r)
//   s := &Service{Name: vars["name"]}
//   t, _ := template.ParseFiles("views/services/show.html")
//   t.Execute(w, s)
// }
