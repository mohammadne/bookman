package main

import "github.com/labstack/echo"

func main() {
	var e = echo.New()

	routeUrls(e)
	e.Start(":8080")
}
