package entity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"os"
	"sync"
)

const (
	curUserFilename = "curUser.txt"
)

type storage struct {
	UserList    []RetUser
	MeetingList []Meeting
	Dirty       bool
	CurUser     UserInfo
}

const root_url = "http://localhost:8080"

var (
	s    *storage
	once sync.Once
)

// Storage ...
func Storage() *storage {
	once.Do(func() {
		s = &storage{}
		s.UserList = make([]RetUser, 10)
		s.MeetingList = make([]Meeting, 10)
	})
	return s
}

// read CurUser from file
func (s *storage) readCurUser() error {
	if err := readFromFile(&s.CurUser, curUserFilename); err != nil {
		return err
	} else if s.CurUser == (UserInfo{}) {
		return errors.New("Failed!please login first")
	}
	return nil
}

// write current user to file
func (s *storage) writeCurUser() error {
	if s.CurUser != (UserInfo{}) {
		return writeToFile(s.CurUser, curUserFilename)
	}
	return nil
}

//read json to datalist from file
func readFromFile(datalist interface{}, filename string) error {
	//read data from files
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		return json.Unmarshal(data, &datalist)
	}
	return nil
}

// write datalist to file
func writeToFile(datalist interface{}, filename string) error {
	data, err := json.Marshal(datalist)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0666)
}

// erase current user file while logout.
func (s *storage) eraseCurUser() error {
	file, err := os.OpenFile(curUserFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	s.CurUser = UserInfo{}
	return file.Truncate(0)
}

//set current user
func (s *storage) setCurUser(user UserInfo) error {
	s.CurUser = user
	return s.writeCurUser()
}

//register on backend service
func (s *storage) Register(user *User) (*RetUser, error) {
	if bs, err := json.Marshal(user); err == nil {
		req := bytes.NewBuffer([]byte(bs))
		//os.Stdout.Write(bs)
		contentType := "application/json;charset=utf-8"
		resp, err := http.Post(root_url+"/v1/users?", contentType, req)
		if err != nil {
			return &RetUser{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == 201 {
			body, _ := ioutil.ReadAll(resp.Body)
			var ret RetUser
			//fmt.Println(string(body))
			json.Unmarshal(body, &ret)
			//fmt.Println(ret)
			return &ret, nil
		}
	}
	return &RetUser{}, registerError
}

// login on backend service
func (s *storage) Login(username string, password string) error {
	var info UserInfo
	info.Username = username
	info.Password = password
	bs, err := json.Marshal(info)
	if err == nil {
		req := bytes.NewBuffer([]byte(bs))
		contentType := "application/json;charset=utf-8"
		resp, err := http.Post(root_url+"/v1/user/login?", contentType, req)
		if err != nil {
			//fmt.Println(resp)
			return err
		}

		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			//fmt.Println(body)
			return err
		}
		if resp.StatusCode == 200 {
			s.setCurUser(info)
			//fmt.Println(string(body))
			return nil
		}
	}
	return loginError
}

func (s *storage) Logout() error {
	s.eraseCurUser()
	return nil

	var info UserInfo
	info.Username = s.CurUser.Username
	info.Password = s.CurUser.Password
	bs, err := json.Marshal(info)
	if err == nil {
		req := bytes.NewBuffer([]byte(bs))
		contentType := "application/json;charset=utf-8"
		resp, err := http.Post(root_url+"/v1/user/logout?", contentType, req)
		if err != nil {
			//fmt.Println(resp)
			return err
		}

		//fmt.Println(resp)
		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			//fmt.Println(body)
			return err
		}
		if resp.StatusCode == 201 {
			s.eraseCurUser()
			//fmt.Println(string(body))
			return nil
		}
	}
	return loginError
}

func (s *storage) DeleteUser(username string, password string) error {
	var info UserInfo
	info.Username = username
	info.Password = password
	if s.CurUser.Username != username || s.CurUser.Password != password {
		return errors.New("wrong username or password!")
	}
	bs, err := json.Marshal(info)
	if err != nil {
		return err
	}
	reqBody := bytes.NewBuffer([]byte(bs))
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", root_url+"/v1/users", reqBody)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		//fmt.Println(string(body))
		s.eraseCurUser()
		return nil
	}
	return errors.New("error sending")
}

//query all users and return them as an Array
func (s *storage) ListAllusers() ([]RetUser, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", root_url+"/v1/users?", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(body)
	json.Unmarshal(body, &s.UserList)

	//fmt.Println(s.UserList)
	return s.UserList, err //is this safe ?
}

// add a meeting to meeting list.
func (s *storage) addMeeting(m Meeting) (*Meeting, error) {
	bs, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
		return &Meeting{}, err
	}

	req := bytes.NewBuffer([]byte(bs))
	contentType := "application/json;charset=utf-8"
	resp, err := http.Post(root_url+"/v1/meetings?", contentType, req)
	if err != nil {
		return &Meeting{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return &Meeting{}, err
	}

	if resp.StatusCode == 200 {
		var retMeeting Meeting
		fmt.Println(resp.Status)
		//fmt.Println(string(body))
		json.Unmarshal(body, &retMeeting)
		//fmt.Println(retMeeting)
		return &retMeeting, nil
	}
	//json.Unmarshal(resp_body, s.MeetingList)
	return &Meeting{}, err
}

// Shwo all meetings that have been created in the system
func (s *storage) ListAllMeetings() ([]Meeting, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", root_url+"/v1/meetings", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return s.MeetingList, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return s.MeetingList, err
	}

	fmt.Println(resp.Status)
	//fmt.Println(string(body))

	json.Unmarshal(body, &s.MeetingList)
	return s.MeetingList, err //is this safe ?
}
