package models

// Profile struct
type Profile struct {
	Name     string
	Username string
	Owned    []BuildingInfo
}

// BuildingInfo basic struct for needed building information
type BuildingInfo struct {
	BuildingID int
	Name       string
	Summary    string
	Built      string
	Height     string
	Architect  string
	ImageURL   string
}
