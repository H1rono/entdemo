package repository

import "context"

type CreateUser struct {
	Age  int
	Name string
}

type User struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func (r *Repository) CreateUser(ctx context.Context, u *CreateUser) (*User, error) {
	res, err := r.c.User.Create().
		SetAge(u.Age).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:   res.ID,
		Age:  res.Age,
		Name: res.Name,
	}, nil
}
