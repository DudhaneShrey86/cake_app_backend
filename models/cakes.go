package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cake struct {
	Id          primitive.ObjectID `form:"_id,omitempty" bson:"_id,omitempty" json:"_id,omitempty"`
	CategoryId  primitive.ObjectID `form:"category_id,omitempty" bson:"category_id,omitempty" json:"category_id,omitempty"`
	Name        string             `form:"name,omitempty" bson:"name,omitempty" json:"name,omitempty"`
	Time        int                `form:"time,omitempty" bson:"time,omitempty" json:"time,omitempty"`
	Ingredients []string           `form:"ingredients,omitempty" bson:"ingredients,omitempty" json:"ingredients,omitempty"`
	Recipe      string             `form:"recipe,omitempty" bson:"recipe,omitempty" json:"recipe,omitempty"`
	ImageLink   string             `form:"image_link,omitempty" bson:"image_link,omitempty" json:"image_link,omitempty"`
}
