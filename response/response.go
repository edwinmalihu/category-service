package response

type ResponseSuccess struct {
	Category string `json:"category"`
	Msg      string `json:"msg"`
}

type ResponseCategory struct {
	Category string `json:"category"`
}
