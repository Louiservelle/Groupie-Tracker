package models

type artist struct {
	id      int      `json:"id"`
	image   string   `json:"image"`
	name    string   `json:"name"`
	members []string `json:"members"`
}
