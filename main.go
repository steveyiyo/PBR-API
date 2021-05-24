package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func indexPage(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func apiService(c *gin.Context) {

	type Result struct {
		Success    bool
		Message    string
		IP_Address string
		Node       string
	}

	var r Result
	r = Result{false, "NULL", c.ClientIP(), "NULL"}

	upstream := c.PostForm("upstream")

	if upstream == "" {
		r = Result{false, "Bad Request!", c.ClientIP(), ""}
		c.JSON(400, r)
	} else if upstream == "TW" {
		if reroute(c.ClientIP(), "TW") {
			r = Result{true, fmt.Sprintf("Your route has been reroute to %s upstream!", upstream), c.ClientIP(), upstream}
			c.JSON(200, r)
		}
	} else if upstream == "JP" {
		if reroute(c.ClientIP(), "JP") {
			r = Result{true, fmt.Sprintf("Your route has been reroute to %s upstream!", upstream), c.ClientIP(), upstream}
			c.JSON(200, r)
		}
	} else {
		r = Result{false, "Bad Request!", c.ClientIP(), ""}
		c.JSON(400, r)
	}
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func reroute(IP_Address string, Node string) bool {
	command := fmt.Sprintf("ip rule add from %s lookup %s", IP_Address, Node)
	cmd := exec.Command("bash", "-c", command)
	cmd.Run()
	return true
}

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("static/*")

	router.GET("/", indexPage)
	router.POST("/api", apiService)
	router.NoRoute(pageNotAvailable)

	router.Run(":1080")
}
