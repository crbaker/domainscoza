package main

import "testing"

func TestUpdateRecords(t *testing.T) {

	records := []Record{}

	records = append(records, Record{
		Name:       "navigator",
		RecordType: "A",
		Content:    "192.168.78.90",
		Ttl:        300,
		Priority:   1,
	})
	records = append(records, Record{
		Name:       "zipkin",
		RecordType: "A",
		Content:    "192.138.68.90",
		Ttl:        300,
		Priority:   1,
	})

	updateRecords("some-key", "wethinkcode", "co.za", records)

	got := getRecords("some-key", "wethinkcode", "co.za")

	if got.ReturnCode != 0 {
		t.Errorf("Error")
	}
	if len(*got.Records) != 2 {
		t.Error("Incorrect number of records")
	}
}
