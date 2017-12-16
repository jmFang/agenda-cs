package entity

import (
	"errors"
)

var (
	nullAgumentError = errors.New("[error]: aguments shouldn't be null")
	timeFormatError  = errors.New("[error]: time format should be like yyyy-mm-dd/hh:mm")
	registerError    = errors.New("[error]: username has registered")
	loginError       = errors.New("[error]: login error,check your username or password")
)

// AgendaService ...
type AgendaService struct {
	AgendaStorage *storage
}

// NewAgendaService ...
func NewAgendaService() *AgendaService {
	return &AgendaService{Storage()}
}

// UserLogin ...
// check if the username match password.
func (as *AgendaService) UserLogin(username string, password string) error {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		if username == "" || password == "" {
			return nullAgumentError
		}
		return as.AgendaStorage.Login(username, password)
	}
	return errors.New("Already logined!")
}

//UserLogout ...
// user logout.
func (as *AgendaService) UserLogout() error {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		return errors.New("please login first!")
	}
	return as.AgendaStorage.Logout()
}

// UserRegister ...
// regist a user.
func (as *AgendaService) UserRegister(username string, password string, email string, phone string) (*RetUser, error) {
	if username == "" || password == "" || phone == "" || email == "" {
		return &RetUser{}, nullAgumentError
	}
	user := NewUser(username, password, email, phone)
	return as.AgendaStorage.Register(user)
}

// DeleteUser ...
// delete a user.
func (as *AgendaService) DeleteUser(username string, password string) error {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		return errors.New("please login first!")
	}
	if username == "" || password == "" {
		return nullAgumentError
	}
	return as.AgendaStorage.DeleteUser(username, password)
}

// ListAllUsers ...
// ListAllUsers list all users from storage
// return the list result. RetUser: contains user id
// RetUser: contains user id
func (as *AgendaService) ListAllUsers() ([]RetUser, error) {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		return as.AgendaStorage.UserList, errors.New("please login first!")
	}
	return as.AgendaStorage.ListAllusers()
}

// AddMeeting ...
// add a meeting.
func (as *AgendaService) AddMeeting(title string, startdate string, enddate string, participators []string) (*Meeting, error) {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		return &Meeting{}, errors.New("please login first!")
	}
	if title == "" || startdate == "" || enddate == "" || participators == nil {
		return &Meeting{}, nullAgumentError
	}
	_, err := StringToDate(startdate)
	if err != nil {
		return &Meeting{}, timeFormatError
	}
	_, err1 := StringToDate(enddate)
	if err1 != nil {
		return &Meeting{}, timeFormatError
	}
	sponsor := as.AgendaStorage.CurUser.Username
	meeting := NewMeeting(sponsor, participators, startdate, enddate, title)
	return as.AgendaStorage.addMeeting(meeting)
}

// ListAllMeetings ...
func (as *AgendaService) ListAllMeetings() ([]Meeting, error) {
	as.AgendaStorage.readCurUser()
	if as.AgendaStorage.CurUser == (UserInfo{}) {
		return as.AgendaStorage.MeetingList, errors.New("please login first!")
	}
	return as.AgendaStorage.ListAllMeetings()
}
