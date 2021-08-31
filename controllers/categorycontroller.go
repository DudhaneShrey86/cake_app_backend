package controllers

import (
	"context"
	"net/http"

	"github.com/DudhaneShrey86/cake_app_back/connection"
	"github.com/DudhaneShrey86/cake_app_back/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllCategories(c *gin.Context) {
	categoryColl := connection.CategoryColl
	cursor, err := categoryColl.Find(context.Background(), bson.M{})
	if err != nil {
		SendError(c, err)
		return
	}
	defer cursor.Close(connection.MainCtx)
	var results []bson.M
	err = cursor.All(connection.MainCtx, &results)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"output": "success",
		"data":   results,
	})
}

func GetCategoryById(c *gin.Context) {
	categoryColl := connection.CategoryColl
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		SendError(c, err)
		return
	}
	var result models.Category
	if err := categoryColl.FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&result); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"output": "success",
		"data":   result,
	})
}
