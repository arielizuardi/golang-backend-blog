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
