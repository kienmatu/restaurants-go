package restaurantStorage

import (
	"context"
	"github.com/kienmatu/restaurants-go/common"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

func (s *sqlStore) ListByCondition(
	ctx context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]*restaurantModel.Restaurant, error) {
	var restaurants []*restaurantModel.Restaurant
	db := s.db.Table(restaurantModel.Restaurant{}.TableName()).Where("status != 0")
	if f := filter; f != nil {
		if f.OwnerID > 0 {
			db.Where("owner_id = ?", filter.OwnerID)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}
