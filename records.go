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

// ListRecordsResp represents response body of ListRecord
// API docs: https://www.dnspod.cn/docs/records.html#record-create
type ListRecordsResp struct {
	Status *Status `json:"status,omitempty"`
	Domain *Domain `json:"domain,omitempty"`
	Info   struct {
		SubDomains  string `json:"sub_domains,omitempty"`
		RecordTotal string `json:"record_total,omitempty"`
		RecordsNum  string `json:"records_num,omitempty"`
	} `json:"info,omitempty"`
	Records []*Record `json:"records,omitempty"`
}

// ListRecords get a list of records
// API docs: https://www.dnspod.cn/docs/records.html#record-list
func (r *RecordsService) ListRecords(opt *ListRecordsOptions) (*ListRecordsResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.List", opt)
	if err != nil {
		return nil, nil, err
	}

	var record ListRecordsResp
	resp, err := r.client.Do(req, &record)
	if err != nil {
		return nil, nil, err
	}

	return &record, resp, nil
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

// CreateRecordResp represents response body of CreateRecord
// API docs: https://www.dnspod.cn/docs/records.html#record-create
type CreateRecordResp struct {
	Status *Status `json:"status,omitempty"`
	Record *Record `json:"record,omitempty"`
}

// CreateRecord create a record.
// API docs: https://www.dnspod.cn/docs/records.html#record-create
func (r *RecordsService) CreateRecord(opt *CreateRecordOptions) (*CreateRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Create", opt)
	if err != nil {
		return nil, nil, err
	}

	var crr CreateRecordResp
	resp, err := r.client.Do(req, &crr)
	if err != nil {
		return nil, nil, err
	}

	return &crr, resp, nil
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

// ModifyRecordResp represents response body of ModifyRecord
// API docs: https://www.dnspod.cn/docs/records.html#record-modify
type ModifyRecordResp struct {
	Status *Status `json:"status,omitempty"`
	Record *Record `json:"record,omitempty"`
}

// ModifyRecord modify specified record by id.
// API docs: https://www.dnspod.cn/docs/records.html#record-modify
func (r *RecordsService) ModifyRecord(opt *ModifyRecordOptions) (*ModifyRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Modify", opt)
	if err != nil {
		return nil, nil, err
	}

	var mrr ModifyRecordResp
	resp, err := r.client.Do(req, &mrr)
	if err != nil {
		return nil, nil, err
	}

	return &mrr, resp, nil
}

// DeleteRecordOptions represents the available DeleteRecord() options.
// API docs: https://www.dnspod.cn/docs/records.html#record-remove
type DeleteRecordOptions struct {
	Domain   string `url:"domain,omitempty"`
	RecordID string `url:"record_id,omitempty"`
}

// DeleteRecordResp represents response body of DeleteRecord
// API docs: https://www.dnspod.cn/docs/records.html#record-remove
type DeleteRecordResp struct {
	Status *Status `json:"status,omitempty"`
	Record *Record `json:"record,omitempty"`
}

// DeleteRecord delete specified record by id.
// API docs: https://www.dnspod.cn/docs/records.html#record-remove
func (r *RecordsService) DeleteRecord(opt *DeleteRecordOptions) (*DeleteRecordResp, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "Record.Delete", opt)
	if err != nil {
		return nil, nil, err
	}

	var drr DeleteRecordResp
	resp, err := r.client.Do(req, &drr)
	if err != nil {
		return nil, nil, err
	}

	return &drr, resp, nil
}
