package restaurantStorage

import (
	"context"
	"errors"
	"fmt"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) FindOneByCondition(ctx context.Context, cond map[string]interface{}) (*restaurantModel.Restaurant, error) {
	var data = restaurantModel.Restaurant{}
	if err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		Where(cond).
		First(&data).Error; err != nil {
		return nil, errors.New("can not find restaurant")
	}
	return &data, nil
}

func (s *sqlStore) FindByID(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {
	var restaurant *restaurantModel.Restaurant
	if err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		First(&restaurant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	fmt.Println("nil?", restaurant == nil)
	return restaurant, nil
}
