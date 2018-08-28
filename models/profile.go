package models

// Profile struct
type Profile struct {
	Name     string
	Username string
	Owned    []BuildingInfo
}

// BuildingInfo basic struct for needed building information
type BuildingInfo struct {
	BuildingID  int
	Title       string
	Address     string
	Description string
	Image       string
}
