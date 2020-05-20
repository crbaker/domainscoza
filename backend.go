package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func doPost(url string, data *url.Values) (resp *http.Response, err error) {
	return http.Post(url, "application/x-www-form-urlencoded", strings.NewReader((data.Encode())))
}

func updateRecords(key string, sld string, tld string, records []Record) {

	data := url.Values{
		"key": {key},
		"sld": {sld},
		"tld": {tld},
	}

	for i := 0; i < len(records); i++ {
		record := records[i]

		data.Set("name"+strconv.Itoa(i+1), record.Name)
		data.Set("type"+strconv.Itoa(i+1), record.RecordType)
		data.Set("content"+strconv.Itoa(i+1), record.Content)
		data.Set("ttl"+strconv.Itoa(i+1), strconv.Itoa(record.Ttl))
		data.Set("prio"+strconv.Itoa(i+1), strconv.Itoa(record.Priority))
	}

	doPost("http://localhost:5000/update-records", &data)

}

func getRecords(key string, sld string, tld string) *Records {

	data := url.Values{
		"key": {key},
		"sld": {sld},
		"tld": {tld},
	}

	r, _ := doPost("http://localhost:5000/get-records", &data)

	bytes, _ := ioutil.ReadAll(r.Body)

	response := &Records{
		Records: &[]Record{},
	}

	json.Unmarshal(bytes, &response)

	return response
}
