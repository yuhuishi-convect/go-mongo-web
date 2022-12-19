package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TodoStatus is the status of a todo item
type TodoStatus string

// TodoStatus constants
const (
	TodoStatusActive   TodoStatus = "active"
	TodoStatusComplete TodoStatus = "complete"
)

type TodoItem struct {
	ID          primitive.ObjectID `json:"id", bson:"_id", omitempty`
	Name        string             `json:"name", bson:"name"`
	Description string             `json:"description", bson:"description"`
	Status      TodoStatus         `json:"status", bson:"status", default:"active"`
}

// TodoItemRepository is the interface for the todo item repository
type TodoItemRepository interface {
	FindAll() ([]TodoItem, error)
	FindByID(id string) (TodoItem, error)
	Save(item TodoItem) error
}

// TodoItemMongoRepository is the mongo implementation of the todo item repository
type TodoItemMongoRepository struct {
	collection *mongo.Collection
}

// NewTodoItemMongoRepository creates a new TodoItemMongoRepository
func NewTodoItemMongoRepository(collection *mongo.Collection) *TodoItemMongoRepository {
	return &TodoItemMongoRepository{collection: collection}
}

// FindAll returns all todo items
func (r *TodoItemMongoRepository) FindAll() ([]TodoItem, error) {
	context := context.Background()
	cursor, err := r.collection.Find(context, bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var items []TodoItem
	err = cursor.All(context, &items)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

// FindByID returns a todo item by id
func (r *TodoItemMongoRepository) FindByID(id string) (TodoItem, error) {
	context := context.Background()
	var item TodoItem
	err := r.collection.FindOne(context, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return TodoItem{}, err
	}
	return item, nil
}

// Save saves a todo item
func (r *TodoItemMongoRepository) Save(item TodoItem) error {
	context := context.Background()
	_, err := r.collection.InsertOne(context, item)
	if err != nil {
		return err
	}
	return nil
}
