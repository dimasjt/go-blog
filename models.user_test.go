package main

import "testing"

func TestCreateUser(t *testing.T) {
	u, _ = createUser("Mac", "apple@mac.com", "123456")
	lu := userList

	if len(lu) != 2 {
		t.Fail()
	}
}
