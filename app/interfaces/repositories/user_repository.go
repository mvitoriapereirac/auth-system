package repositories

import (
    "context"
    "auth-system/app/domain"
    "gorm.io/gorm"
    "auth-system/app/core/constants"
    "fmt"
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
    user.Senha = ""
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

func (r *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
    var user domain.User
    if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
        return domain.User{}, err
    }
    return user, nil
}

func (r *userRepository) ExistsByCPF(cpf string) bool {
	var count int64
	r.db.Model(&domain.User{}).Where("cpf = ?", cpf).Count(&count)
	return count > 0
}

func (r *userRepository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) CreateWithRole(user domain.User) (domain.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Cria o usuário
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// Associa o tipo de usuário à respectiva entidade
		tipo := ""
		if user.Tipo != nil {
			tipo = *user.Tipo
		}
		switch tipo {
		case constants.UserTypeConsumidor:
			consumidor := domain.Consumidor{UserID: user.ID, Score: 1000.00}
			if err := tx.Create(&consumidor).Error; err != nil {
				return err
			}
		case constants.UserTypeEmpresa:
			empresa := domain.Empresa{UserID: user.ID}
			if err := tx.Create(&empresa).Error; err != nil {
				return err
			}
		default:
			return fmt.Errorf("tipo de usuário inválido: %s", user.Tipo)
		}

		return nil
	})

	return user, err
}
