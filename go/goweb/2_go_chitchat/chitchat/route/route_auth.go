package route

import (
	"crypto/md5"
	"fmt"
	"goweb/2_go_chitchat/chitchat/data"
	"goweb/2_go_chitchat/chitchat/utilities"
	"io"
	"net/http"
	"strconv"
	"time"
)

// GET /err?msg=
// Show the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := utilities.Session(writer, request)
	if err != nil {
		utilities.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utilities.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	threads, err := data.Threads()
	if err != nil {
		utilities.RespErrMsg(writer, request, "Cannot get threads")
		utilities.Warning(err)
	} else {
		m := make(map[string][]data.Thread)
		m["thread"] = threads
		m["hotspot"] = data.GetHotspotThreads()
		_, err = utilities.Session(writer, request)
		if err != nil {
			utilities.GenerateHTML(writer, m, "layout", "public.navbar", "index")
		} else {
			utilities.GenerateHTML(writer, m, "layout", "private.navbar", "index")
		}
	}
}

// GET /login
// show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := utilities.ParseTemplateFiles("login.layout", "public.navbar", "login")
	currTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(currTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	t.ExecuteTemplate(writer, "layout", token)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	utilities.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccout(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utilities.Danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utilities.Danger(err, "Cannot create user")
		utilities.RespErrMsg(writer, request, "account have exist")
	} else {
		utilities.Info("created user:" + user.Email)
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		utilities.Danger(err, "Cannot find user")
		utilities.RespErrMsg(w, r, "no account")
		return
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utilities.Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		utilities.RespErrMsg(w, r, "password incorrect")
	}
}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utilities.Warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
