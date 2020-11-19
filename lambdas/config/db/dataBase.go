package db

// DataBase representation basic actions on data base
type DataBase interface {
	OpenConnection() error
	GetConnection() interface{}
}
