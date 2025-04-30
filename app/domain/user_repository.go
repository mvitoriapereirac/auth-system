package domain

import (
    "context"
)

type UserRepository interface {
    Create(user User) (User, error)
    GetByID(id uint) (User, error)
    GetByEmail(ctx context.Context, email string) (User, error)
    ExistsByEmail(email string) (bool)
    ExistsByCPF(cpf string) (bool)
    CreateWithRole(user User) (User, error)
}