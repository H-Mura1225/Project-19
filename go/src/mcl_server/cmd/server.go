package main

import (
	"net/http"

	"src/mcl_server/pkg/keyHandler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://*"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "p1", "p2"},
		AllowMethods: []string{echo.GET},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "暗号とセキュリティ2020")
	})
	e.GET("mcl_c384_256.js", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.GET("mcl_c384_256.wasm", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.GET("/signup", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.GET("/secretkey", keyHandler.AuthMiddleware(keyHandler.Secretkey))
	e.GET("/secretkey2", keyHandler.AuthMiddleware(keyHandler.Secretkey2))
	e.GET("/publickey", keyHandler.AuthMiddleware(keyHandler.Publickey))
	e.GET("/publickey2", keyHandler.AuthMiddleware(keyHandler.Publickey2))
	e.GET("/KtTimeKey", keyHandler.KtTimeKey)
	e.GET("/TtTimeKey", keyHandler.TtTimeKey)

	e.Logger.Fatal(e.Start(":8080"))
}
