package domain

// Service , interface defining all CRUD operations
type Service interface {
	Find() error
	Store() error
	Update() error
	FindAll() error
	Delete() error
}

// Repository , interface acting like a port for the database implementation
type Repository interface {
	Find() error
	Store() error
	Update() error
	FindAll() error
	Delete() error
}
