package dnspod

// Domain represents domain
type Domain struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Punycode  string   `json:"punycode,omitempty"`
	Grade     string   `json:"grade,omitempty"`
	Owner     string   `json:"owner,omitempty"`
	ExtStatus string   `json:"ext_status,omitempty"`
	TTL       int      `json:"ttl,omitempty"`
	MinTTL    int      `json:"min_ttl,omitempty"`
	DnspodNs  []string `json:"dnspod_ns,omitempty"`
	Status    string   `json:"status,omitempty"`
}
