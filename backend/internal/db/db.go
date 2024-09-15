package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDb(datasourcename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", datasourcename)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	// Define the SQL statements to create tables
	createtablesql := `
	CREATE TABLE IF NOT EXISTS taille (
		"taille" TEXT PRIMARY KEY NOT NULL,
		"quantity" INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS user (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL,
		"isadmin" BOOLEAN NOT NULL DEFAULT 0,
	    "isauthenticated" BOOLEAN NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS product (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL UNIQUE,
		"price" INTEGER NOT NULL,
		"image" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"category" TEXT NOT NULL,
		"quantity" INTEGER NOT NULL,
		"taille" TEXT NOT NULL,
		FOREIGN KEY ("taille") REFERENCES taille("taille")
	);

	CREATE TABLE IF NOT EXISTS commande (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER NOT NULL,
		"product_id" INTEGER NOT NULL,
		"quantity" INTEGER NOT NULL,
		"taille" TEXT NOT NULL,
		FOREIGN KEY ("user_id") REFERENCES user("id"),
		FOREIGN KEY ("product_id") REFERENCES product("id"),
		FOREIGN KEY ("taille") REFERENCES taille("taille")
	);
	`

	// Execute the SQL statements to create tables
	_, err = db.Exec(createtablesql)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return nil, err
	}

	return db, nil
}
