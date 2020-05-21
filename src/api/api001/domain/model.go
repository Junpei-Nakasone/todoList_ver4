package domain

// ResponseData is for response
type ResponseData struct {
	Todos []Todo `json:"todos"`
}

// Todo id containing todo
type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
