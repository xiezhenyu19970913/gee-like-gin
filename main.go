package main

import (
	"7day/day1/gee"
	"net/http"
)

func main() {
	r:=gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK,"Hello Geektutu\n")
	})
	r.GET("/panic", func(c *gee.Context) {
		name:=[]string{"geektutu"}
		c.String(http.StatusOK,name[100])
	})
	r.Run(":8080")
}