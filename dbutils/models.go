package dbutils

type LoginData struct {
	Token string `json:"token"`
	Error *int   `json:"error,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Sighting struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	SeenAt        string `json:"SeenAt"`
	SeenAtLat     string `json:"SeenAtLat"`
	SeenAtLon     string `json:"SeenAtLon"`
	PhotoLocation string `json:"PhotoLocation"`
}

type UserDataWithPassword struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type TigerData struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	UserName    string      `json:"userName"`
	Name        string      `json:"name"`
	DateOfBirth string      `json:"dateOfBirth"`
	Sightings   []*Sighting `json:"Sightings" gorm:"foreignKey:ID;references:ID"`
}

type TigerDataLastSeen struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	UserName     string    `json:"userName"`
	Name         string    `json:"name"`
	DateOfBirth  string    `json:"dateOfBirth"`
	LastSighting *Sighting `json:"LastSighting" gorm:"foreignKey:ID;references:ID"`
}
