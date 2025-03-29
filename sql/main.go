package sql

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      int
	Name    string
	EmpID   int
	Gender  string
	Premium int
}

func (usr *User) sanitize() error {
	err := errors.New("Wrong Values")
	if usr.Name == "" {
		return err
	}
	if !(usr.EmpID == 999 || (usr.EmpID <= 100 && usr.EmpID%1 == 0 && usr.EmpID != 0)) {
		return err
	}
	return nil
}

func Add(usr *User) (id int64, err error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	err = usr.sanitize()
	if err != nil {
		return 0, err
	}

	insertSQL := `INSERT INTO data (name, emp_id,gender,premium) VALUES (?, ?,?,?)`
	result, err := db.Exec(insertSQL, usr.Name, usr.EmpID, usr.Gender, usr.Premium)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Delete(id int) error {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	deleteSQL := `DELETE FROM data WHERE id = ?`
	result, err := db.Exec(deleteSQL, id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("No Value Found with the Id")
	}
	return nil

}

func Get(id int) (usr *User, err error) {
	var usr_id int
	var name string
	var emp_id int
	var gender string
	var premium int

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return &User{}, err
	}
	defer db.Close()

	query := `SELECT * FROM data WHERE id = ?`
	err = db.QueryRow(query, id).Scan(&usr_id, &name, &emp_id, &gender, &premium)
	if err != nil {
		return &User{}, err
	}

	user := User{
		Id:      usr_id,
		Name:    name,
		EmpID:   emp_id,
		Gender:  gender,
		Premium: premium,
	}
	return &user, nil
}

func Set(usr *User) error {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	updateSQL := `UPDATE data SET name = ?, emp_id = ?, gender=?, premium=? WHERE id = ?;`
	result, err := db.Exec(updateSQL, usr.Name, usr.EmpID, usr.Gender, usr.Premium, usr.Id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("No Value Found with the Id")
	}
	return nil
}
