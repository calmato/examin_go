package request

type Auth struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
