package restaurantBiz

import (
	"context"
	restaurantModel "github.com/kienmatu/restaurants-go/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	Update(ctx context.Context, data *restaurantModel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (r *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, data *restaurantModel.RestaurantUpdate) error {
	if err := r.store.Update(ctx, data); err != nil {
		return err
	}
	return nil
}
