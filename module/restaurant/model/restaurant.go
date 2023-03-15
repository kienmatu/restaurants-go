package restaurantModel

import (
	"errors"
	"github.com/kienmatu/restaurants-go/common"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	OwnerID         int    `json:"ownerId" gorm:"column:owner_id"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
	Status          string `json:"status" gorm:"column:status"`
}

func (r Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Addr    string `json:"addr" gorm:"column:addr"`
	OwnerID int    `json:"ownerId" gorm:"column:owner_id"`
}
type RestaurantUpdate struct {
	ID     int    `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	Status string `json:"status" gorm:"column:status"`
}

func (r RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
func (r RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
func (r *RestaurantCreate) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("Empty name")
	}
	return nil
}

/**
CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) NULL,
  `name` varchar(50) NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int(11) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json NULL,
  `logo` json NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
