package users

func validNewUser(user newUserModel) (errors []string) {
	var invalidityReasons []string

	if len(user.Password) < 8 {
		invalidityReasons = append(invalidityReasons, "Password is too short")
	}

	if len(invalidityReasons) > 0 {
		return invalidityReasons
	}

	return nil
}
