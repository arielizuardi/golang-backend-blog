package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arielizuardi/sph-backend-coding-challenge/article/handler"
	"github.com/arielizuardi/sph-backend-coding-challenge/article/mocks"
	"github.com/arielizuardi/sph-backend-coding-challenge/model"
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
