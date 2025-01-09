package repository

import (
	"context"
	"github.com/H1rono/entdemo/ent"
)

type UserAndCars struct {
	User
	Cars []*Car
}

func (r *Repository) GetUserCars(ctx context.Context, userID int) ([]*Car, error) {
	user, err := r.c.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	cars, err := user.QueryCars().All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*Car, 0, len(cars))
	for _, c := range cars {
		res = append(res, fromEntCar(c))
	}
	return res, nil
}

func (r *Repository) GetUserAndCars(ctx context.Context, userID int) (*UserAndCars, error) {
	user, err := r.c.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	cars, err := user.QueryCars().All(ctx)
	if err != nil {
		return nil, err
	}
	resCars := make([]*Car, 0, len(cars))
	for _, c := range cars {
		resCars = append(resCars, fromEntCar(c))
	}
	resUser := fromEntUser(user)
	userAndCars := &UserAndCars{
		User: *resUser,
		Cars: resCars,
	}
	return userAndCars, nil
}

func (r *Repository) UpdateUserCars(ctx context.Context, userID int, cars []*Car) ([]*Car, error) {
	// FIXME: transaction
	user, err := r.c.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	entCars := make([]*ent.Car, 0, len(cars))
	for _, c := range cars {
		entCar, err := r.c.Car.Get(ctx, c.ID)
		if err != nil {
			return nil, err
		}
		entCars = append(entCars, entCar)
	}

	user, err = user.Update().ClearCars().Save(ctx)
	if err != nil {
		return nil, err
	}
	user, err = user.Update().AddCars(entCars...).Save(ctx)
	if err != nil {
		return nil, err
	}
	resEntCars, err := user.QueryCars().All(ctx)
	if err != nil {
		return nil, err
	}
	resCars := make([]*Car, 0, len(resEntCars))
	for _, c := range resEntCars {
		resCars = append(resCars, fromEntCar(c))
	}
	return resCars, nil
}

func (r *Repository) UnlinkUserCars(ctx context.Context, userID int, cars []*Car) ([]*Car, error) {
	// FIXME: transaction
	user, err := r.c.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	entCars := make([]*ent.Car, 0, len(cars))
	for _, c := range cars {
		entCar, err := r.c.Car.Get(ctx, c.ID)
		if err != nil {
			return nil, err
		}
		entCars = append(entCars, entCar)
	}

	user, err = user.Update().RemoveCars(entCars...).Save(ctx)
	if err != nil {
		return nil, err
	}
	resEntCars, err := user.QueryCars().All(ctx)
	if err != nil {
		return nil, err
	}
	resCars := make([]*Car, 0, len(resEntCars))
	for _, c := range resEntCars {
		resCars = append(resCars, fromEntCar(c))
	}
	return resCars, nil
}
