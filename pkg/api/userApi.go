package api

import (
	"log"
	"redbook/internal/pkg/encryption"
	"redbook/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateUser(u *models.User) (int64, error) {
	insertStatement := "INSERT INTO users(username, password, birthday, email, created_at, last_updated) VALUES(?, ?, ?, ?, ?, ?)"
	statement, err := dbmap.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer statement.Close()
	hashPassword := encryption.HashAndSalt([]byte(u.GetPassword()))
	result, err := statement.Exec(u.GetUsername(), hashPassword, u.GetBirthday().AsTime(), u.GetEmail(), timestamppb.Now().AsTime(), timestamppb.Now().AsTime())
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

func DeleteUser(u int32) error {
	return nil
}
func GetUser(u int32) (*models.User, error) {
	return nil, nil
}
