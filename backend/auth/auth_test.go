package auth

import "testing"

func TestPasswordHash(t *testing.T) {
	pwd := "asdfgfgjh"
	s, err := hashPassword(pwd)
	if err != nil {
		t.Errorf("hashPassword failed")
	}
	b := checkPasswordHash(pwd, s)
	if !b {
		t.Errorf("checkPassword failed")
	}
}
