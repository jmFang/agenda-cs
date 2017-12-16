package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dzc15331066/agenda-cs/entities"
	"github.com/unrolled/render"
	"net/http"
)

type State struct {
	Success bool   `json:"success"`
	State   string `json:"state"`
}

var (
	inputError = errors.New("Bad input")
)

func postUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var user entities.User
		err := decoder.Decode(&user)
		fmt.Println(user)
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorInfo string }{"Bad Input!"})
			return
		}
		err = entities.UserService.Save(&user)
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, State{false, "register failed"})
			return
		}

		formatter.JSON(w, http.StatusCreated, user)
	}

}

func getAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ulist := entities.UserService.FindAllUser()
		formatter.JSON(w, http.StatusOK, ulist)

	}

}

func userLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var userinfo entities.UserInfo
		decoder.Decode(&userinfo)

		user := entities.UserService.FindByUsernameAndPassword(userinfo.Username, userinfo.Password)
		if user != nil {
			formatter.JSON(w, http.StatusOK, State{true, "login"})
		} else {
			formatter.JSON(w, http.StatusUnauthorized, State{false, "unauthorized"})
		}

	}
}

func userLogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var userinfo entities.UserInfo
		decoder.Decode(&userinfo)
		user := entities.UserService.FindByUsernameAndPassword(userinfo.Username, userinfo.Password)
		if user != nil {
			formatter.JSON(w, http.StatusOK, State{true, "logout"})
		} else {
			formatter.JSON(w, http.StatusUnauthorized, State{false, "unauthorized"})
		}

	}
}

func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var userinfo entities.UserInfo
		decoder.Decode(&userinfo)
		affect, err := entities.UserService.DeleteUser(userinfo.Username, userinfo.Password)

		if err != nil {
			panic(err)
		}
		if affect > 0 {
			formatter.JSON(w, http.StatusOK, State{true, "deleted"})
		} else {
			formatter.JSON(w, http.StatusUnauthorized, State{false, "unauthorized"})
		}
	}
}

func checkInput(req *http.Request, keys ...string) error {
	for _, key := range keys {
		if len(req.Form[key][0]) == 0 {
			return inputError
		}
	}
	return nil
}
