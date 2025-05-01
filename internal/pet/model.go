package pet

type Pet struct {
	Ascii       string `json:"ascii"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
