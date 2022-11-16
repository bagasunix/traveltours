package base

type Repository interface {
	GetModelName() string
	GetConnection() (T any)
}
