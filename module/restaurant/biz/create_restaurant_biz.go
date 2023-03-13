package restaurantBiz

import (
	"context"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error
	//Update(ctx context.Context, data *restaurantModel.RestaurantUpdate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (r *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := r.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
