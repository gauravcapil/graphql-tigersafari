// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginData struct {
	Token string `json:"token"`
	Error *int   `json:"error,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Sighting struct {
	SeenAt        string `json:"SeenAt"`
	SeenAtLat     string `json:"SeenAtLat"`
	SeenAtLon     string `json:"SeenAtLon"`
	PhotoLocation string `json:"PhotoLocation"`
}

type TigerData struct {
	UserName    string      `json:"userName"`
	Name        string      `json:"name"`
	DateOfBirth string      `json:"dateOfBirth"`
	Sightings   []*Sighting `json:"Sightings"`
}

type TigerDataLastSeen struct {
	UserName     string    `json:"userName"`
	Name         string    `json:"name"`
	DateOfBirth  string    `json:"dateOfBirth"`
	LastSighting *Sighting `json:"LastSighting"`
}

type UserData struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
