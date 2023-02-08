package ent

type UserCreateInput struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func (i *UserCreateInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	m.SetPassword(i.Password)
}

func (c *UserCreate) SetInput(i *UserCreateInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

type UserUpdateInput struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (i *UserUpdateInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*i.Name)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*i.Password)
	}
}

func (u *UserUpdate) SetInput(i *UserUpdateInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

func (u *UserUpdateOne) SetInput(i *UserUpdateInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
