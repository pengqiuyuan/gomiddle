package entity

type ResponseList struct {
	Choose   string   `json:"choose"`
	Success  string   `json:"success"`
	ObjFail  []string `json:"objFail"`
	Fail     string   `json:"fail"`
}