package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/simple-golang-echo/configs"
	"github.com/wildanfaz/simple-golang-echo/internal/pkg"
	"github.com/wildanfaz/simple-golang-echo/internal/repositories"
	"github.com/wildanfaz/simple-golang-echo/internal/services/books"
)

func InitRouter() {
	// init config
	cfg := configs.InitConfig()

	// init db
	db := configs.InitMySQL(cfg.MySqlDSN)

	// init logger
	log := pkg.InitLogger()

	// init echo router
	e := echo.New()
	apiV1 := e.Group("/api/v1")

	// repositories
	booksRepo := repositories.NewBooksRepository(db)

	// services
	booksSvc := books.NewService(booksRepo, log)

	// books
	apiV1.POST("/books", booksSvc.AddBook)
	apiV1.GET("/books/:id", booksSvc.GetBook)
	apiV1.GET("/books", booksSvc.ListBooks)
	apiV1.PUT("/books/:id", booksSvc.UpdateBook)
	apiV1.DELETE("/books/:id", booksSvc.DeleteBook)

	e.Logger.Fatal(e.Start(":1323"))
}
