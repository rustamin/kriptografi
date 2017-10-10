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
  // io.WriteString(w, "<a href='/readcookie'>Read Cookie <ba><br/>" + "\n")
  // io.WriteString(w, "<a href='/writecookie'>Write Cookie <ba><br/>" + "\n")
  // io.WriteString(w, "<a href='/deletecookie'>Delete Cookie <ba><br/>" + "\n")
  io.WriteString(w, "<a href='/login'>Login <ba><br/>" + "\n")
  io.WriteString(w, "<a href='/register'>Register <ba><br/>" + "\n")

}

// set session in login ketika ada cookie
// di welcome ada button logout

// register page
// alurnya
// the func register

// testing algo
// deploy

func IndexServer(w http.ResponseWriter, req *http.Request) {
  // draw menu
  DrawMenu(w)
}


func ReadCookieServer(w http.ResponseWriter, req *http.Request) {

  // draw menu
  DrawMenu(w)

  // read cookie
  var cookie,err = req.Cookie("testcookiename")
  if err == nil {
    var cookievalue = cookie.Value
    io.WriteString(w, "<b>get cookie value is " + cookievalue + "</b>\n")
  }

}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
  // set cookies.
  expire := time.Now().AddDate(0, 1, 0)
  cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire/*, MaxAge: 86400 *  30*/}

  http.SetCookie(w, &cookie)

  //
  // we can not set cookie after writing something to ResponseWriter
  // if so ,we cannot set cookie succefully.
  //
  // so we have draw menu after set cookie
  DrawMenu(w)

}


func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {

  // set cookies.
  cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
  http.SetCookie(w, &cookie)

  // ABOUT MaxAge
  // MaxAge=0 means no 'Max-Age' attribute specified.
  // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
  // MaxAge>0 means Max-Age attribute present and given in seconds

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

  // password tidak sama
  if password != password2 {
    fmt.Println("password tidak sama")
    http.Redirect(res, req, "/register", 301)
		return
  }

	var databaseUsername string

  // cek username
	errQuery := db.QueryRow("SELECT  username FROM users WHERE username=?", username).Scan(&databaseUsername)

  // ada user di database, kembali ke login
	if errQuery == nil {
    fmt.Println("ada user di database, kembali ke login")
		http.Redirect(res, req, "/register", 301)
		return
	}

  // enkrip password
  fmt.Println("password awal: ", password)
  // scytale
  password  = kriptoalgo.Scytale(password)
  fmt.Println("scytale: ",password)
  // transposisi
  password  = kriptoalgo.Transposisi(password)
  fmt.Println("transposisi: ",password)
  // // caesar
  password  = kriptoalgo.Caesar(password)
  fmt.Println("Caesar: ",password)
  // return


  // insert user ke database
  _, err = db.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", username, email, password)
  if err != nil { http.Error(res, "Server error, unable to create your account.", 500); return; }

  http.Redirect(res, req, "/login", 301)
  return



}

// res http.ResponseWriter, req *http.Request
func login(res http.ResponseWriter, req *http.Request) {


  if req.Method != "POST" {

    http.ServeFile(res, req, "login.html")
    return
  }


  // userName := sess.CAttr("UserName")
  // fmt.Println(userName)
  //
  // return



  username := req.FormValue("username")
	password := req.FormValue("password")

  var databaseUserId int
	var databaseUsername string
  var databaseEmail string
	var databasePassword string

  // cek username
	errQuery := db.QueryRow("SELECT id, email, username, password FROM users WHERE username=?", username).Scan(&databaseUserId, &databaseEmail, &databaseUsername, &databasePassword)

	if errQuery != nil {
    fmt.Println("ga ada user di database")
		http.Redirect(res, req, "/login", 301)
		return
	}

  // cek password

  // scytale
  password  = kriptoalgo.Scytale(password)
  fmt.Println("scytale: ",password)
  // transposisi
  password  = kriptoalgo.Transposisi(password)
  fmt.Println("transposisi: ",password)
  // // caesar
  password  = kriptoalgo.Caesar(password)
  fmt.Println("Caesar: ",password)

  if password != databasePassword {
    http.Redirect(res, req, "/login", 301)
		return
  }

  // cek cookie dari client
  // jika tidak ada cookie langsung ke validasi page, generate kode, kirim email, insert ke database kodenya
  var cookie,err = req.Cookie("testcookiename")
  // tidak ada cookie
  if err != nil {
    fmt.Println("tidak ada cookie di client")

    createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
    return
  }

  // ada cookie di client
  // select cookies di database berdasarkan user id, value cookie-nya dan statusnya aktif
  // jika ada set session, masuk ke index
  // jika tidak ada createVerificationCode

  // INI SEBELUM REVISI
  // get all cookie related to user in db and active.
  // cocokan cookies in db and client
  // jika cocok, masuk ke index
  //jika tidak cocok ke validasi page, generate kode, kirim email, insert ke db kodenya
  fmt.Println("ada cookie")
  fmt.Println("cookienya yaitu " + cookie.Value)

  var cnt int

  errQuery = db.QueryRow("SELECT COUNT(*) FROM cookies WHERE active=1 and user_id=? and cookie_value=?", databaseUserId, cookie.Value).Scan(&cnt)

  if errQuery != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

  fmt.Println("banyaknya cookie adalah ", cnt)

  if cnt == 0 {
    fmt.Println("cookie dari client tidak cocok dengan di db ")
    createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
		return
  } else {
    fmt.Println("cookie dari client cocok dengan di db ")
    // HARUSNYA SET SESSION
    // set SESSION
    // // Get a session. Get() always returns a session, even if empty.
    // session, err := store.Get(req, "session-name")
    // if err != nil {
    //    http.Error(res, err.Error(), http.StatusInternalServerError)
    //    return
    // }
    //
    // // Set some session values.
    // session.Values["UserName"] = databaseUsername
    // session.Values["authenticated"] = true
    // // Save it before we write to the response/return from the handler.
    // session.Save(req, res)
    //
    // fmt.Println("cek session dulu yaaa")
    // // fmt.Println(session.Values["foo"])
    //
    // if session != nil {
    //   fmt.Println("ada session bung fdsjfdlf")
    //   http.Redirect(res, req, "/welcome", 301)
    //   return
    // } else {
    //   fmt.Println("gagal set session")
    //   http.Redirect(res, req, "/login", 301)
    //   return
    // }


    // not working
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

    // http.Redirect(res, req, "/welcome", 301)
    // return
  }




  // rows, err := db.Query("SELECT code FROM cookies WHERE active=1 and user_id=? and cookie_value=?", databaseUserId, cookie.Value)
  // if err != nil {
  //   log.Fatal(err)
  // }
  // defer rows.Close()
  //
  //
  //
  // var dbCookies []string
  // var code string
  //
  // for rows.Next() {
  //   err := rows.Scan(&code)
  //   if err != nil {
  //     log.Fatal(err)
  //     fmt.Println(err)
  //   }
  //   // push to slice cookie
  //   dbCookies = append(dbCookies, code)
  // }
  // err = rows.Err()
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // // cek cookie di db ada ga
  // if len(dbCookies) == 0 { // ga ada di db
  //   fmt.Println("ga ada di db")
  //
  //   createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
  //   return
  //
  // } else {
  //   // ada cookies di db, cocokan dengan coockie client
  //   fmt.Println("ada di db")
  //
  //   // search cookie
  //   for i := range dbCookies {
  //     if dbCookies[i] == cookie.Value {
  //       // Found!
  //       fmt.Println("cookie cocok antara db dan dari client, boleh masuk ke index")
  //       // HARUSNYA SET SESSION
  //       http.Redirect(res, req, "/index", 301)
  //       break
  //     }
  //   }
  //
  //   // cookie ga cocok andara db dan client
  //   fmt.Println("cookie ga cocok antara db dan client")
  //
  //   createVerificationCode(res, req, databaseUserId, databaseEmail) // create kode, insert ke db, kirim ke email
  //   return
  // }
  //
  // /* jika ada cookie
  //       select all cookie in database related to the user
  //       cocokkan dengan cookies di db dan di browser user
  //       jika ada masuk ke index page
  //       jika tidak ada yg cocok masuk ke validasi page, generate kode, kirim email
  // */


}




func verification(res http.ResponseWriter, req *http.Request) {
  sess := session.Get(req)
  if sess != nil {
    fmt.Println("ada session")
    http.Redirect(res, req, "/welcome", 301)
    return
  }
  // else {
  //   fmt.Println("ga ada session")
  //   http.Redirect(res, req, "/login", 301)
  //   return
  // }

  if req.Method != "POST" {
    http.ServeFile(res, req, "verification.html")
    return
  }

  /*  method POST
      kode dari client
      get all cookie code di db based on kode verifikasi dan status not active
      kode dari client cocok dengan cookie code di db
      jika ya,
        generate cookie
        update status cookie dan tambahakn generate cookie  menjadi aktif,
        set session dan masuk ke index
      jika return halaman verification dengan pesan error
  */

  codeClient := req.FormValue("code")

  var cnt int

  errQuery := db.QueryRow("SELECT COUNT(*) FROM cookies WHERE active=0 and code=?", codeClient).Scan(&cnt)

  if errQuery != nil {
		http.Redirect(res, req, "/verification", 301)
		return
	}

  if cnt == 0 {
    fmt.Println("kode verifikasi tidak sesuai dengan yang di database ")
    // harusnya dengan pesan error
    http.Redirect(res, req, "/verification", 301)
		return
  } else {
    fmt.Println("kode verifikasi sesuai dengan yang di database ")
    // LIAT FLOWCHART
    // UPDATE STATUS CREATE COOKIE STORE IN DB
    // HARUSNYA SET SESSION

    // buat cookie
    expire := time.Now().AddDate(0, 1, 0)
    cookievalue := randstring.RandomStringForCookie(25)
    cookie := http.Cookie{Name: "testcookiename", Value: cookievalue, Path: "/", Expires: expire/*, MaxAge: 86400*/}

    http.SetCookie(res, &cookie)

    // insert kedatabase cookie value nya
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
    fmt.Println(userName)

    if sess != nil {
        // No session (yet)
        // ada session
        fmt.Println("ada session")
        http.Redirect(res, req, "/welcome", 301)
        return
    }

  }
}

func welcome(res http.ResponseWriter, req *http.Request) {

  // // Get a session. Get() always returns a session, even if empty.
  // session, err := store.Get(req, "session-name")
  // if err != nil {
  //    http.Error(res, err.Error(), http.StatusInternalServerError)
  //    return
  // }
  //
  //
  // fmt.Println("cek session dulu yaaa di welcome")
  // // fmt.Println(session.Values["foo"])
  //
  // if session.Values["authenticated"] == true {
  //   fmt.Println("ada session bung fdsjfdlf")
  //   http.ServeFile(res, req, "welcome.html")
  //   return
  // } else {
  //   fmt.Println("gagal set session")
  //   http.Redirect(res, req, "/login", 301)
  //   return
  // }

  sess := session.Get(req)
  // userName := sess.CAttr("UserName")
  // fmt.Println(userName)
  //
  // return

  if sess == nil {
      // No session (yet)
      http.Redirect(res, req, "/login", 301)
      return
  } else {
    // cek session dulu harusnya
    http.ServeFile(res, req, "welcome.html")
    userName := sess.CAttr("UserName")
    fmt.Println(userName)
    return
  }


}

func dfsfsdfdsfsd(res http.ResponseWriter, req *http.Request) {
  // session, err := store.Get(req, "session-name")
  //
  // if req.Method != "POST" {
  //
  //   // Get a session. Get() always returns a session, even if empty.
  //
  //   if err != nil {
  //      http.Error(res, err.Error(), http.StatusInternalServerError)
  //      return
  //   }
  //
  //
  //   fmt.Println("cek session dulu yaaa di logout")
  //   // fmt.Println(session.Values["foo"])
  //
  //   session.Values["authenticated"] = false
  //   session.Values["UserName"] = ""
  //   session.Save(req, res)
  //
  //   if session.Values["authenticated"] == false {
  //     fmt.Println("udah logout bung masuk ke login")
  //     http.Redirect(res, req, "/login", 301)
  //     return
  //   } else {
  //     fmt.Println("belum logout")
  //     http.Redirect(res, req, "/welcome", 301)
  //     return
  //   }
  // }
  //




  sess := session.Get(req)
  fmt.Println("akfdsjldjfs")
  log.Println("Session:", sess)

  if sess == nil {
      // tidak ada session
      fmt.Println("tidak ada session bung towel")
  } else {
    fmt.Println("ada session bung towel")
  }

  userName := sess.CAttr("UserName")
  fmt.Println(userName)

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
  http.HandleFunc("/readcookie", ReadCookieServer)
  http.HandleFunc("/writecookie", WriteCookieServer)
  http.HandleFunc("/deletecookie", DeleteCookieServer)

  http.HandleFunc("/welcome", welcome)
  http.HandleFunc("/login", login)
  http.HandleFunc("/register", register)
  http.HandleFunc("/verification", verification)
  http.HandleFunc("/logout", dfsfsdfdsfsd)



  fmt.Println("listen on 3000")
  err := http.ListenAndServe(":3000", context.ClearHandler(http.DefaultServeMux))

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}


func createVerificationCode(res http.ResponseWriter, req *http.Request, databaseUserId int, databaseEmail string) {
  // generate kode yang akan dikirim lewat email
  var randString string = randstring.RandomStringForVerify(6)
  fmt.Println(randString)
  // insert kedatabase kodenya
  _, err = db.Exec("INSERT INTO cookies(user_id, code) VALUES(?, ?)", databaseUserId, randString)
  if err != nil { http.Error(res, "Server error, unable to create your account.", 500); return; }
  // kirim email
  email.SendMail(databaseEmail, randString)
  // redirect ke vilidasi page
  http.Redirect(res, req, "/verification", 301)

}


func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
