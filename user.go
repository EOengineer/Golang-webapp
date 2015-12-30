package main

import (
  "fmt"
  "time"
  "log"
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/gorilla/mux"
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
  enforcer(resp, req, authenticated)

  // get the query string parameters
  vars := mux.Vars(req)
  // get the id
  uid := vars["uid"]


  // open the database connection
  db, err := sql.Open("postgres", db_config)
  if err != nil {
    log.Fatal(err)
  }

  // interpolate and run the query
  row := db.QueryRow("SELECT * FROM users WHERE Uid = $1", uid)
  if err != nil {
    log.Fatal(err)
  }

  // initialize the struct
  RetreivedUser := &User{}


  // coerse database results to user struct
  err = row.Scan(&RetreivedUser.Uid, &RetreivedUser.Username, &RetreivedUser.Created)
  if err != nil {
    log.Fatal(err)
  }

  t, _ := template.ParseFiles("views/users/show.html")
  t.Execute(resp, &RetreivedUser)


}

