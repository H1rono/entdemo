package repository

import (
	"context"
	"time"

	"github.com/H1rono/entdemo/ent"
)

type CreateCar struct {
	Model        string
	RegisteredAt time.Time
}

type Car struct {
	ID           int
	Model        string
	RegisteredAt time.Time
}

func fromEntCar(car *ent.Car) *Car {
	return &Car{
		ID:           car.ID,
		Model:        car.Model,
		RegisteredAt: car.RegisteredAt,
	}
}

func (r *Repository) CreateCar(ctx context.Context, car *CreateCar) (*Car, error) {
	res, err := r.c.Car.
		Create().
		SetModel(car.Model).
		SetRegisteredAt(car.RegisteredAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntCar(res), nil
}

func (r *Repository) GetCars(ctx context.Context) ([]*Car, error) {
	res, err := r.c.Car.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	cars := make([]*Car, 0, len(res))
	for _, car := range res {
		cars = append(cars, fromEntCar(car))
	}
	return cars, nil
}

func (r *Repository) GetCar(ctx context.Context, id int) (*Car, error) {
	res, err := r.c.Car.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return fromEntCar(res), nil
}

func (r *Repository) UpdateCar(ctx context.Context, id int, car *CreateCar) (*Car, error) {
	res, err := r.c.Car.UpdateOneID(id).
		SetModel(car.Model).
		SetRegisteredAt(car.RegisteredAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntCar(res), nil
}

func (r *Repository) DeleteCar(ctx context.Context, id int) error {
	return r.c.Car.DeleteOneID(id).Exec(ctx)
}
