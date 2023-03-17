package restaurantStorage

import (
	"context"
	"github.com/kienmatu/restaurants-go/common"

	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context, data *restaurantModel.RestaurantUpdate) error {
	if err := s.db.Table(restaurantModel.RestaurantUpdate{}.TableName()).Save(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
