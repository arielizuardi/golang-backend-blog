package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arielizuardi/golang-backend-blog/article/handler"
	"github.com/arielizuardi/golang-backend-blog/article/mocks"
	"github.com/arielizuardi/golang-backend-blog/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestArticleHandler_HandleCreateArticle(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`{"title":"my-title", "content":"my-content", "author":"my-author"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	u := new(mocks.Usecase)
	u.On("CreateArticle", &model.Article{Title: "my-title", Content: "my-content", Author: "my-author"}).Return(nil)

	h := handler.NewArticleHTTPHandler(u)

	// Assertions
	if assert.NoError(t, h.HandleCreateArticle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"status":200,"message":"Success","data":{"id":0}}`, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestArticleHandler_HandleCreateArticle_UsecaseReturnError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`{"title":"my-title", "content":"my-content", "author":"my-author"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	u := new(mocks.Usecase)
	u.On("CreateArticle", &model.Article{Title: "my-title", Content: "my-content", Author: "my-author"}).Return(validator.ValidationErrors{})

	h := handler.NewArticleHTTPHandler(u)

	// Assertions
	if assert.NoError(t, h.HandleCreateArticle(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `{"status":400,"message":"","data":null}`, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestArticleHandler_HandleGetArticleByID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles/:article_id")
	c.SetParamNames("article_id")
	c.SetParamValues("1")

	result := &model.Article{Title: "my-title", Content: "my-content", Author: "my-author"}
	u := new(mocks.Usecase)
	u.On("GetArticleByID", 1).Return(result, nil)

	h := handler.NewArticleHTTPHandler(u)
	// Assertions
	if assert.NoError(t, h.HandleGetArticleByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"status\":200,\"message\":\"Success\",\"data\":{\"id\":0,\"title\":\"my-title\",\"content\":\"my-content\",\"author\":\"my-author\"}}", strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestArticleHandler_HandleGetArticleByID_ArticleDoesNotExist(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles/:article_id")
	c.SetParamNames("article_id")
	c.SetParamValues("1")

	u := new(mocks.Usecase)
	u.On("GetArticleByID", 1).Return(nil, nil)

	h := handler.NewArticleHTTPHandler(u)
	// Assertions
	if assert.NoError(t, h.HandleGetArticleByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "{\"status\":404,\"message\":\"article is not found\",\"data\":null}", strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestArticleHandler_HandleGetAllArticle(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result := []*model.Article{
		&model.Article{Title: "my-title", Content: "my-content", Author: "my-author"},
		&model.Article{Title: "my-title-2", Content: "my-content-2", Author: "my-author-2"},
	}

	u := new(mocks.Usecase)
	u.On("GetAllArticle").Return(result, nil)

	h := handler.NewArticleHTTPHandler(u)
	// Assertions
	if assert.NoError(t, h.HandleGetAllArticle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"status\":200,\"message\":\"Success\",\"data\":[{\"id\":0,\"title\":\"my-title\",\"content\":\"my-content\",\"author\":\"my-author\"},{\"id\":0,\"title\":\"my-title-2\",\"content\":\"my-content-2\",\"author\":\"my-author-2\"}]}", strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestArticleHandler_HandleGetAllArticle_RepositoryReturnError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	result := []*model.Article{}

	u := new(mocks.Usecase)
	u.On("GetAllArticle").Return(result, errors.New("whoops"))

	h := handler.NewArticleHTTPHandler(u)
	// Assertions
	if assert.NoError(t, h.HandleGetAllArticle(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "{\"status\":500,\"message\":\"whoops\",\"data\":null}", strings.Trim(rec.Body.String(), "\n"))
	}
}
