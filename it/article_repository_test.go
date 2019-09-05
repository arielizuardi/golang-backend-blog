// +build integration

package it_test

import (
	"github.com/arielizuardi/sph-backend-coding-challenge/article/repository"
	"github.com/arielizuardi/sph-backend-coding-challenge/model"
)

func (p *PostgresRepositoryTestSuite) TestPostgresArticleRepository_CreateArticle() {
	newArticle := &model.Article{
		Title:   "my-title",
		Content: "my-content",
		Author:  "my-author",
	}

	r := repository.NewPostgresArticleRepository(p.gormDB)
	p.Assert().NoError(r.CreateArticle(newArticle))

	id := newArticle.ID // new id is created

	var result model.Article
	p.Assert().NoError(p.gormDB.First(&result, model.Article{ID: id}).Error)
	p.Assert().Equal(newArticle.Title, result.Title)
	p.Assert().Equal(newArticle.Content, result.Content)
	p.Assert().Equal(newArticle.Author, result.Author)
}

func (p *PostgresRepositoryTestSuite) TestPostgresArticleRepository_GetArticleByID() {
	myArticle := &model.Article{
		Title:   "find-my-title",
		Content: "find-my-content",
		Author:  "find-my-author",
	}

	r := repository.NewPostgresArticleRepository(p.gormDB)
	p.Assert().NoError(r.CreateArticle(myArticle))

	result, err := r.GetArticleByID(myArticle.ID)
	p.Assert().NoError(err)
	p.Assert().Equal(myArticle.Title, result.Title)
	p.Assert().Equal(myArticle.Content, result.Content)
	p.Assert().Equal(myArticle.Author, result.Author)
}

func (p *PostgresRepositoryTestSuite) TestPostgresArticleRepository_GetAllArticle() {
	firstArticle := &model.Article{
		Title:   "first-title",
		Content: "first-content",
		Author:  "first-author",
	}

	secondArticle := &model.Article{
		Title:   "second-title",
		Content: "second-content",
		Author:  "second-author",
	}

	r := repository.NewPostgresArticleRepository(p.gormDB)
	p.Assert().NoError(r.CreateArticle(firstArticle))
	p.Assert().NoError(r.CreateArticle(secondArticle))

	result, err := r.GetAllArticle()
	p.Assert().NoError(err)
	p.Assert().Len(result, 2)
}
