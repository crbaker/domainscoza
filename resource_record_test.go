package main

import "testing"

func TestGetRecords(t *testing.T) {
	got := getRecords("some-key", "wethinkcode", "co.za")
	if got.ReturnCode != 0 {
		t.Errorf("Error")
	}
	if len(*got.Records) != 2 {
		t.Error("Incorrect number of records")
	}
}
