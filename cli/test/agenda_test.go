package test

import (
	"io/ioutil"
	"testing"

	"github.com/dzc15331066/agenda-cs/cli/entity"
)

func TestUserRegister(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	email := "1235@qq.com"
	phone := "1315766578"
	retUser, err := as.UserRegister(name, pass, email, phone)
	if err != nil {
		t.Error(err)
	} else {
		res := (retUser.Name == name && retUser.Password == pass && retUser.Email == email && retUser.Phone == phone)
		want := `[{"Name":"Username","Password":"pass","Email":"1235@qq.com","Phone":"1315766578"}]`
		if !res {
			t.Errorf("want %q but get %q", want, retUser)
		}
	}

}

func TestUserLogin(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	as.UserLogout()
	err := as.UserLogin(name, pass)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := `{"username":"Username","password":"pass"}`
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestUserLogout(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.UserLogout()
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := ""
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestListAllUsers(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.UserLogin("Username", "pass")
	users, err := as.ListAllUsers()
	eq := true
	if err != nil {
		t.Error(err)
	}
	want := entity.User{"Username", "pass", "1235@qq.com", "1315766578"}
	res := make([]entity.User, 0)
	res = append(res, want)
	equ := (len(users) == len(res))
	if !equ {
		t.Errorf("want %q but get %q", want, users)
	}
	for i := 0; i < len(users); i++ {
		if users[i].Name == want.Name {
			eq = (users[i].Password == want.Password && users[i].Email == want.Email && users[i].Phone == want.Phone)
			break
		}
	}
	if !eq {
		t.Errorf("want %q but get %q", want, users)
	}
}

func TestDeleteUser(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	err := as.DeleteUser(name, pass)
	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := ""
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestAddMeeting(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	email := "1235@qq.com"
	phone := "1315766578"
	as.UserRegister(name, pass, email, phone)
	as.UserRegister("part1", "pass1", "12143@qq.com", "13443443535")
	as.UserRegister("part2", "pass2", "121434@qq.com", "1334345345")
	as.UserLogin("Username", "pass")
	title := "meeting"
	startdate := "2001-11-11/12:00"
	enddate := "2005-12-11/13:00"
	participators := []string{"part1", "part2"}
	retMeeting, err := as.AddMeeting(title, startdate, enddate, participators)
	if err != nil {
		t.Error(err)
	} else {
		want := entity.Meeting{"Username", title, startdate, enddate, participators}
		res := equalMeeting(*retMeeting, want)
		if !res {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestListAllMeetings(t *testing.T) {
	as := entity.NewAgendaService()
	ms, err := as.ListAllMeetings()
	if err != nil {
		t.Error(err)
	}
	start := "2001-11-11/12:00"
	end := "2005-12-11/13:00"
	want := make([]entity.Meeting, 0)
	m := entity.Meeting{"Username", "meeting", start, end, []string{"part1", "part2"}}
	want = append(want, m)
	l := min(len(want), len(ms))
	for i := 0; i < l; i++ {
		if !equalMeeting(ms[i], want[i]) {
			t.Errorf("want %q but get %q", want, ms)
		}
	}
}

func equalMeeting(m1 entity.Meeting, m2 entity.Meeting) bool {
	eq := m1.Sponsor == m2.Sponsor && m1.Title == m2.Title && m1.Start == m2.Start && m1.End == m2.End
	l := min(len(m1.Participators), len(m2.Participators))
	for i := 0; i < l; i++ {
		if m1.Participators[i] != m2.Participators[i] {
			return false
		}
	}
	return eq
}

func min(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
