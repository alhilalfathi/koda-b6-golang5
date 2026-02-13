package auth

func (a *Auth) ForgotPass(email, newPassword string) string {
	for i, user := range a.Users {
		if user.Email == email {
			a.Users[i].Password = hash(newPassword)
			return "Password reset successfully."
		}
	}
	return "Email not registered."
}
