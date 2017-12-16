package entities

import (
	"time"
)

const format = "2006-01-02/15:04"

type Meeting struct {
	Sponsor       string   `json:"sponsor"`
	Title         string   `json:"title"`
	Start         string   `json:"start"`
	End           string   `json:"end"`
	Participators []string `json:"participators"`
}

func NewMeeting(sponsor string, participator []string, start string, end string, title string) *Meeting {
	participators := make([]string, 0)
	for _, p := range participator {
		participators = append(participators, p)
	}
	return &Meeting{sponsor, title, start, end, participators}
}

func (m *Meeting) ParticipatorIndex(username string) int {
	for i, par := range m.Participators {
		if par == username {
			return i
		}
	}
	return -1
}

// change string "yyyy-mm-dd/hh:mm" to Date.
func StringToDate(str string) (time.Time, error) {
	utc, _ := time.LoadLocation("UTC")
	date, err := time.ParseInLocation(format, str, utc)
	return date, err
}

// change date to string format "yyyy-mm-dd/hh:mm".
func DateToString(date time.Time) string {
	return date.Format(format)
}

// check whether meeting conflicts the dates.
//func (meeting *Meeting) OverLap(start time.Time, end time.Time) bool {
//	return ((meeting.Start.Before(start) || meeting.Start.Equal(start)) && (start.Before(meeting.End) || start.Equal(meeting.End))) || ((start.Before(meeting.Start) || start.Equal(meeting.Start)) && (meeting.Start.Before(end) || meeting.Start.Before(end))) || ((start.Before(meeting.Start) || start.Equal(meeting.Start)) && (end.After(meeting.End) || end.Equal(meeting.End))) || ((meeting.Start.Before(start) || meeting.Start.Equal(start)) && (meeting.End.Equal(end) || meeting.End.After(end)))
//}
