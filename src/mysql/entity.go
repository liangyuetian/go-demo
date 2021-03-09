package mysql

import "database/sql"

type BaseEntity struct {
	Id          int `db:"id"`
	Status      int `db:"status"`
	CreatedTime int `db:"created_time"`
	UpdatedTime int `db:"updated_time"`
	Version     int `db:"version"`
}

type PreOrderEntity struct {
	BaseEntity
}

type ForwardSupplyEntity struct {
	BaseEntity
}

type UserFarmerLogEntity struct {
	BaseEntity
	ForwardSupplyEntity
	CustomerId     int            `db:"customer_id"`
	CustomerName   string         `db:"customer_name"`
	Role           int            `db:"role"`
	SellerCid      int            `db:"seller_cid"`
	SellerUserName sql.NullString `db:"seller_user_name"`
	ShareToken     sql.NullString `db:"share_token"`
	SupplyId       int            `db:"supply_id"`
	SupplyName     sql.NullString `db:"supply_name"`
	PageSource     sql.NullString `db:"page_source"`
}
