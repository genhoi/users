package app

import (
	"database/sql"
	"fmt"
	"github.com/aofei/air"
	"github.com/deanishe/go-env"
	"github.com/genhoi/users/app/actions"
	"github.com/genhoi/users/module/db"
	"github.com/genhoi/users/module/user"
	"github.com/genhoi/users/module/user/jsonl"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	instance      *container
	onceContainer sync.Once

	onceDb           sync.Once
	onceAir          sync.Once
	onceMigrate      sync.Once
	onceUserImporter sync.Once
	onceUserRepo     sync.Once

	onceSearchAction sync.Once
	onceUiAction     sync.Once
)

type container struct {
	config config

	logger       logrus.FieldLogger
	db           *sql.DB
	air          *air.Air
	migrate      *migrate.Migrate
	userImporter *jsonl.Importer
	userRepo     *user.Repository
	searchAction *actions.Search
	uiAction     *actions.Ui
}

func (c *container) UiAction() *actions.Ui {
	onceUiAction.Do(func() {
		c.uiAction = &actions.Ui{}
	})

	return c.uiAction
}

func (c *container) SearchAction() *actions.Search {
	onceSearchAction.Do(func() {
		c.searchAction = actions.NewSearch(c.UserRepo())
	})

	return c.searchAction
}

func (c *container) Air() *air.Air {
	onceAir.Do(func() {
		c.air = air.New()
		c.air.DebugMode = c.config.debug
		c.air.Address = c.config.listenAddr
	})

	return c.air
}

func (c *container) UserRepo() *user.Repository {
	onceUserRepo.Do(func() {
		database := c.Db()
		c.userRepo = user.NewRepository(database)
	})

	return c.userRepo
}

func (c *container) UserImporter() *jsonl.Importer {
	onceUserImporter.Do(func() {
		query := db.NewQuery("users")
		database := c.Db()

		c.userImporter = jsonl.NewImporter(query, database)
	})

	return c.userImporter
}

func (c *container) Migrate() *migrate.Migrate {
	onceMigrate.Do(func() {
		var err error
		c.migrate, err = migrate.New(c.config.migrationsDSN, c.config.databaseDSN)
		if err != nil {
			panic(err)
		}
	})

	return c.migrate
}

func (c *container) Db() *sql.DB {
	onceDb.Do(func() {
		c.db = stdlib.OpenDB(c.config.databaseConfig)
	})

	return c.db
}

func (c *container) Logger() logrus.FieldLogger {
	return c.logger
}

func Container() *container {
	onceContainer.Do(func() {
		instance = createContainer()
	})

	return instance
}

type config struct {
	databaseDSN    string
	databaseConfig pgx.ConnConfig
	listenAddr     string
	migrationsDSN  string
	debug          bool
}

func createContainer() *container {
	port := env.Get("APP_PORT", "80")
	databaseDsn := env.Get("APP_DATABASE", "postgres://postgres:root@postgres:5432/users?sslmode=disable&timezone=GMT")
	migrationsDsn := env.Get("APP_MIGRATIONS", "file://migrations")
	debugEnv := env.Get("APP_DEBUG", "")
	debug := false
	if debugEnv != "" {
		debug = true
	}

	dbConf, err := pgx.ParseConfig(databaseDsn)
	if err != nil {
		panic(err)
	}

	conf := config{
		listenAddr:     fmt.Sprintf(":%s", port),
		databaseDSN:    databaseDsn,
		databaseConfig: *dbConf,
		migrationsDSN:  migrationsDsn,
		debug:          debug,
	}

	return &container{config: conf, logger: logrus.StandardLogger()}
}
