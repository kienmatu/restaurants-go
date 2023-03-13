package restaurantStorage

import (
	"context"
	"errors"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := s.db.Create(data); err != nil {
		return errors.New("can not create restaurant")
	}
	return nil
}

func (s *sqlStore) Update(ctx context.Context, data *restaurantModel.RestaurantUpdate) error {
	if err := s.db.Updates(data); err != nil {
		return errors.New("can not create restaurant")
	}
	return nil
}
