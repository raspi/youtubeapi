package shared

import "fmt"

type APIError struct {
	Error APIErrorError `json:"error"`
}

var _ error = APIErrorError{}

type APIErrorError struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorElement `json:"errors"`
	Status  string         `json:"status"`
	Details []Detail       `json:"details"`
}

func (A APIErrorError) Error() string {
	// TODO add more details
	return fmt.Sprintf(`%v %v: %v`, A.Code, A.Status, A.Message)
}

type Detail struct {
	Type     string   `json:"@type"`
	Reason   string   `json:"reason"`
	Domain   string   `json:"domain"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Service string `json:"service"`
}

type ErrorElement struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}
