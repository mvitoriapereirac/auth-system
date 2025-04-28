package domain

type UserRepository interface {
    Create(user User) (User, error)
    GetByID(id uint) (User, error)
}