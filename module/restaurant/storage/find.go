package restaurantStorage

import (
	"context"
	"github.com/kienmatu/restaurants-go/common"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) FindOneByCondition(ctx context.Context, cond map[string]interface{}) (*restaurantModel.Restaurant, error) {
	var data = restaurantModel.Restaurant{}
	if err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		Where(cond).
		First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}

// using
func (s *sqlStore) FindByID(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {
	var restaurant *restaurantModel.Restaurant
	if err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		First(&restaurant, "id = ?", id).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return restaurant, nil
}
