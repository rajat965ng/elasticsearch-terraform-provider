// resource_server.go
package main

import (
	"elasticsearch/es_ops"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"doc_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"payload": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	id := d.Get("doc_id").(int)
	name := d.Get("name").(string)
	d.SetId(strconv.Itoa(id))

	payload := d.Get("payload").(string)

	err := es_ops.Create(id, name, payload)
	if err != nil {
		panic(err)
	}

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
