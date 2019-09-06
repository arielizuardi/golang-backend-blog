// +build integration

package it_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"

	"github.com/arielizuardi/sph-backend-coding-challenge/config"
	"github.com/arielizuardi/sph-backend-coding-challenge/model"
	"github.com/arielizuardi/sph-backend-coding-challenge/server"
)

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
	dbMigration     *migrate.Migrate
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.Require().NoError(config.Load())

	s.port = config.Port()
	s.dbConnectionStr = config.DBConnectionURL()
	dbConn, err := gorm.Open("postgres", s.dbConnectionStr)
	s.Require().NoError(err)
	s.dbConn = dbConn

	s.dbMigration, err = migrate.New("file://../db/migration", s.dbConnectionStr)
	s.Require().NoError(err)
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}

	serverReady := make(chan bool)

	server := server.Server{
		Port:        config.Port(),
		DBConn:      s.dbConn,
		ServerReady: serverReady,
	}

	go server.Start()
	<-serverReady
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) SetupTest() {
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}
}

func (s *e2eTestSuite) TearDownTest() {
	s.NoError(s.dbMigration.Down())
}

func (s *e2eTestSuite) Test_EndToEnd_CreateArticle() {
	reqStr := `{"title":"e2eTitle", "content": "e2eContent", "author":"e2eauthor"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/articles", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":{"id":1}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetArticleByID() {
	article := model.Article{
		Title:   "my-title",
		Author:  "my-author",
		Content: "my-content",
	}

	s.NoError(s.dbConn.Create(&article).Error)

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/articles/1", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":{"id":1,"title":"my-title","content":"my-content","author":"my-author"}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllArticle() {
	article := model.Article{
		Title:   "my-title",
		Author:  "my-author",
		Content: "my-content",
	}

	s.NoError(s.dbConn.Create(&article).Error)

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/articles", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":[{"id":1,"title":"my-title","content":"my-content","author":"my-author"}]}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
