package portal

type Users struct {
	client *Client
}

func (u Users) CreateUser() {}

func (u Users) DeleteUser() {}

func (u Users) InviteUser() {}

func (u Users) UpdateUser() {}

func (u Users) GetUserByID() {}

func (u Users) ListUsers() {}
