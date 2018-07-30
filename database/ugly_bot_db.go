//file contains functions to connect to db
//file contains functions to read/write to db
package ugly_bot_db

import (
	"database/sql"
)

//connects to databse
func ConnectToDB() (db *sql.DB, err error) {
	//change the 2nd argument to fit
	//the credentials for your own local database
	//root -> username
	//mysql -> password
	//tcp -> protocl type (leave that the same)
	//127.0.0.1 -> port where database is hosted locally
	//ugly_bot -> name i gave to the database instance
	db, okay := sql.Open("mysql",
		"root:mysql@tcp(127.0.0.1)/ugly_bot")
	if okay != nil {
		return nil, okay

	}
	return db, nil
}
