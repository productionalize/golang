package app

import "github.com/globalsign/mgo/bson"

// Todo represents a todo item.
type Todo struct {
	ID        bson.ObjectId `bson:"_id"`
	Content   string        `json:"content"`
	Completed bool          `json:"completed"`
}

// Todos represents a list of todo items.
type Todos []Todo
