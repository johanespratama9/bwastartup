package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"Occupation" binding:"required"`
	Email      string `json:"mail" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
