package service

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/db"
	"github.com/iiran/lltt/pkg/model"
)

func QueryUserDisplayname(username string) (displayname string, err error) {
	var (
		rows         *sql.Rows
		_displayname sql.NullString
	)
	if rows, err = db.Query(`pqs`, "select display_name from users where username = ?", username); err != nil {
		return displayname, err
	}
	defer rows.Close()
	if !rows.Next() {
		return displayname, errors.GetErr(errors.USER_NOT_FOUND)
	}
	if err = rows.Scan(&_displayname); err != nil {
		return "", err
	}
	return _displayname.String, nil
}

func SetUserDisplayname(username string, displayname string) (err error) {
	var (
		q = `update users set display_name = ? where username = ?`
	)
	if _, err = db.Exec(`pqs`, q, displayname, username); err != nil {
		return errors.ErrAppend(err, errors.DB_INSERT_ERR)
	}
	return nil
}

func QueryUser(username string) (user model.User, err error) {
	var (
		rows         *sql.Rows
		_displayname sql.NullString
		_email       sql.NullString
		q            string
	)
	q = `select id, display_name, email, created_at, score, weight from users where username = ?`
	if rows, err = db.Query(`pqs`, q, username); err != nil {
		return user, err
	}
	defer rows.Close()
	if !rows.Next() {
		return user, errors.GetErr(errors.USER_NOT_FOUND)
	}
	if err = rows.Scan(&user.ID, &_displayname, &_email, &user.CreatedAt, &user.Score, &user.Weight); err != nil {
		return user, err
	}
	user.Displayname = _displayname.String
	user.Email = _email.String
	user.Username = username
	return user, nil
}

func ListUsers(offset int64, limit int64) (users []model.User, err error) {
	if limit <= 0 {
		return make([]model.User, 0), nil
	}
	var (
		rows *sql.Rows
		q    string
	)
	q = `select id, username, display_name, email, created_at, score, weight from users limit ? offset ?`
	if rows, err = db.Query(`pqs`, q, limit, offset); err != nil {
		return nil, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	users = make([]model.User, 0)
	for rows.Next() {
		user := model.User{}
		var _displayname sql.NullString
		var _email sql.NullString
		if err = rows.Scan(&user.ID, &user.Username, &_displayname, &_email, &user.CreatedAt, &user.Score, &user.Weight); err != nil {
			return users, errors.ErrAppend(err, errors.DB_SCAN_ERR)
		}
		user.Displayname = _displayname.String
		user.Email = _email.String
		users = append(users, user)
	}
	return
}

func CreateUser(user model.User) (resUser model.User, err error) {
	q := `insert into users (username, display_name, email, score, weight) values (?, ?, ?, ?, ?)`
	if _, err = db.Exec("pqs", q, user.Username, user.Displayname, user.Email, user.Score, user.Weight); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserID(username string) (id int64, err error) {
	var (
		rows *sql.Rows
		q    = `select id from users where username = ?`
	)
	if rows, err = db.Query(`pqs`, q, username); err != nil {
		return -1, errors.GetErr(errors.UNKNOWN_ERR_INTERNAL)
	}
	defer rows.Close()
	if !rows.Next() {
		return -1, errors.GetErr(errors.USER_NOT_FOUND)
	}
	if err = rows.Scan(&id); err != nil {
		return -1, errors.GetErr(errors.UNKNOWN_ERR_INTERNAL)
	}
	return
}
