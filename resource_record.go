package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	host := d.Get("host").(string)
	d.SetId(host)

	return resourceRecordRead(d, m)
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRecordRead(d, m)
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"points_to": {
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
