package model

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func UserFromGitlabUser(glu *GitLabUser) *User {
	return &User{
		UserName: glu.Username,
		Email:    glu.Email,
		Name:     glu.Name,
	}
}

func UserFromGithubUser(ghu *GithubUser) *User {
	return &User{
		UserName: ghu.Email,
		Email:    ghu.Email,
		Name:     ghu.Name,
	}
}

func UserFromGoogleUser(ggu *GoogleUser) *User {
	user := &User{}
	user.Name = ggu.DisplayName

	if len(ggu.Emails) > 0 {
		user.Email = ggu.Emails[0].Value
		user.UserName = ggu.Emails[0].Value
	}

	return user
}
