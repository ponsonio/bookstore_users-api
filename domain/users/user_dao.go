package users

import (
	"fmt"
	"github.com/jcabrera/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jcabrera/bookstore_users-api/utils/date_utils"
	"github.com/jcabrera/bookstore_users-api/utils/errors"
	"github.com/jcabrera/bookstore_users-api/utils/mysql_utils"
)

var(
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser ="INSERT INTO users(first_name, last_name, email, date_created, password, status) values (?,?,?,?,?,?) "
	queryGetUser ="SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ? "
	queryUpdateUser = "UPDATE users SET first_name=? , last_name=?, email=? WHERE id = ?"
	queryDeleteUser = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
)


func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status) ; getErr != nil {
		mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowDatabaseFormatString()
	errVal := user.Validate()
	if errVal != nil {
		return errVal
	}
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.Id = userId
	return nil
}


func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User , *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var results = make([]User, 0)

	for rows.Next(){
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status: %s",status))
	}
	return results, nil
}
