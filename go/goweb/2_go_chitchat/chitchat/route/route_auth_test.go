package route

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"text/template"
	"time"
)

func Test_Get_Login(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("Get", "/login", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	// fmt.Println(body)
	if strings.Contains(body, "Sign in") == false {
		t.Errorf("Body does not contain Sign in")
	}
}

func Test_test(test *testing.T) {
	fmt.Println(time.Now())

	if testing.Short() {
		test.SkipNow()
	}
	files := []string{"login.layout.html", "public.navbar.html", "login.html"}
	for i := 0; i < len(files); i++ {
		files[i] = "templates/" + files[i]
	}
	// t := template.New("test")
	t := template.Must(template.ParseFiles(files...))
	fmt.Println(t)
	t.ExecuteTemplate(os.Stdout, "layout", nil)
}
