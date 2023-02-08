package db

import (
	"app/ent"
	"app/ent/migrate"
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
)

func Init() {
	db, err := ent.Open("sqlite3", "enber.sqlite?_fk=1&cahche=shared")
	if err != nil {
		log.Fatalln(err)
	}

	db.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
	)

	Client = db
}
