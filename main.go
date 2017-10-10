package main

import (
  "io"
  "net/http"
  "log"
  "fmt"
  "time"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/icza/session"
import "github.com/rustamin/kriptografi/randstring"
import "github.com/rustamin/kriptografi/email"
import "github.com/rustamin/kriptografi/kriptoalgo"
import "github.com/gorilla/sessions"
import "github.com/gorilla/context"

var db *sql.DB
var err error

var store = sessions.NewCookieStore([]byte("something-very-secret-rahasia"))

func DrawMenu(w http.ResponseWriter){
  w.Header().Set("Content-Type", "text/html")
  io.WriteString(w, "<a href='/'>HOME <ba><br/>" + "\n")
  io.WriteString(w, "<a href='/login'>Login <ba><br/>" + "\n")
  io.WriteString(w, "<a href='/register'>Register <ba><br/>" + "\n")

}

/* TODO
    deploy on heroku
*/

func IndexServer(w http.ResponseWriter, req *http.Request) {
  // draw menu
  DrawMenu(w)
}


func register(res http.ResponseWriter, req *http.Request) {
  if req.Method != "POST" {
    http.ServeFile(res, req, "register.html")
    return
  }

  username := req.FormValue("username")
  email := req.FormValue("email")
  password := req.FormValue("password")
  password2 := req.FormValue("password")

  // password doesn't match
  if password != password2 {
    http.Redirect(res, req, "/register", 301)
		return
  }

	var databaseUsername string

  // cek username
	errQuery := db.QueryRow("SELECT  username FROM users WHERE username=?", username).Scan(&databaseUsername)

  // username already taken
	if errQuery == nil {
    fmt.Println("ada user di database, kembali ke login")
		http.Redirect(res, req, "/register", 301)
		return
	}

  // encrypt password
  fmt.Println("original password: ", password)
  // scytale
  password  = kriptoalgo.Scytale(password)
  fmt.Println("scytale: ",password)
  // transposition
  password  = kriptoalgo.Transposisi(password)
  fmt.Println("transposition: ",password)
  // // caesar
  password  = kriptoalgo.Caesar(password)
  fmt.Println("Caesar: ",password)

  // insert user ke database
  _, err = db.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", username, email, password)
  if err != nil { http.Error(res, "Server error, unable to create your account.", 500); return; }

  http.Redirect(res, req, "/login", 301)
  return
}

func login(res http.ResponseWriter, req *http.Request) {

  if req.Method != "POST" {
    http.ServeFile(res, req, "login.html")
    return
  }

  username := req.FormValue("username")
	password := req.FormValue("password")

  var databaseUserId int
	var databaseUsername string
  var databaseEmail string
	var databasePassword string

  // check username in databse
	errQuery := db.QueryRow("SELECT id, email, username, password FROM users WHERE username=?", username).Scan(&databaseUserId, &databaseEmail, &databaseUsername, &databasePassword)

  // there is no user in db
	if errQuery != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

  // check password
  // scytale
  password  = kriptoalgo.Scytale(password)
  fmt.Println("scytale: ",password)
  // transposition
  password  = kriptoalgo.Transposisi(password)
  fmt.Println("transposition: ",password)
  // caesar
  password  = kriptoalgo.Caesar(password)
  fmt.Println("Caesar: ",password)

  if password != databasePassword {
    http.Redirect(res, req, "/login", 301)
		return
  }

  // check cookie in client
  // if there is no cookie, go to validation page, generate code, email and insert ot db the code
  var cookie,err = req.Cookie("testcookiename")
  // tidak ada cookie
  if err != nil {
    createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
    return
  }

  // there is a cookie in client
  // select cookies in database based on user id, value cookie-nya and status
  // if cookie in client match with cookie in db, set session, go to index page
  // else createVerificationCode
  var cnt int

  errQuery = db.QueryRow("SELECT COUNT(*) FROM cookies WHERE active=1 and user_id=? and cookie_value=?", databaseUserId, cookie.Value).Scan(&cnt)

  if errQuery != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

  // cookies does not match
  if cnt == 0 {
    createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
		return
  } else { // cookies match

    sess := session.NewSessionOptions(&session.SessOptions{
        CAttrs: map[string]interface{}{"UserName": databaseUsername},
    })
    session.Add(sess, res)
    userName := sess.CAttr("UserName")
    fmt.Println(userName)

    if sess != nil {
        // No session (yet)
        // ada session
        fmt.Println("ada session")
        http.Redirect(res, req, "/welcome", 301)
        return
    } else {
      fmt.Println("gagal set session")
      http.Redirect(res, req, "/login", 301)
      return
    }
  }
}


func verification(res http.ResponseWriter, req *http.Request) {

  sess := session.Get(req)
  if sess != nil {
    http.Redirect(res, req, "/welcome", 301)
    return
  }

  if req.Method != "POST" {
    http.ServeFile(res, req, "verification.html")
    return
  }

  /*  method POST
      verification code from client
      get all cookie code in db based on verfication code and status not active
      if code verfication match
        generate cookie
        update status cookie dan add cookie value, status active in db
        set session and go to index page
      else return to verification page with error message (error message do not finish yet)
  */

  codeClient := req.FormValue("code")

  var cnt int

  errQuery := db.QueryRow("SELECT COUNT(*) FROM cookies WHERE active=0 and code=?", codeClient).Scan(&cnt)

  if errQuery != nil {
		http.Redirect(res, req, "/verification", 301)
		return
	}

  // verification code does not match
  if cnt == 0 {
    http.Redirect(res, req, "/verification", 301)
		return
  } else { // match verification code

    // create cookie
    expire := time.Now().AddDate(0, 1, 0)
    cookievalue := randstring.RandomStringForCookie(25)
    cookie := http.Cookie{Name: "testcookiename", Value: cookievalue, Path: "/", Expires: expire/*, MaxAge: 86400*/}

    http.SetCookie(res, &cookie) // set cookie

    // insert  cookie value to db
    stmt, err := db.Prepare("UPDATE cookies set active = ?, cookie_value = ? where code=?")
    checkErr(err)
    _, err = stmt.Exec(1, cookievalue, codeClient)
    checkErr(err)

    // get username
    var databaseUsername string
    errQuery := db.QueryRow("SELECT username FROM users WHERE id in (select user_id from cookies where code = ? and cookie_value = ?)", codeClient, cookievalue).Scan(&databaseUsername)
    checkErr(errQuery)

    // set session
    sess := session.NewSessionOptions(&session.SessOptions{
        CAttrs: map[string]interface{}{"UserName": databaseUsername},
    })
    session.Add(sess, res)
    userName := sess.CAttr("UserName")

    if sess != nil {
        // session already set
        http.Redirect(res, req, "/welcome", 301)
        return
    }
  }
}

func welcome(res http.ResponseWriter, req *http.Request) {

  sess := session.Get(req)

  if sess == nil {
      // No session (yet)
      http.Redirect(res, req, "/login", 301)
      return
  } else {
    http.ServeFile(res, req, "welcome.html")
    userName := sess.CAttr("UserName")
    return
  }
}

/*
TODO fix logout/destroy session
logout not working properly, icza/session library failed to delete the session when logout
*/

func logout(res http.ResponseWriter, req *http.Request) {
  sess := session.Get(req)

  session.Remove(sess, res)
  sess = nil

  return
  http.Redirect(res, req, "/login", 301)
}



func main() {
  session.Global.Close()
	session.Global = session.NewCookieManagerOptions(session.NewInMemStore(), &session.CookieMngrOptions{AllowHTTP: true})

  db, err = sql.Open("mysql", "root:@/golang-example")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  http.HandleFunc("/", IndexServer)

  http.HandleFunc("/welcome", welcome)
  http.HandleFunc("/login", login)
  http.HandleFunc("/register", register)
  http.HandleFunc("/verification", verification)
  http.HandleFunc("/logout", logout)

  fmt.Println("listen on 3000")
  err := http.ListenAndServe(":3000", context.ClearHandler(http.DefaultServeMux))

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}


func createVerificationCode(res http.ResponseWriter, req *http.Request, databaseUserId int, databaseEmail string) {
  // generate verification code
  var randString string = randstring.RandomStringForVerify(6)
  fmt.Println(randString)
  // insert the code to db
  _, err = db.Exec("INSERT INTO cookies(user_id, code) VALUES(?, ?)", databaseUserId, randString)
  if err != nil { http.Error(res, "Server error, unable to create your account.", 500); return; }
  // send email the code
  email.SendMail(databaseEmail, randString)

  http.Redirect(res, req, "/verification", 301)
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
