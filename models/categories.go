package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id        primitive.ObjectID `form:"_id,omitempty" bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `form:"name,omitempty" bson:"name,omitempty" json:"name,omitempty"`
	ImageLink string             `form:"image_link,omitempty" bson:"image_link,omitempty" json:"image_link,omitempty"`
}
