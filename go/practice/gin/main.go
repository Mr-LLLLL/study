package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// jsonFormat(r)
	// htmlFormat(r)
	// query(r)
	// post(r)
	// path(r)
	// bindStruct(r)
	test(r)
	r.Run(":1234")
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func test(r *gin.Engine) {
	group := r.Group("hello")
	group.GET("test", func(c *gin.Context) {
		c.String(http.StatusOK, "test the router")
	})
	group.GET("test/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id)
	})
	r.GET("/test/hello/world", func(c *gin.Context) {
		req := new(UserInfo)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, "Parameters invalid")
		}
		c.JSON(http.StatusOK, gin.H{
			"name": req.Username,
			"pass": req.Password,
		})
	})
}

func bindStruct(r *gin.Engine) {
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/user", func(c *gin.Context) {
		var u UserInfo
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
}

func path(r *gin.Engine) {
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year": year,
			"mont": month,
		})
	})
}

func post(r *gin.Engine) {
	r.LoadHTMLFiles("templates/public/login.html", "templates/public/private.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("username")
		// password := c.PostForm("password")
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "1234")
		c.HTML(http.StatusOK, "private.html", gin.H{
			"Username": username,
			"Password": password,
		})
	})
}

func query(r *gin.Engine) {
	r.GET("/web", func(c *gin.Context) {
		// name := c.Query("query")
		name := c.DefaultQuery("query", "somebody")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
}

func jsonFormat(r *gin.Engine) {
	r.GET("/json", func(c *gin.Context) {
		// data := map[string]interface{}{
		//     "name":    "prince",
		//     "message": "hello world",
		//     "age":     18,
		// }
		data := gin.H{
			"name":    "prince",
			"message": "hello world",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}

	r.GET("/json1", func(c *gin.Context) {
		data := msg{
			Name:    "prince",
			Message: "hello world",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/sleep", func(c *gin.Context) {
		time.Sleep(time.Hour)
		data := msg{
			Name:    "prince",
			Message: "hello world",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
}

func htmlFormat(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.Static("/xxx", "./statics")

	// r.LoadHTMLFiles("templates/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts.com",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>liwenzhou blog</a>",
		})
	})
}
