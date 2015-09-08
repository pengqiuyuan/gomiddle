package entity

type ResponseList struct {
	Choose   int      `json:"choose"`
	Success  int      `json:"success"`
	ObjFail  []string `json:"objFail"`
	Fail     int      `json:"fail"`
}