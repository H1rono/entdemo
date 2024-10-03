package repository

import (
	"context"
	"time"
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

func (r *Repository) CreateCar(ctx context.Context, car *CreateCar) (*Car, error) {
	res, err := r.c.Car.
		Create().
		SetModel(car.Model).
		SetRegisteredAt(car.RegisteredAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &Car{
		ID:           res.ID,
		Model:        res.Model,
		RegisteredAt: res.RegisteredAt,
	}, nil
}

func (r *Repository) GetCars(ctx context.Context) ([]*Car, error) {
	res, err := r.c.Car.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	cars := make([]*Car, 0, len(res))
	for _, car := range res {
		cars = append(cars, &Car{
			ID:           car.ID,
			Model:        car.Model,
			RegisteredAt: car.RegisteredAt,
		})
	}
	return cars, nil
}

func (r *Repository) GetCar(ctx context.Context, id int) (*Car, error) {
	res, err := r.c.Car.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Car{
		ID:           res.ID,
		Model:        res.Model,
		RegisteredAt: res.RegisteredAt,
	}, nil
}

func (r *Repository) UpdateCar(ctx context.Context, id int, car *CreateCar) (*Car, error) {
	res, err := r.c.Car.UpdateOneID(id).
		SetModel(car.Model).
		SetRegisteredAt(car.RegisteredAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &Car{
		ID:           res.ID,
		Model:        res.Model,
		RegisteredAt: res.RegisteredAt,
	}, nil
}

func (r *Repository) DeleteCar(ctx context.Context, id int) error {
	return r.c.Car.DeleteOneID(id).Exec(ctx)
}
