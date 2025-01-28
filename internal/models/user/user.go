package voyageraUser

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
