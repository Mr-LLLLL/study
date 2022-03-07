package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func Test(t *testing.T) {
	TestingT(t)
}

// first execute function
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	Suite(&PostTestSuite{})
}

func tearDown() {}

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Content, Equals, "Updated post")
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing post test suite")
}
