package model

import "time"

//Province data
type Province struct {
	Id         int8      `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

//Provinces data
type Provinces struct {
	Data      []Province `json:"data"`
	TotalData int        `json:"totalData"`
}
