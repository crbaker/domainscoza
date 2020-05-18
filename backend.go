package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func doPost(url string, data *url.Values) (resp *http.Response, err error) {
	return http.Post(url, "application/x-www-form-urlencoded", strings.NewReader((data.Encode())))
}

func getRecords(key string, sld string, tld string) *Records {

	data := url.Values{
		"key": {key},
		"sld": {sld},
		"tld": {tld},
	}

	r, _ := doPost("http://localhost:5000/records", &data)

	bytes, _ := ioutil.ReadAll(r.Body)

	response := &Records{
		Records: &[]Record{},
	}

	json.Unmarshal(bytes, &response)

	return response
}
