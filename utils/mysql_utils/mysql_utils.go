package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jcabrera/bookstore_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("duplicated key"))
	}
	return errors.NewInternalServerError("error procesing request ")

}


