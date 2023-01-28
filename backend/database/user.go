package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	EMail     string    `json:"eMail"`
	Pwd       string    `json:"pwd"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *DBService) GetUserById(id uuid.UUID) (*User, error) {
	res, err := s.db.Query("SELECT * FROM users WHERE id=$1", id.String())
	if err != nil {
		return nil, err
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var mail string
	var pwd string
	var createdAt time.Time
	var updatedAt time.Time

	if err := res.Scan(&uid, &mail, &pwd, &createdAt, &updatedAt); err != nil {
		return nil, err
	}
	user := User{Id: uid, EMail: mail, Pwd: pwd, CreatedAt: createdAt, UpdatedAt: updatedAt}
	return &user, nil
}

func (s *DBService) GetUserByMail(m string) (*User, error) {
	res, err := s.db.Query("SELECT * FROM users WHERE e_mail=$1", m)
	if err != nil {
		return nil, err
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var mail string
	var pwd string
	var createdAt time.Time
	var updatedAt time.Time

	if err := res.Scan(&uid, &mail, &pwd, &createdAt, &updatedAt); err != nil {
		return nil, err
	}
	user := User{Id: uid, EMail: mail, Pwd: pwd, CreatedAt: createdAt, UpdatedAt: updatedAt}
	return &user, nil
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

func (s *DBService) InsertUser(email string, pwd string) (uuid.UUID, error) {
	if email == "" || pwd == "" {
		return uuid.Nil, errors.New("InsertUser: e-Mail or password is empty")
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
	return id, err
}

func (s *DBService) UpdateMail(id uuid.UUID, mail string) error {
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE users SET e_mail = $1 WHERE id = $2",
				mail,
				id,
			)
			return err
		})
}

func (s *DBService) UpdatePwd(id uuid.UUID, pwd string) error {
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE users SET pwd = $1 WHERE id = $2",
				pwd,
				id,
			)
			return err
		})
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
