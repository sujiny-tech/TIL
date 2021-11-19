package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // go-sqlite3 라이브러리 import
)

func main() {
	//os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-example.db") // "sqlite-example.db" SQLite file 생성
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-example.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-example.db") // 해당 SQLite File 열기
	defer sqliteDatabase.Close()                                    // Defer Closing the database **
	createTable(sqliteDatabase)                                     // Database Tables 생성

	// 데이터 삽입
	insertUser(sqliteDatabase, "1", "Sujiny", "00")
	insertUser(sqliteDatabase, "2", "Jiny", "01")
	insertUser(sqliteDatabase, "3", "Tech", "02")
	insertUser(sqliteDatabase, "4", "sujiny-tech", "03")

	// 해당 테이블에 있는 데이터 출력
	displayUser(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE User (
		"Id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"password" TEXT,
		UNIQUE(Id, name)
	  );` // SQL Statement for Create Table

	log.Println("Create User table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("User table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, Id string, name string, password string) {
	log.Println("Inserting User record ...")
	insertUserSQL := `INSERT INTO User (Id, name, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.

	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Id, name, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayUser(db *sql.DB) {
	row, err := db.Query("SELECT * FROM User ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var Id int
		var name string
		var password string
		row.Scan(&Id, &name, &password)
		log.Println("User: ", Id, " ", name, " ", password)
	}
}
