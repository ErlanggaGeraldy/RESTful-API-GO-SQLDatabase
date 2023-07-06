package request

type BookRequest struct {
	Title     string `json:"title" validate:"required"`
	Published string `json:"published" validate:"required"`
	Isbn      string `json:"isbn" validate:"required"`
}
type BookRequestUpdate struct {
	Title     string `json:"title" validate:"required"`
	Published string `json:"published" validate:"required"`
	Isbn      string `json:"isbn" validate:"required"`
}