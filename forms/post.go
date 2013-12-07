package forms

type Post struct {
	Title   string `form:"title" json:"title" required`
	Content string `form:"content" json:"content" required`
}
