package usecase_test

import (
	"errors"
	"testing"

	"github.com/arielizuardi/sph-backend-coding-challenge/article"
	"github.com/arielizuardi/sph-backend-coding-challenge/article/mocks"
	"github.com/arielizuardi/sph-backend-coding-challenge/article/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {
	testCases := []struct {
		name          string
		article       *article.Article
		repoError     error
		expectedError error
	}{
		{
			"Create Article with required fields",
			&article.Article{
				Title:   "my-title",
				Content: "this-is-content",
				Author:  "my-author",
			},
			nil,
			nil,
		},
		{
			"Create Article with missing fields",
			&article.Article{
				Title:   "my-title",
				Content: "this-is-content",
			},
			nil,
			errors.New("Key: 'Article.Author' Error:Field validation for 'Author' failed on the 'required' tag"),
		},
		{
			"Create Article and repository return error",
			&article.Article{
				Title:   "my-title",
				Content: "this-is-content",
				Author:  "my-author",
			},
			errors.New("whoops"),
			errors.New("whoops"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := new(mocks.Repository)
			repo.On("CreateArticle", testCase.article).Return(testCase.repoError)
			u := usecase.NewArticleUsecase(repo)
			err := u.CreateArticle(testCase.article)
			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestGetArticleByID(t *testing.T) {
	testCases := []struct {
		name          string
		articleID     int
		returnArticle *article.Article
		repoError     error
		expectedError error
	}{
		{
			"Article exists with given ID",
			1,
			&article.Article{
				ID:      1,
				Title:   "my-title",
				Content: "this-is-content",
				Author:  "my-author",
			},
			nil,
			nil,
		},
		{
			"Article does not exist with given ID",
			1,
			nil,
			nil,
			nil,
		},
		{
			"Repository return error",
			1,
			nil,
			errors.New("whoops"),
			errors.New("whoops"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := new(mocks.Repository)
			repo.On("GetArticleByID", testCase.articleID).Return(testCase.returnArticle, testCase.repoError)
			u := usecase.NewArticleUsecase(repo)
			result, err := u.GetArticleByID(testCase.articleID)
			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError.Error(), err.Error())
			}

			assert.Equal(t, testCase.returnArticle, result)
		})
	}
}

func TestGetAllArticle(t *testing.T) {
	testCases := []struct {
		name string

		returnArticles []*article.Article
		repoError      error
		expectedError  error
	}{
		{
			"Number of articles greater than zero",
			[]*article.Article{
				{ID: 1, Title: "my-title", Content: "this-is-content", Author: "my-author"},
			},
			nil,
			nil,
		},
		{
			"Number of articles is zero",
			nil,
			nil,
			nil,
		},
		{
			"Repository return error",
			nil,
			errors.New("whoops"),
			errors.New("whoops"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := new(mocks.Repository)
			repo.On("GetAllArticle").Return(testCase.returnArticles, testCase.repoError)
			u := usecase.NewArticleUsecase(repo)
			result, err := u.GetAllArticle()
			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError.Error(), err.Error())
			}

			assert.Equal(t, len(testCase.returnArticles), len(result))
		})
	}
}
