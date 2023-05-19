package handler

import (
	"net/http"
	"sample/models"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
}

func InitArticle() ArticleHandler {
	return ArticleHandler{}
}

func (h ArticleHandler) FetchArticles(c echo.Context) (err error) {
	datas := []models.Article{
		models.Article {
			ID: "1"
			Title: "Hello World! This is Title"
			Body: "Hello World! This is Body"
		}
	}

	return c.JSON(http.StatusOK, datas)
}