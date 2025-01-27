package api

type ArrayResponse struct {
	Items interface{} `json:"items"`
	Count int         `json:"count"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
