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
	if username == "" || password == "" {
		return nullAgumentError
	}
	return as.AgendaStorage.Login(username, password)
}

//UserLogout ...
// user logout.
func (as *AgendaService) UserLogout() error {
	return as.AgendaStorage.Logout()
}

// UserRegister ...
// regist a user.
func (as *AgendaService) UserRegister(username string, password string, email string, phone string) error {
	if username == "" || password == "" || phone == "" || email == "" {
		return nullAgumentError
	}
	user := NewUser(username, password, email, phone)
	return as.AgendaStorage.Register(user)
}

// DeleteUser ...
// delete a user.
func (as *AgendaService) DeleteUser(username string, password string) error {
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
	return as.AgendaStorage.ListAllusers()
}

// AddMeeting ...
// add a meeting.
func (as *AgendaService) AddMeeting(title string, startdate string, enddate string, participators []string) error {
	if title == "" || startdate == "" || enddate == "" || participators == nil {
		return nullAgumentError
	}
	_, err := StringToDate(startdate)
	if err != nil {
		return timeFormatError
	}
	_, err1 := StringToDate(enddate)
	if err1 != nil {
		return timeFormatError
	}
	sponsor := as.AgendaStorage.CurUser.Name
	meeting := NewMeeting(sponsor, participators, startdate, enddate, title)
	return as.AgendaStorage.addMeeting(meeting)
}

// ListAllMeetings ...
func (as *AgendaService) ListAllMeetings() ([]Meeting, error) {
	return as.AgendaStorage.ListAllMeetings()
}
