package domain

// RequestParam receives id
// type RequestParam struct {
// 	ID int `json:"id"`
// }

// Todo id containing todo
type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateData struct {
	Name string `json:"name"`
}
