# go-dnspod

### example

```go
package main

import (
	"log"

	dnspod "github.com/xujintao/go-dnspod"
)

var (
	httpClient *http.Client
	dns *dnspod.Client
)

func main() {
    // http client
    tp := http.DefaultTransport.(*http.Transport)
    httpClient = &http.Client{
        Transport: tp,
        Timeout: 10 * time.Second,
    }

    // dns
    dnsID := "75458"
    dnsToken := "4bf6ee843b92e4cd153eea3acff3f7ca"
    token := fmt.Sprintf("%s,%s", dnsID, dnsToken)
    dns = dnspod.NewClient(httpClient, token)

    // list records
    optListRecords := dnspod.ListRecordsOptions{
        Domain:  "github.com",
        KeyWord: "c1",
    }
    lrr, _, err := dns.Records.ListRecords(&optListRecords)
    if err != nil {
        log.Println(err)
        return nil, fmt.Errorf("dns获取域名记录列表失败")
    }
    
    log.Println(lrr)
}
```