package presenter

type BasicResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}
