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

type storage struct {
	UserList    []RetUser
	MeetingList []Meeting
	Dirty       bool
	CurUser     User
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

//register on backend service
func (s *storage) Register(user *User) error {
	if bs, err := json.Marshal(user); err == nil {
		req := bytes.NewBuffer([]byte(bs))
		os.Stdout.Write(bs)
		contentType := "application/json;charset=utf-8"
		resp, err := http.Post(root_url+"/v1/users?key=1e3576bt", contentType, req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == 201 {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			return nil
		}
	}
	return registerError
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
		resp, err := http.Post(root_url+"/v1/user/login?key=1e3576bt", contentType, req)
		if err != nil {
			fmt.Println(resp)
			return err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(body)
			return err
		}
		if resp.StatusCode == 200 {
			s.CurUser.Name = username
			s.CurUser.Password = password
			fmt.Println(string(body))
			return nil
		}
	}
	return loginError
}

func (s *storage) Logout() error {
	var info UserInfo
	info.Username = s.CurUser.Name
	info.Password = s.CurUser.Password
	bs, err := json.Marshal(info)
	if err == nil {
		req := bytes.NewBuffer([]byte(bs))
		contentType := "application/json;charset=utf-8"
		resp, err := http.Post(root_url+"/v1/user/logout?key=1e3576bt", contentType, req)
		if err != nil {
			fmt.Println(resp)
			return err
		}

		fmt.Println(resp)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(body)
			return err
		}
		if resp.StatusCode == 201 {
			fmt.Println(string(body))
			return nil
		}
	}
	return loginError
}

func (s *storage) DeleteUser(username string, password string) error {
	var info UserInfo
	info.Username = username
	info.Password = password
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		fmt.Println(string(body))
		s.CurUser.Name = ""
		s.CurUser.Password = ""
		return nil
	}
	return errors.New("error sending")
}

//query all users and return them as an Array
func (s *storage) ListAllusers() ([]RetUser, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", root_url+"/v1/users?key=1e3576bt", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &s.UserList)

	fmt.Println(s.UserList)
	return s.UserList, err //is this safe ?
}

// add a meeting to meeting list.
func (s *storage) addMeeting(m Meeting) error {
	bs, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
		return err
	}

	req := bytes.NewBuffer([]byte(bs))
	contentType := "application/json;charset=utf-8"
	resp, err := http.Post(root_url+"/v1/meetings?key=1e3576bt", contentType, req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if resp.StatusCode == 201 {
		fmt.Println(resp.Status)
		fmt.Println(string(body))
		return nil
	}
	//json.Unmarshal(resp_body, s.MeetingList)
	return err
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
	fmt.Println(string(body))

	json.Unmarshal(body, &s.MeetingList)
	return s.MeetingList, err //is this safe ?
}
