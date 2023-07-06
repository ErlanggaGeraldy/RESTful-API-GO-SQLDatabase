package request

type AuthorRequest struct {
	Name    string `json:"name" validate:"required"`
	Country string `json:"country" validate:"required"`
}
type AuthorRequestUpdate struct {
	Name    string `json:"name" `
	Country string `json:"country" `
}
