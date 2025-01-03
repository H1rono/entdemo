package repository

import (
	"context"

	"github.com/H1rono/entdemo/ent"
)

type CreateUser struct {
	Age  int
	Name string
}

type UpdateUser = CreateUser

type User struct {
	ID   int
	Age  int
	Name string
}

func fromEntUser(u *ent.User) *User {
	return &User{
		ID:   u.ID,
		Age:  u.Age,
		Name: u.Name,
	}
}

func (r *Repository) CreateUser(ctx context.Context, u *CreateUser) (*User, error) {
	res, err := r.c.User.Create().
		SetAge(u.Age).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntUser(res), nil
}

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	users, err := r.c.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*User, 0, len(users))
	for _, u := range users {
		res = append(res, fromEntUser(u))
	}
	return res, nil
}

func (r *Repository) GetUser(ctx context.Context, id int) (*User, error) {
	u, err := r.c.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return fromEntUser(u), nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	return r.c.User.DeleteOneID(id).Exec(ctx)
}

func (r *Repository) UpdateUser(ctx context.Context, id int, u *UpdateUser) (*User, error) {
	res, err := r.c.User.UpdateOneID(id).
		SetAge(u.Age).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntUser(res), nil
}
