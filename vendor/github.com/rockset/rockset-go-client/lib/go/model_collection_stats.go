/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package rockset
import (
    "bytes"
    "encoding/json"
    "fmt"
    
)

type CollectionStats struct {
	// number of documents in the collection
	DocCount int64 `json:"doc_count,omitempty"`
	// total collection size in bytes
	TotalSize int64 `json:"total_size,omitempty"`
	// number between 0 and 1 that indicates progress of collection creation
	FillProgress float64 `json:"fill_progress,omitempty"`
	// number of documents purged from the collection
	PurgedDocCount int64 `json:"purged_doc_count,omitempty"`
	// total collection size in bytes purged
	PurgedDocSize int64 `json:"purged_doc_size,omitempty"`
	// milliseconds since Unix epoch Jan 1, 1970
	LastUpdatedMs int64 `json:"last_updated_ms,omitempty"`
	// milliseconds since Unix epoch Jan 1, 1970
	LastQueriedMs int64 `json:"last_queried_ms,omitempty"`
}
func (m CollectionStats) PrintResponse() {
    r, err := json.Marshal(m)
    var out bytes.Buffer
    err = json.Indent(&out, []byte(string(r)), "", "    ")
    if err != nil {
        fmt.Println("error parsing string")
        return
    }

    fmt.Println(out.String())
}

