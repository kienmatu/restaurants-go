package restaurantBiz

import (
	"context"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

type GetRestaurantStore interface {
	FindByID(ctx context.Context, id int) (*restaurantModel.Restaurant, error)
	FindByCondition(ctx context.Context, cond map[string]interface{}) ([]*restaurantModel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (r *getRestaurantBiz) FindByID(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {
	data, err := r.store.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *getRestaurantBiz) FindByCondition(ctx context.Context, cond map[string]interface{}) ([]*restaurantModel.Restaurant, error) {
	data, err := r.store.FindByCondition(ctx, cond)
	if err != nil {
		return nil, err
	}
	return data, nil
}
