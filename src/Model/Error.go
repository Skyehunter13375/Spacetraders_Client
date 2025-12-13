package Model

type ApiError struct {
	Err  string `json:"error"`
	Msg  string `json:"message"`
	Code int    `json:"code"`
}
