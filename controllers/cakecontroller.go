package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/DudhaneShrey86/cake_app_back/connection"
	"github.com/DudhaneShrey86/cake_app_back/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SendError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"output": "error",
		"data":   err.Error(),
	})
}

func GetAllCakes(c *gin.Context) {
	cakeColl := connection.CakeColl
	id, err := primitive.ObjectIDFromHex(c.Query("category_id"))
	if err != nil {
		SendError(c, err)
		return
	}
	opts := options.Find()
	if c.Query("limit") != "" {
		l, _ := strconv.Atoi(c.Query("limit"))
		opts.SetLimit(int64(l))
	}
	cursor, err := cakeColl.Find(context.Background(), bson.M{
		"category_id": id,
	}, opts)
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

func GetCakeById(c *gin.Context) {
	cakeColl := connection.CakeColl
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		SendError(c, err)
		return
	}
	var result models.Cake
	if err := cakeColl.FindOne(context.Background(), bson.M{
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

func GetCakeByName(c *gin.Context) {
	cakeColl := connection.CakeColl
	searchString := strings.ToLower(c.Query("search_string"))
	cursor, err := cakeColl.Find(context.Background(), bson.M{
		"name": bson.M{"$regex": primitive.Regex{Pattern: searchString, Options: "i"}},
	})
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
