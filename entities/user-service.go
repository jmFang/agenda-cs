package entities

type UserAtomicService struct{}

var UserService = UserAtomicService{}

// Save .
func (*UserAtomicService) Save(u *User) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userDao{tx}
	err = dao.Save(u)
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return err
}

// FindAllUser
func (*UserAtomicService) FindAllUser() []User {
	dao := userDao{mydb}
	return dao.FindAllUser()
}

// Find by username and password .
func (*UserAtomicService) FindByUsernameAndPassword(username string, password string) *User {
	dao := userDao{mydb}
	return dao.FindByUsernameAndPassword(username, password)
}

// delete user

func (*UserAtomicService) DeleteUser(username string, password string) (int, error) {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userDao{tx}

	affect, err := dao.DeleteUser(username, password)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return affect, err
}
