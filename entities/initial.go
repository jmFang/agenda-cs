package entities

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var mydb *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	mydb = db
	createMeetingTable()
	createUserTable()

}

type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecer
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createMeetingTable() {
	tblStmt := `create table if not exists meeting(
                    sponsor text not null,
                    title text primary key,
                    start not null,
                    end not null,
                    participators blob not null
                    
                    );`
	_, err := mydb.Exec(tblStmt)
	checkErr(err)
}

func createUserTable() {
	tblStmt := `create table if not exists user(
                    uid integer primary key autoincrement,
                    username varchar(64) not null unique,
                    password varchar(64) not null,
                    email varchar(64) not null,
                    phone varchar(64) not null
                    );`
	_, err := mydb.Exec(tblStmt)
	checkErr(err)
}
