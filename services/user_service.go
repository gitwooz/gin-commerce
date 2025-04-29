package services

import (
	"context"
	"errors"
	"time"

	"github.com/gitwooz/go-gin-app/models"
	"github.com/gitwooz/go-gin-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(databaseName string) *UserService {
	return &UserService{
		collection: utils.GetCollection(databaseName, "users"),
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, user)
	return err
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.GetUserByID(id) // Alias for GetUserByID
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *UserService) UpdateUser(id string, updatedUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedUser})
	return err
}

func (s *UserService) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
