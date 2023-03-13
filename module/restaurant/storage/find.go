package restaurantStorage

import (
	"context"
	"errors"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) FindByCondition(ctx context.Context, cond map[string]interface{}) ([]*restaurantModel.Restaurant, error) {
	data := []*restaurantModel.Restaurant{}
	if err := s.db.Find(data); err != nil {
		return nil, errors.New("can not find restaurant")
	}
	return data, nil
}

func (s *sqlStore) FindByID(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {
	var restaurant restaurantModel.Restaurant
	if err := s.db.Find(&restaurant, id); err != nil {
		return nil, errors.New("can not find restaurant")
	}
	return &restaurant, nil
}
