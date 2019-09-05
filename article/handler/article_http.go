package handler

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/arielizuardi/sph-backend-coding-challenge/article"
	"github.com/arielizuardi/sph-backend-coding-challenge/model"
	"github.com/labstack/echo"
)

// ArticleHTTPHandler represents article http handler
type ArticleHTTPHandler struct {
	Usecase article.Usecase
}

// NewArticleHTTPHandler return new instances of article http handler
func NewArticleHTTPHandler(usecase article.Usecase) *ArticleHTTPHandler {
	return &ArticleHTTPHandler{Usecase: usecase}
}

// HandleCreateArticle handle create article
func (h *ArticleHTTPHandler) HandleCreateArticle(c echo.Context) error {
	var article model.Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &Response{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := h.Usecase.CreateArticle(&article); err != nil {
		httpStatus := http.StatusInternalServerError
		if reflect.TypeOf(err).String() == `validator.ValidationErrors` {
			httpStatus = http.StatusBadRequest
		}

		return c.JSON(httpStatus, &Response{
			Status:  httpStatus,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]int{
			"id": article.ID,
		},
	})
}

// HandleGetArticleByID handle get article by id
func (h *ArticleHTTPHandler) HandleGetArticleByID(c echo.Context) error {
	articleIDStr := c.Param("article_id")
	id, err := strconv.Atoi(articleIDStr)
	res, err := h.Usecase.GetArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if res == nil {
		return c.JSON(http.StatusNotFound, &Response{
			Status:  http.StatusNotFound,
			Message: "article is not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// HandleGetAllArticle handle get all article
func (h *ArticleHTTPHandler) HandleGetAllArticle(c echo.Context) error {
	res, err := h.Usecase.GetAllArticle()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    res,
	})
}
