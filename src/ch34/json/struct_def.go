package json

type BasicInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type JobInfo struct {
	Skills   []string `json:"skills"`
	Company  string   `json:"company"`
	Position string   `json:"position"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
