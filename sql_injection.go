package main

import "database/sql"

func LoadUser(db *sql.DB, username string) {
    // vulnerable query
    db.Query("SELECT * FROM users WHERE username = '" + username + "'")
}

