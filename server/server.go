package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arielizuardi/golang-backend-blog/article/handler"
	"github.com/arielizuardi/golang-backend-blog/article/repository"
	"github.com/arielizuardi/golang-backend-blog/article/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Server represents server
type Server struct {
	Port        int
	DBConn      *gorm.DB
	ServerReady chan bool
}

// Start start http server
func (s *Server) Start() {
	appPort := fmt.Sprintf(":%d", s.Port)

	repo := repository.NewPostgresArticleRepository(s.DBConn)
	articleUsecase := usecase.NewArticleUsecase(repo)
	articleHandler := handler.NewArticleHTTPHandler(articleUsecase)

	e := echo.New()

	e.POST("/articles", articleHandler.HandleCreateArticle)
	e.GET("/articles", articleHandler.HandleGetAllArticle)
	e.GET("/articles/:article_id", articleHandler.HandleGetArticleByID)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"pong": "ok",
		})
	})

	go func() {
		if err := e.Start(appPort); err != nil {
			logrus.Errorf(err.Error())
			logrus.Infof("shutting down the server")
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
