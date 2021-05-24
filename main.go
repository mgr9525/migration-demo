package main

import (
	"database/sql"
	"flag"
	"github.com/adelowo/migration-demo/deps"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	//var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	var mysqlDSN = flag.String("mysql.dsn", os.Getenv("MYSQL_DSN"), "Mysql DSN")

	flag.Parse()

	db, err := sql.Open("mysql", *mysqlDSN)
	if err != nil {
		log.Fatalf("could not connect to postgresql database... %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	// Run migrations
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	s := bindata.Resource(deps.AssetNames(), deps.Asset)
	sc, err := bindata.WithInstance(s)
	if err != nil {
		log.Fatalf("could not bindata... %v", err)
	}
	m, err := migrate.NewWithInstance(
		"bindata", sc,
		"mysql", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
	os.Exit(0)
}
