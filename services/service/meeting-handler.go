package service

import (
	"encoding/json"
	"fmt"
	"github.com/dzc15331066/agenda-cs/entities"

	"github.com/unrolled/render"
	"net/http"
)

func postMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var m entities.Meeting
		err := decoder.Decode(&m)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		fmt.Println(m)
		err = entities.MeetingService.Save(&m)
		if err != nil {
			fmt.Println(err)
			formatter.JSON(w, http.StatusBadRequest, nil)
		} else {
			formatter.JSON(w, http.StatusOK, m)
		}

	}
}

func getAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		mlist := entities.MeetingService.FindAll()
		formatter.JSON(w, http.StatusOK, mlist)
	}

}
