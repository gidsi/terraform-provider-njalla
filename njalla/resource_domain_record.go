package njalla

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

type AddRecordResponseResult struct {
	ID   	int64     `json:"id"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Content	string	`json:"content"`
	TTL		int		`json:"ttl"`
}
type AddRecordResponse struct {
	Result AddRecordResponseResult `json:"result"`
}


type ListRecordsResponse struct {
	Result	ListRecordsResponseResult `json:"result"`
}
type ListRecordsResponseResult struct {
	Records	[]ListRecordsResponseRecords `json:"records"`
}
type ListRecordsResponseRecords struct {
	ID   	int64     `json:"id"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Content	string	`json:"content"`
	TTL		int		`json:"ttl"`
}

type AddRecordRequest struct {
	Domain	string	`json:"domain"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Content	string	`json:"content"`
	TTL		int		`json:"ttl"`
}
type RemoveRecordRequest struct {
	Domain	string	`json:"domain"`
	ID		int64	`json:"id"`
}
type EditRecordRequest struct {
	Domain	string	`json:"domain"`
	ID		int64	`json:"id"`
	Content	string	`json:"content"`
}
type ListRecordsRequest struct {
	Domain string	`json:"domain"`
}

func resourceDomainRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainRecordCreate,
		Read:   resourceDomainRecordRead,
		Update: resourceDomainRecordUpdate,
		Delete: resourceDomainRecordDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "A",
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  10800,
			},
		},
	}
}

func resourceDomainRecordCreate(d *schema.ResourceData, m interface{}) error {
	njallaClient := m.(NjallaClient)
	values := AddRecordRequest{
		Domain: d.Get("domain").(string),
		Name: d.Get("name").(string),
		Content: d.Get("content").(string),
		Type: d.Get("type").(string),
		TTL: d.Get("ttl").(int),
	}

	var result AddRecordResponse
	err := njallaClient.DoRequest("add-record", values, &result)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(result.Result.ID, 10))
	err = d.Set("name", result.Result.Name)
	if err != nil {
		return err
	}
	err = d.Set("type", result.Result.Type)
	if err != nil {
		//return errors.New("test2")
		return err
	}
	err = d.Set("content", result.Result.Content)
	if err != nil {
		return err
	}
	err = d.Set("ttl", result.Result.TTL)
	if err != nil {
		return err
	}

	return nil
}
func resourceDomainRecordRead(d *schema.ResourceData, m interface{}) error {
	njallaClient := m.(NjallaClient)
	param := ListRecordsRequest {
		Domain: d.Get("domain").(string),
	}

	var result ListRecordsResponse
	err := njallaClient.DoRequest("list-records", param, &result)
	if err != nil {
		return err
	}

	for _, value := range result.Result.Records {
		if strconv.FormatInt(value.ID, 10) == d.Id() {
			err = d.Set("name", value.Name)
			if err != nil {
				return err
			}
			err = d.Set("type", value.Type)
			if err != nil {
				//return errors.New("test2")
				return err
			}
			err = d.Set("content", value.Content)
			if err != nil {
				return err
			}
			err = d.Set("ttl", value.TTL)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}
func resourceDomainRecordUpdate(d *schema.ResourceData, m interface{}) error {
	njallaClient := m.(NjallaClient)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	param := EditRecordRequest{
		Domain:  d.Get("domain").(string),
		ID:      id,
		Content: d.Get("content").(string),
	}

	err = njallaClient.DoRequest("edit-record", param, nil)
	if err != nil {
		return err
	}

	return nil
}
func resourceDomainRecordDelete(d *schema.ResourceData, m interface{}) error {
	njallaClient := m.(NjallaClient)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	values := RemoveRecordRequest{
		Domain: d.Get("domain").(string),
		ID: id,
	}

	err = njallaClient.DoRequest("remove-record", values, nil)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

