package model

import (
	"database/sql"
	"../config"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Datastore interface {

}

type DB struct {
	*sql.DB
}

const MAX_RETRIES = 3

func NewDb(params config.MySQLConfig)(*DB, error) {
	conf := mysql.Config{User: params.User, Passwd: params.Pass, DBName: params.Database, Net: "tcp", Addr: params.Url+":"+params.Port}

	log.Printf("%+v", conf)

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}