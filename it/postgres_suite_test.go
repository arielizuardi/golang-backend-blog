// +build integration

package it_test

import (
	"log"
	"testing"

	"github.com/arielizuardi/sph-backend-coding-challenge/config"

	"github.com/golang-migrate/migrate"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	connStr string
	err     error
)

type PostgresRepositoryTestSuite struct {
	gormDB *gorm.DB
	suite.Suite
}

func (p *PostgresRepositoryTestSuite) SetupSuite() {
	if err := config.Load(); err != nil {
		logrus.WithError(err).Fatal("Failed to load config")
	}

	connStr = config.DBConnectionURL()
	p.gormDB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Can't connect to posgres", err)
	}
}

func TestPostgresRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &PostgresRepositoryTestSuite{})
}

func (p *PostgresRepositoryTestSuite) SetupTest() {
	m, err := migrate.New("file://../db/migration", connStr)
	assert.NoError(p.T(), err)

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			// just ignore
			return
		}

		panic(err)
	}
}

func (p *PostgresRepositoryTestSuite) TearDownTest() {
	m, err := migrate.New("file://../db/migration", connStr)
	assert.NoError(p.T(), err)
	assert.NoError(p.T(), m.Down())
}
