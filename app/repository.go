package app

type Repository interface {
	FindAllTodos() (Todos, error)
}
