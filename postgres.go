package gsql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	username string
	password string
	database string
	port     string
	sslmode  string
)

func init() {
	config, err := ReadConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	username = config.DbConfig.User
	password = config.DbConfig.Password
	database = config.DbConfig.Database
	port = config.DbConfig.Port
	sslmode = config.DbConfig.SSLMode
}

func DbConnect() int {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", username, password, database, port)
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()
	if err != nil {
		fmt.Printf("error %v", err)
		return 1
	}
	rows, err := db.Query("SELECT name FROM account")
	defer rows.Close()
	if err != nil {
		fmt.Printf("error %v", err)
		return 1
	}
	if rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("error %v", err)
			return 1
		}
		fmt.Println("the name is ", name)
	} else {
		fmt.Println("there is no user.")
	}
	return 0
}
