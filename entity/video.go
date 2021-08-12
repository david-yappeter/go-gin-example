package entity

type Person struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Location string `json:"location"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=20" validate:"is-cool"`
	Description string `json:"description" binding:"max=10"`
	URL         string `json:"url" binding:"required,url"`
	Author      string `json:"author" binding:"required"`
}
