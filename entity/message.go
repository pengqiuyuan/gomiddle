package entity

//{"message":"success"} or {"message":"fail"} or {"message":"error"}
type Message struct {
	Message   string      `json:"message"`
}