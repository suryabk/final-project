package repository

import (
	"database/sql"
	"final-project/structs"
	"time"
)

func RegisterUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (email, password, created_at, updated_at) VALUES($1, $2, $3, $4)"
	errs := db.QueryRow(sql, user.Email, user.Password, time.Now(), time.Now())

	return errs.Err()
}

func LoginUser(db *sql.DB, user structs.User) (result structs.User, err error) {
	sql := "SELECT * FROM users WHERE email = $1"
	rows, err := db.Query(sql, user.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err := rows.Scan(&user.UserID, &user.Email, &user.Password, &user.UpdatedAt, &user.CreatedAt)
		if err != nil {
			panic(err)
		}

		result = user
	}

	return
}
