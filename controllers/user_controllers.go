package controllers

import (
	"context"
	"net/http"

	"github.com/faztweb/go-fiber-mongodb/config"
	"github.com/faztweb/go-fiber-mongodb/models"
	"github.com/faztweb/go-fiber-mongodb/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": validationErr.Error()},
		})
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	result, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "Created Successfully",
		Data: &fiber.Map{
			"data": result,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	var user models.User

	objectId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Interal server error",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: &fiber.Map{
			"data": user,
		},
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	results, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{})
	}

	for results.Next(context.TODO()) {
		var user models.User
		results.Decode(&user)
		users = append(users, user)
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"users": users},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objId}})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}
