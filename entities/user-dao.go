package entities

import (
	"database/sql"
	"fmt"
)

type userDao DaoSource

var userInsertStmt = "insert into user(username,password,email,phone) values(?,?,?,?)"

// Save .
func (dao *userDao) Save(u *User) error {
	stmt, err := dao.Prepare(userInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(NewNullString(u.Name), NewNullString(u.Password), NewNullString(u.Email), NewNullString(u.Phone))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.UID = int(id)
	return nil
}

var userQueryAll = "select * from user"
var userQueryByInfo = "select * from user where username = ? and password = ?"

func (dao *userDao) FindAllUser() []User {
	rows, err := dao.Query(userQueryAll)
	checkErr(err)
	defer rows.Close()

	ulist := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.UID, &u.Name, &u.Password, &u.Email, &u.Phone)
		checkErr(err)
		ulist = append(ulist, u)
	}
	return ulist

}

// Find by username and password .
func (dao *userDao) FindByUsernameAndPassword(username string, password string) *User {
	stmt, err := dao.Prepare(userQueryByInfo)
	checkErr(err)

	defer stmt.Close()

	row := stmt.QueryRow(username, password)
	u := User{}
	err = row.Scan(&u.UID, &u.Name, &u.Password, &u.Email, &u.Phone)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		checkErr(err)
	}

	return &u
}

// delete a user

func (dao *userDao) DeleteUser(username string, password string) (int, error) {
	stmt, err := dao.Prepare("delete from user where username = ? and password = ?")
	checkErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(username, password)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	return int(affect), err
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
