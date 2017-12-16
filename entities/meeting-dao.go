package entities

import (
	"encoding/json"
)

type meetingDao DaoSource

// Save .
func (dao *meetingDao) Save(m *Meeting) error {
	meetingInsertStmt := "insert into meeting(sponsor, title, start, end, participators) values(?,?,?,?,?)"
	stmt, err := dao.Prepare(meetingInsertStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	partsBlob, err := json.Marshal(m.Participators)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(m.Sponsor, m.Title, m.Start, m.End, partsBlob)
	return err
}

// find by sponsor

func (dao *meetingDao) FindBySponsor(sponsor string) []Meeting {
	stmt := "select * from meeting where sponsor = ?"
	return dao.execQueryStmt(stmt)
}

// find all meetings

func (dao *meetingDao) FindAll() []Meeting {
	stmt := "select * from meeting"
	return dao.execQueryStmt(stmt)
}

// find by participators
func (dao *meetingDao) FindByPart(part string) []Meeting {
	mlist := dao.FindAll()
	res := make([]Meeting, 0)
	for _, m := range mlist {
		if m.ParticipatorIndex(part) != -1 {
			res = append(res, m)
		}
	}
	return res
}

func (dao *meetingDao) execQueryStmt(stmt string) []Meeting {
	rows, err := dao.Query(stmt)
	checkErr(err)
	defer rows.Close()
	mlist := make([]Meeting, 0)
	for rows.Next() {
		m := Meeting{}
		var partsBlob []byte
		err := rows.Scan(&m.Sponsor, &m.Title, &m.Start, &m.End, &partsBlob)
		checkErr(err)
		err = json.Unmarshal(partsBlob, &m.Participators)
		checkErr(err)
		mlist = append(mlist, m)

	}
	return mlist
}
