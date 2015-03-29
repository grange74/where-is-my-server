package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // global variable to share it between main and the HTTP handler

func NewDB() {
	var err error
	db, err = sql.Open("sqlite3", "where-is-my-server.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS servers(name text, ip text)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(name text, email text, server_name text)")
	if err != nil {
		panic(err)
	}
}

func GetServersFromDB() []Server {
	rows, err := db.Query("SELECT name, ip FROM servers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	servers := []Server{}
	for rows.Next() {
		var name, ip string
		if err := rows.Scan(&name, &ip); err != nil {
			panic(err)
		}
		servers = append(servers, Server{name, ip})
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return servers
}

func AddServerToDB(server Server) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO servers VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(server.Name, server.IP)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
