package restaurantBiz

import (
	"context"

	"github.com/kienmatu/restaurants-go/common"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListByCondition(ctx context.Context, filter *restaurantModel.Filter, paging *common.Paging, moreKeys ...string) ([]*restaurantModel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (r *listRestaurantBiz) ListByCondition(ctx context.Context, filter *restaurantModel.Filter, paging *common.Paging) ([]*restaurantModel.Restaurant, error) {
	data, err := r.store.ListByCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
