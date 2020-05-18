package main

type Records struct {
	Records    *[]Record `json:"arrRecords"`
	ReturnCode int       `json:"intReturnCode,omitempty"`
	Message    string    `json:"strMessage,omitempty"`
}

type Record struct {
	Name       string `json:"name,omitempty"`
	RecordType string `json:"type,omitempty"`
	Content    string `json:"content,omitempty"`
	Ttl        int    `json:"ttl,omitempty"`
	Priority   int    `json:"prio,omitempty"`
}

func equals(this *Record, that *Record) bool {
	return this.Name == that.Name
}

func exists(records []Record, this Record) bool {
	return indexOf(records, func(that Record) bool { return that.Name == this.Name }) >= 0
}

func indexOf(records []Record, pred func(record Record) bool) int {
	for i := 0; i < len(records); i++ {
		r := records[i]
		if pred(r) {
			return i
		}
	}
	return -1
}
