package entities

type MeetingAtomicService struct{}

var MeetingService = MeetingAtomicService{}

// Save .
func (*MeetingAtomicService) Save(m *Meeting) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := meetingDao{tx}
	err = dao.Save(m)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	checkErr(err)
	return nil
}

// FindAll .

func (*MeetingAtomicService) FindAll() []Meeting {
	dao := meetingDao{mydb}
	return dao.FindAll()
}

// find by sponsor .

func (*MeetingAtomicService) FindBySponsor(sponsor string) []Meeting {
	dao := meetingDao{mydb}
	return dao.FindBySponsor(sponsor)
}

// find by participators

func (*MeetingAtomicService) FindByPart(part string) []Meeting {
	dao := meetingDao{mydb}
	return dao.FindByPart(part)
}
