package ugly_bot_db

import (
	"database/sql"
)

//connects to databse
func ConnectToDB() (db *sql.DB, err error) {
	db, okay := sql.Open("mysql",
		"root:mysql@tcp(127.0.0.1)/ugly_bot")
	if okay != nil {
		return nil, okay

	}
	return db, nil
}
