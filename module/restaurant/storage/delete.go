package restaurantStorage

import (
	"context"
	"github.com/kienmatu/restaurants-go/common"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
