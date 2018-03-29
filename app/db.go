package app

import (
	"gopkg.in/mgo.v2"
)

type DB struct {
	addr     string
	database string
}

func NewDB(host, port, database string) *DB {
	addr := host + ":" + port
	return &DB{addr, database}
}

func (db *DB) FindAllTodos() (Todos, error) {
	session, err := mgo.Dial(db.addr)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var results Todos
	c := session.DB(db.database).C("todos")
	c.Find(nil).All(&results)
	return results, err
}
