package models
type Fav_loc struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
}
