package service

import "testing"

import "FileStorageServer/db"

func TestLoginIn(t *testing.T) {

	_ = db.InitDb()

	user := LoginIn("admin", "1234567")

	if user == nil {
		t.Fatal("user is null")
	}

	t.Log(user)

}
