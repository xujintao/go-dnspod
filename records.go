package dnspod

import "net/http"

// RecordsService https://www.dnspod.cn/docs/records.html
type RecordsService struct {
	client *Client
}

// represent all record type
const (
	RecordTypeA           string = "A"
	RecordTypeCNAME       string = "CNAME"
	RecordTypeMX          string = "MX"
	RecordTypeTXT         string = "TXT"
	RecordTypeNS          string = "NS"
	RecordTypeAAAA        string = "AAAA"
	RecordTypeSRV         string = "SRV"
	RecordTypeURLExplicit string = "显性URL"
	RecordTypeURLImplicit string = "显性URL"
)

// Record represents a record of domain
type Record struct {
	ID            string      `json:"id,omitempty"`
	TTL           string      `json:"ttl,omitempty"`
	Value         string      `json:"value,omitempty"`
	Enabled       string      `json:"enabled,omitempty"`
	Status        string      `json:"status,omitempty"`
	UpdatedOn     string      `json:"updated_on,omitempty"`
	Name          string      `json:"name,omitempty"`
	Line          string      `json:"line,omitempty"`
	LineID        string      `json:"line_id,omitempty"`
	Type          string      `json:"type,omitempty"`
	Weight        interface{} `json:"weight,omitempty"`
	MonitorStatus string      `json:"monitor_status,omitempty"`
	Remark        string      `json:"remark,omitempty"`
	UseAqb        string      `json:"use_aqb,omitempty"`
	Mx            string      `json:"mx,omitempty"`
	Hold          string      `json:"hold,omitempty,omitempty"`
}

// ListRecordsOptions represents the available ListRecords() options.
// API docs: https://www.dnspod.cn/docs/records.html#record-list
type ListRecordsOptions struct {
	Domain  string `url:"domain,omitempty"`
	KeyWord string `url:"keyword,omitempty"`
}

// ActionRecordResp represents response body of
// ListRecords, CreateRecord, ModifyRecord, DeleteRecord
// API docs: https://www.dnspod.cn/docs/records.html#record-create
type ActionRecordResp struct {
	Status *Status `json:"status,omitempty"`
	Domain *Domain `json:"domain,omitempty"`
	Info   struct {
		SubDomains  string `json:"sub_domains,omitempty"`
		RecordTotal string `json:"record_total,omitempty"`
		RecordsNum  string `json:"records_num,omitempty"`
	} `json:"info,omitempty"`
	Records []*Record `json:"records,omitempty"`
	Record  *Record   `json:"record,omitempty"`
}

// ListRecords get a list of records
// API docs: https://www.dnspod.cn/docs/records.html#record-list
func (r *RecordsService) ListRecords(opt *ListRecordsOptions) (*ActionRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.List", opt)
	if err != nil {
		return nil, nil, err
	}

	var arr ActionRecordResp
	resp, err := r.client.Do(req, &arr)
	if err != nil {
		return nil, nil, err
	}

	return &arr, resp, nil
}

// CreateRecordOptions represents the available CreateRecord() options.
// API docs: https://www.dnspod.cn/docs/records.html#record-create
type CreateRecordOptions struct {
	Domain     string `url:"domain,omitempty"`
	SubDomain  string `url:"sub_domain,omitempty"`
	RecordType string `url:"record_type,omitempty"`
	RecordLine string `url:"record_line,omitempty"`
	Value      string `url:"value,omitempty"`
}

// CreateRecord create a record.
// API docs: https://www.dnspod.cn/docs/records.html#record-create
func (r *RecordsService) CreateRecord(opt *CreateRecordOptions) (*ActionRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Create", opt)
	if err != nil {
		return nil, nil, err
	}

	var arr ActionRecordResp
	resp, err := r.client.Do(req, &arr)
	if err != nil {
		return nil, nil, err
	}

	return &arr, resp, nil
}

// ModifyRecordOptions represents the available ModifyRecord() options.
// API docs: https://www.dnspod.cn/docs/records.html#record-modify
type ModifyRecordOptions struct {
	Domain     string `url:"domain,omitempty"`
	RecordID   string `url:"record_id,omitempty"`
	SubDomain  string `url:"sub_domain,omitempty"`
	RecordType string `url:"record_type,omitempty"`
	RecordLine string `url:"record_line,omitempty"`
	Value      string `url:"value,omitempty"`
}

// ModifyRecord modify specified record by id.
// API docs: https://www.dnspod.cn/docs/records.html#record-modify
func (r *RecordsService) ModifyRecord(opt *ModifyRecordOptions) (*ActionRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Modify", opt)
	if err != nil {
		return nil, nil, err
	}

	var arr ActionRecordResp
	resp, err := r.client.Do(req, &arr)
	if err != nil {
		return nil, nil, err
	}

	return &arr, resp, nil
}

// DeleteRecordOptions represents the available DeleteRecord() options.
// API docs: https://www.dnspod.cn/docs/records.html#record-remove
type DeleteRecordOptions struct {
	Domain   string `url:"domain,omitempty"`
	RecordID string `url:"record_id,omitempty"`
}

// DeleteRecord delete specified record by id.
// API docs: https://www.dnspod.cn/docs/records.html#record-remove
func (r *RecordsService) DeleteRecord(opt *DeleteRecordOptions) (*ActionRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Remove", opt)
	if err != nil {
		return nil, nil, err
	}

	var arr ActionRecordResp
	resp, err := r.client.Do(req, &arr)
	if err != nil {
		return nil, nil, err
	}

	return &arr, resp, nil
}
