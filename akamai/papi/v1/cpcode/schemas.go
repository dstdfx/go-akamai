package cpcode

import "time"

// CPCode represents content provider code which is used to track all
// web traffic handled by Akamai servers.
type CPCode struct {
	ID string `json:"cpcodeId"`
	Name string `json:"cpcodeName"`
	ProductIDs []string `json:"productIds"`
	CreatedDate *time.Time `json:"createdDate"`
}
