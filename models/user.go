package models

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Enable     string    `json:"enable"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DisabledAt time.Time `json:"disabled_at"`
}

var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func GetAllUsers(c *gin.Context) {
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	email := user.Email
	password := user.Password
	id := guuid.New().String()

	newUser := User{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")

	user := User{}
	err := collection.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})
	return
}

func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	var user User
	c.BindJSON(&user)

	newData := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": userId}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Edited Successfully",
	})
	return
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": userId})
	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}
