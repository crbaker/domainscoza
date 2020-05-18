package main

import "testing"

func TestShouldFindIndexOfRecord(t *testing.T) {
	records := &Records{
		Records: &[]Record{
			{Content: "", Name: "navigator", Priority: 1, RecordType: "A", Ttl: 1},
			{Content: "", Name: "zipkin", Priority: 1, RecordType: "A", Ttl: 1},
		},
		ReturnCode: 1, Message: "Some Message",
	}

	i := indexOf(*records.Records, func(record Record) bool { return record.Name == "navigator" })
	if i < 0 {
		t.Error("Should find record")
	}

	i = indexOf(*records.Records, func(record Record) bool { return record.Name == "random" })
	if i >= 0 {
		t.Error("Should not find record")
	}
}

func TestShouldFindExistingRecord(t *testing.T) {
	records := &Records{
		Records: &[]Record{
			{Content: "", Name: "navigator", Priority: 1, RecordType: "A", Ttl: 1},
			{Content: "", Name: "zipkin", Priority: 1, RecordType: "A", Ttl: 1},
		},
		ReturnCode: 1, Message: "Some Message",
	}

	recordToMatch := Record{Name: "navigator"}

	if !exists(*records.Records, recordToMatch) {
		t.Error("Should find record")
	}
}
