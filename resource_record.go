package main

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func recordFromResourceData(d *schema.ResourceData) *Record {
	name := d.Get("name").(string)
	content := d.Get("content").(string)
	recordType := d.Get("type").(string)
	priority := d.Get("priority").(int)
	ttl := d.Get("ttl").(int)

	return &Record{Content: content, Name: name, Priority: priority, RecordType: recordType, Ttl: ttl}
}

func getCurrentRecords(d *schema.ResourceData) *Records {
	name := d.Get("name").(string)
	log.Println("Reading Records " + name)

	key := d.Get("key").(string)
	tld := d.Get("tld").(string)
	sld := d.Get("sld").(string)

	return getRecords(key, tld, sld)
}

func saveNewRecords(d *schema.ResourceData, records []Record) {

	log.Println("Sending " + strconv.Itoa(len(records)) + " records")

	key := d.Get("key").(string)
	tld := d.Get("tld").(string)
	sld := d.Get("sld").(string)

	updateRecords(key, tld, sld, records)
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	records := getCurrentRecords(d)

	log.Println("Found " + strconv.Itoa(len(*records.Records)) + " Records ")

	recordToAdd := recordFromResourceData(d)

	if !exists(*records.Records, *recordToAdd) {

		log.Println("Creating Record")
		log.Println(recordToAdd.Name)

		updatedRecords := append(*records.Records, *recordToAdd)
		saveNewRecords(d, updatedRecords)
	}

	d.SetId(recordToAdd.Name)

	return resourceRecordRead(d, m)
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	records := getCurrentRecords(d)
	recordToRead := recordFromResourceData(d)
	if !exists(*records.Records, *recordToRead) {
		d.SetId("")
		return nil
	}
	d.SetId(recordToRead.Name)
	return nil
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
	records := getCurrentRecords(d)
	recordToUpdate := recordFromResourceData(d)

	existingRecord := find(*records.Records, *recordToUpdate)

	if existingRecord != nil {

		existingRecord.Name = recordToUpdate.Name
		existingRecord.RecordType = recordToUpdate.RecordType
		existingRecord.Content = recordToUpdate.Content
		existingRecord.Ttl = recordToUpdate.Ttl
		existingRecord.Priority = recordToUpdate.Priority

		saveNewRecords(d, *records.Records)
	}

	d.SetId(recordToUpdate.Name)

	return nil
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	records := getCurrentRecords(d)
	recordToDelete := recordFromResourceData(d)

	updatedRecords := remove(*records.Records, *recordToDelete)
	saveNewRecords(d, updatedRecords)

	return nil
}

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sld": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tld": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
