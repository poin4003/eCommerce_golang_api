package vo

type UserRegistratorRequest struct {
	Email   string `json:"email" biding:"required,email"`
	Purpose string `json:"purpose" biding:"required,purpose"`
}
