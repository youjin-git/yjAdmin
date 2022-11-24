package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
