
package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Static("/static", "src/web")

    // Route => handler
    e.File("/", "src/web/index.html")
    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}