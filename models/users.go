package models

type UsersModel struct {
	Users []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Age  string `json:"age"`
		Social struct {
			Facebook string `json:"facebook"`
			Twitter string `json:"twitter"`
		} `json:"social"`
	} `json:"users"`
}
