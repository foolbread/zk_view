package models

type webUser struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type webUserManager struct {
	Users []*webUser `json:"users"`
}

func (m *webUserManager) CheckUser(usr string, pwd string) bool {
	for _, v := range m.Users {
		if v.User == usr {
			return v.Pwd == pwd
		}
	}

	return false
}
