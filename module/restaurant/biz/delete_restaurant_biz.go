package restaurantBiz

import (
	"context"
)

type DeleteRestaurantStore interface {
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (r *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	if err := r.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
