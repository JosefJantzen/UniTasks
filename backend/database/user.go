package database

import (
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id"`
	EMail string    `json:"eMail"`
	Pwd   string    `json:"pwd"`
}

func (s *DBService) GetUserById(id uuid.UUID) *User {
	res, err := s.db.Query("SELECT * FROM users WHERE id=$1", id.String())
	if err != nil {
		return nil
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var mail string
	var pwd string

	if err := res.Scan(&uid, &mail, &pwd); err != nil {
		return nil
	}
	user := User{Id: uid, EMail: mail, Pwd: pwd}
	return &user
}

func (s *DBService) GetUserByMail(m string) *User {
	res, err := s.db.Query("SELECT * FROM users WHERE e_mail=$1", m)
	if err != nil {
		return nil
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var mail string
	var pwd string

	if err := res.Scan(&uid, &mail, &pwd); err != nil {
		return nil
	}
	user := User{Id: uid, EMail: mail, Pwd: pwd}
	return &user
}
