package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"Repositories/E-COM/models"

)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already existed"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"error": ""})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone number is already existed"})
			return
		}
	}
}

func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func searchProduct() gin.HandlerFunc {

}

func serachProductByQuery() gin.HandlerFunc {

}
