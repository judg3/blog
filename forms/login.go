package forms

type Login struct {
	Username string `form:"username" json:"username" required`
	Password string `form:"password" json:"password" required`
}
