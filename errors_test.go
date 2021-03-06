package gocql

import (
	"testing"
)

func TestErrorsParse(t *testing.T) {
	session := createSession(t)
	defer session.Close()

	if err := session.Query(`CREATE TABLE errors_parse (id int primary key)`).Exec(); err != nil {
		t.Fatal("create:", err)
	}

	if err := session.Query(`CREATE TABLE errors_parse (id int primary key)`).Exec(); err == nil {
		t.Fatal("Should have gotten already exists error from cassandra server.")
	} else {
		switch e := err.(type) {
		case RequestErrAlreadyExists:
			if e.Table != "errors_parse" {
				t.Fatal("Failed to parse error response from cassandra for ErrAlreadyExists.")
			}
		default:
			t.Fatal("Failed to parse error response from cassandra for ErrAlreadyExists.")
		}
	}
}
