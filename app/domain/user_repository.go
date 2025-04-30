package domain

import (
	"context"
)

type UserRepository interface {
	Create(user User) (User, error)
	GetByID(id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByCPF(ctx context.Context, cpf string) (User, error)
	ExistsByEmail(email string) bool
	ExistsByCPF(cpf string) bool
	CreateWithRole(user User) (User, error)
	IsEmpresa(userID uint) (bool, error)
	IsConsumidor(userID uint) (bool, error)
	FindUserByCPF(cpf string) (uint, error)
	FindEmpresaIDByUserID(userID uint) (uint, error)
	GetScoreByUserID(userID uint) (float64, error)
	UpdateScoreByUserID(userID uint, novoScore float64) error
}
