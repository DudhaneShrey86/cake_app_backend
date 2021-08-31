package routes

import (
	con "github.com/DudhaneShrey86/cake_app_back/controllers"
	"github.com/DudhaneShrey86/cake_app_back/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.GET("/getallcakes", con.GetAllCakes)
	r.GET("/getcake/:id", con.GetCakeById)
	r.GET("/getcakebysearch", con.GetCakeByName)
	r.GET("/getallcategories", con.GetAllCategories)
	r.GET("/getcategory/:id", con.GetCategoryById)
	r.Run(":8181")
}
