package models

type User struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	Following    []*User `json:"following,omitempty" bson:"-"`
	Followers    []*User `json:"followers,omitempty" bson:"-"`
	FollowingIds []int   `json:"followingIds" bson:"following"`
	FollowersIds []int   `json:"followersIds" bson:"followers"`
}
