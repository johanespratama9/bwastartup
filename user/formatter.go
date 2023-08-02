package user

type UserFormatter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"emai"`
	Token      string `json:"token"`
}

func FormatterUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Id:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}
