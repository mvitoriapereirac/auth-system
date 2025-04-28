package repositories

import (
    "auth-system/app/domain"
    "gorm.io/gorm"
)

// userRepository implementa a interface UserRepository
type userRepository struct {
    db *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepository{db: db}
}

// Create insere um novo usuário no banco de dados
func (repo *userRepository) Create(user domain.User) (domain.User, error) {
    if err := repo.db.Create(&user).Error; err != nil {
        return domain.User{}, err
    }
    return user, nil
}

// GetByID retorna um usuário pelo ID
func (repo *userRepository) GetByID(id uint) (domain.User, error) {
    var user domain.User
    if err := repo.db.First(&user, id).Error; err != nil {
        return domain.User{}, err
    }
    return user, nil
}
