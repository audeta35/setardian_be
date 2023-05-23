package main

import (
	"database/sql"
	"log"
	httpHandler "steradian_be/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "steradian"

	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	handler := httpHandler.InitSteradian(db)
	echoServer := echo.New()

	// articles handler
	echoServer.POST("/articles", handler.FetchArticles)
	echoServer.POST("/articles/add", handler.InsertArticles)
	echoServer.POST("/articles/:id", handler.GetDetailArticle)
	echoServer.POST("/articles/edit", handler.EditArticles)
	echoServer.DELETE("/articles/delete/:id", handler.DeleteArticle)

	// admin handler
	echoServer.POST("/admin/login", handler.LoginAdmin)
	echoServer.POST("/admin/register", handler.RegisterAdmin)

	// user handler
	echoServer.POST("/user/login", handler.LoginUser)
	echoServer.POST("/user/register", handler.RegisterUser)

	// order handler
	echoServer.POST("/orders", handler.OrderFetch)
	echoServer.POST("/orders/add", handler.OrderAdd)
	echoServer.POST("/orders/edit", handler.OrderEdit)
	echoServer.DELETE("/orders/delete", handler.OrderDelete)

	// cars handler
	echoServer.POST("/cars", handler.CarsFetch)
	echoServer.POST("/cars/add", handler.CarsAdd)
	echoServer.POST("/cars/edit", handler.CarsEdit)
	echoServer.DELETE("/cars/delete", handler.CarsDelete)

	// Start the server
	echoServer.Start(":9090")
}
