package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type book struct {
	id    int    `json:id`
	name  string `json:name`
	title string `json:title`
}

var books = []book{book{id: 1, name: "book1", title: "title1"}, book{id: 2, name: "book2", title: "title2"}}

func main() {
	router := gin.Default()
	//tmpl, err := template.New("name").Parse(...)
	//var funcMap template.FuncMap
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper})
	router.LoadHTMLGlob("template/*.html")
	router.GET("/", getBooks)
	router.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", gin.H{
			"content": "This is an about page...",
		})
	})

}
func getBooks(c *gin.Context) {
	if books == nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(http.StatusOK, books)
	}

}
