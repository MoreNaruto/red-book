package api

import (
	"database/sql"
	"log"
	"redbook/internal/pkg/encryption"
	"redbook/pkg/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserApi struct {
	db *sql.DB
}

func (d *UserApi) CreateUser(u *models.User) (int64, error) {
	insertStatement := "INSERT INTO users(username, password, birthday, email) VALUES($1, $2, $3, $4)"
	statement, err := d.db.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	hashPassword := encryption.HashAndSalt([]byte(u.GetPassword()))
	result, err := statement.Exec(insertStatement, u.GetUsername(), hashPassword, u.GetBirthday(), u.GetEmail())
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	log.Printf("Created user with id = %d and the number of rows affected = %d\n", lastId, rowCnt)
	return lastId, err
}

func (d *UserApi) DeleteUser(u int32) error {
	return nil
}
func (d *UserApi) GetUser(u int32) (*models.User, error) {
	return nil, nil
}
