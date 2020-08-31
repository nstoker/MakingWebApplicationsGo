package viewmodel

// Base structure for viewmodel
type Base struct {
	Title string
}

// NewBase creates a new base model
func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
