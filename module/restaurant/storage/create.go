package restaurantStorage

import (
	"context"
	"github.com/kienmatu/restaurants-go/common"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := s.db.Table(restaurantModel.RestaurantCreate{}.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
