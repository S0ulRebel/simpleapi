package mongo

import (
	"context"
	"os"
	"simple-api/errors"
	"simple-api/initializer"
	"simple-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
}

func (r *MongoUserRepository) CreateUser(user model.User) (model.User, *errors.AppError) {
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return user, errors.NewErrorService().InternalServerError(err)
	}
	return user, nil
}

func (r *MongoUserRepository) GetUsers() ([]model.User, *errors.AppError) {
	var users []model.User
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return users, errors.NewErrorService().InternalServerError(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (r *MongoUserRepository) GetUserByID(id string) (model.User, *errors.AppError) {
	var user model.User
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")
	err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		return user, errors.NewErrorService().NotFound("User")
	}
	return user, nil
}

func (r *MongoUserRepository) GetUserByEmail(email string) (model.User, *errors.AppError) {
	var user model.User
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")
	err := collection.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return user, errors.NewErrorService().NotFound("User")
	}
	return user, nil
}

func (r *MongoUserRepository) UpdateUser(id string, updatedUser model.User) (model.User, *errors.AppError) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var user model.User
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")

	err := collection.FindOneAndUpdate(context.Background(), bson.D{{"_id", id}}, bson.D{{"$set", updatedUser}}, opts).Decode(&user)
	if err != nil {
		return user, errors.NewErrorService().InternalServerError(err)
	}
	return user, nil
}

func (r *MongoUserRepository) DeleteUser(id string) *errors.AppError {
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return errors.NewErrorService().InternalServerError(err)
	}
	return nil
}
