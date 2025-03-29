package pkg

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {

	//conStr := "postgres://postgres:password@localhost:5050/"
	db, err := sql.Open("mysql",
		"root:password@tcp(localhost:3306)/newdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createProductTable(db)

}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal()
	}
}
