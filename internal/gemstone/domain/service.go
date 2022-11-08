package domain

// IService , interface defining all CRUD operations
type IService interface {
	Find() error
	Store() error
	Update() error
	FindAll() error
	Delete() error
}

// IRepository , interface acting like a port for the database implementation
type IRepository interface {
	Find() error
	Store() error
	Update() error
	FindAll() error
	Delete() error
}
