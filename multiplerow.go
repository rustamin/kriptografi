package main

import (

  "net/http"
  "log"
  "fmt"

)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB
var err error


func multi(res http.ResponseWriter, req *http.Request) {


  var (
    alamat string
    nama string
  )

  rows, err := db.Query("select nama, alamat from coba where nama = ?", "rustam")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  fmt.Println(rows)

  var cookies []string


  for rows.Next() {
    err := rows.Scan(&nama, &alamat)
    if err != nil {
      log.Fatal(err)
    }
    // log.Println(nama, alamat)
    // res.Write([]byte("Hello " + nama))

    // push to slice cookie
    cookies = append(cookies, alamat)
  }
  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(cookies)

  // search cookie
  for i := range cookies {
    if cookies[i] == "alamat1" {
      // Found!
      fmt.Println("cookie ditemukan")
      break
    }
  }


  // res.Write([]byte("Hello" + databaseUsername))


}


func main() {

  db, err = sql.Open("mysql", "root:@/golang-example")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  http.HandleFunc("/", multi)


  fmt.Println("listen on 3000")
  err := http.ListenAndServe(":3000", nil)

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
