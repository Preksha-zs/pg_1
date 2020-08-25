package models

// User schema of the user table
type FavLoc struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
}
