package db

import (
	
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log")



func InitDb(datasourcename string) (*sql.DB, error)  {
	db,err:= sql.Open("sqlite3",datasourcename)
	if err!= nil {
		log.Fatal(err)
		return nil,err}

	createtablesql:= `
	
	CREATE TABLE IF NOT EXISTS user (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS product (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL UNIQUE,
		"price" INTEGER NOT NULL,
		"image" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"category" TEXT NOT NULL,
		"quantity" INTEGER NOT NULL
	)`

	_,err=db.Exec(createtablesql)
	if err!= nil {
		log.Fatal(err)
		return nil,err
	}
	return db,nil
}

