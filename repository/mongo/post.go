package mongo

import (
	"context"
	"fmt"
	"os"
	"simple-api/errors"
	"simple-api/initializer"
	"simple-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPostRepository struct {
}

func (r *MongoPostRepository) CreatePost(post model.Post) (model.Post, *errors.AppError) {
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return post, errors.NewErrorService().InternalServerError(err)
	}
	return post, nil
}

func (r *MongoPostRepository) GetPosts() ([]model.Post, *errors.AppError) {
	var posts []model.Post
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return posts, errors.NewErrorService().InternalServerError(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post model.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *MongoPostRepository) GetPostByID(id string) (model.Post, *errors.AppError) {
	var post model.Post
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&post)
	if err != nil {
		return post, errors.NewErrorService().NotFound("Post")
	}
	return post, nil
}

func (r *MongoPostRepository) GetPostsByUserID(userID string) ([]model.Post, *errors.AppError) {
	fmt.Println(userID)
	var posts []model.Post
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	cursor, err := collection.Find(context.Background(), bson.D{{"userid", userID}})
	if err != nil {
		return posts, errors.NewErrorService().InternalServerError(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post model.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	fmt.Println(posts)
	return posts, nil
}

func (r *MongoPostRepository) UpdatePost(id string, updatedPost model.Post) (model.Post, *errors.AppError) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var post model.Post
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	err := collection.FindOneAndUpdate(context.Background(), bson.D{{"_id", id}}, bson.D{{"$set", updatedPost}}, opts).Decode(&post)
	if err != nil {
		return post, errors.NewErrorService().InternalServerError(err)
	}
	return post, nil
}

func (r *MongoPostRepository) DeletePost(id string) *errors.AppError {
	collection := initializer.MGDB.Database(os.Getenv("MONGO_DB")).Collection("posts")
	_, err := collection.DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return errors.NewErrorService().InternalServerError(err)
	}
	return nil
}
