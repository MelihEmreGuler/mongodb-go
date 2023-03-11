package repository

import (
	"context"
	"errors"
	"time"

	"github.com/MelihEmreGuler/mongodb-go/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection //
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // cancel when we are finished (context.CancelFunc returned by context.WithTimeout)

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil { // if there is an error or no id is returned
		errors.New("Error while inserting todo")
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{
		TodoCollection: dbClient, // dbClient is the mongo client
	}
}
