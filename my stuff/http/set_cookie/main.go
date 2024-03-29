package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/setcookie", setcookie)
	log.Println("server running http://192.168.1.155:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	sessionId, _ := r.Cookie("session-id")
	if sessionId != nil {
		isLoggedIn := "logged in"
		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>` + isLoggedIn + `</h1>
	<div>
  <form method="POST" action="/setcookie">     
      <label>name</label><input name="name" type="text" value="" />
      <input type="submit" value="submit" />
  </form>
</div>
</body>
</html>
`
		log.Println("cookie found")
		w.Write([]byte(html))

	} else {
		isLoggedIn := "not logged in"
		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>` + isLoggedIn + `</h1>
	<div>
  <form method="POST" action="/setcookie">     
      <label>name</label><input name="name" type="text" value="" />
      <input type="submit" value="submit" />
  </form>
</div>
</body>
</html>
`
		log.Println("cookie not found")
		w.Write([]byte(html))
	}
}

func setcookie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("POST /setcookie")
		c := &http.Cookie{
			Name:     "session-id",
			Value:    "MTIzNDY1NDY1NDY1NDY1NDY1NDEzMjE2",
			MaxAge:   300,
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
