package auth

func (a *Auth) Login(email, password string) string {

	for _, user := range a.Users {
		if user.Email == email {
			if user.Password != hash(password) {
				return "Incorrect password"
			}
			return "Login successful"
		}

	}
	return "Email not registered"
}
