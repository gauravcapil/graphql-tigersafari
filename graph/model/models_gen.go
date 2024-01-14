// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginData struct {
	Token      string `json:"token"`
	Userid     int    `json:"userid"`
	Expiration string `json:"expiration"`
	Error      *int   `json:"error,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Sighting struct {
	ID            int     `json:"id" gorm:"primary_key"`
	TigerID       int     `json:"TigerId"`
	SeenAt        string  `json:"SeenAt"`
	SeenAtLat     float64 `json:"SeenAtLat"`
	SeenAtLon     float64 `json:"SeenAtLon"`
	PhotoLocation string  `json:"PhotoLocation"`
}

type TigerData struct {
	ID          int         `json:"id" gorm:"primary_key"`
	UserName    string      `json:"userName"`
	Name        string      `json:"name"  gorm:"uniqueIndex"`
	DateOfBirth string      `json:"dateOfBirth"`
	Sightings   []*Sighting `json:"Sightings"  gorm:"foreignKey:TigerID;references:ID"`
}

type TigerDataResponse struct {
	TigerID       int     `json:"TigerID"`
	SeenAt        string  `json:"SeenAt"`
	SeenAtLat     float64 `json:"SeenAtLat"`
	SeenAtLon     float64 `json:"SeenAtLon"`
	PhotoLocation string  `json:"PhotoLocation"`
	UserName      string  `json:"UserName"`
	Name          string  `json:"Name"`
	DateOfBirth   string  `json:"DateOfBirth"`
}


type UserData struct {
	ID       int    `json:"id"  gorm:"primary_key"`
	UserName string `json:"userName"  gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
}

type UserDataWithPassword struct {
	ID       int    `json:"id"  gorm:"primary_key"`
	UserName string `json:"userName"  gorm:"uniqueIndex"`
	Password string `json:"password"`
}
