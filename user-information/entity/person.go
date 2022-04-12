package entity

type Person struct {
	FullName string `json:"full_name,omitempty"`
	Age      uint   `json:"age,omitempty"`
	Address  string `json:"address,omitempty"`
}
