package database

import (
	"database/sql"

	envConfig "github.com/gfbatista/xy-inc/util"
)

//Exec funcão para executar um comando sql
func Exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	return result
}

//Create função para criar a base
func Create() {
	password := envConfig.Get().Database.Password
	usuario := envConfig.Get().Database.Usuario

	db, err := sql.Open("mysql", usuario+":"+password+"@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Exec(db, "create database if not exists xyinc")
	Exec(db, "use xyinc")
	Exec(db, "drop table if exists poi")
	Exec(db, `create table poi (
		id integer auto_increment,
		nome varchar(120),
		coordenada_x int(10),
		coordenada_y int(10),
		PRIMARY KEY (id)
	)`)
}

// OpenConection função para abrir a conexão
func OpenConection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:a@/xyinc")

	return db, err
}
