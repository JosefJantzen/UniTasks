package database

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id"`
	EMail string    `json:"eMail"`
	Pwd   string    `json:"pwd"`
}

func (user User) empty() bool {
	if user.Id.String() == "" || user.EMail == "" || user.Pwd == "" {
		return true
	}
	return false
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

func (s *DBService) CheckMailUsed(m string) bool {
	res, err := s.db.Query("SELECT id FROM users WHERE e_mail=$1", m)
	if err != nil {
		return false
	}

	defer res.Close()
	res.Next()
	var uid uuid.UUID

	if err := res.Scan(&uid); err != nil {
		return false
	}
	return true
}

func (s *DBService) InsertUser(email string, pwd string) uuid.UUID {
	if email == "" || pwd == "" {
		return uuid.Nil
	}
	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO users (e_mail, pwd) VALUES ($1, $2) RETURNING id",
				email,
				pwd,
			).Scan(&id)

			if err != nil {
				return err
			}
			return nil
		})

	if err != nil {
		return uuid.Nil
	}

	return id
}

func (s *DBService) DeleteUser(id uuid.UUID) error {
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"DELETE FROM recurring_tasks WHERE parentUser=$1;",
				id,
			)
			if err != nil {
				return err
			}
			_, err = tx.Exec(
				"DELETE FROM tasks WHERE parentUser=$1",
				id,
			)
			if err != nil {
				return err
			}
			_, err = tx.Exec(
				"DELETE FROM users WHERE id=$1",
				id,
			)
			return err
		})
}
