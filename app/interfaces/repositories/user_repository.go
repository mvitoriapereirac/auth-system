package repositories

import (
	"auth-system/app/core/constants"
	"auth-system/app/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *userRepository) GetByCPF(ctx context.Context, cpf string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("cpf = ?", cpf).First(&user).Error; err != nil {
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
			consumidor := domain.Consumidor{UserID: user.ID, Score: constants.UserMaxScore} //Idealmente, numero estaria em constante
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

func (r *userRepository) IsEmpresa(userID uint) (bool, error) {
	var exists bool
	query := `
    SELECT EXISTS(
        SELECT 1 
        FROM empresa 
        WHERE user_id = ? 
        LIMIT 1
    )`

	if err := r.db.Raw(query, userID).Scan(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) IsConsumidor(userID uint) (bool, error) {
	var exists bool
	query := `
    SELECT EXISTS(
        SELECT 1 
        FROM consumidor 
        WHERE user_id = ? 
        LIMIT 1
    )`

	if err := r.db.Raw(query, userID).Scan(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) FindUserByCPF(cpf string) (uint, error) {
	var user domain.User
	err := r.db.Where("cpf = ?", cpf).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.New("usuário não encontrado")
		}
		return 0, err
	}

	return user.ID, nil
}

func (r *userRepository) FindEmpresaIDByUserID(userID uint) (uint, error) {
	var empresaID uint
	query := `
        SELECT id 
        FROM empresa 
        WHERE user_id = ? 
        LIMIT 1
    `

	if err := r.db.Raw(query, userID).Scan(&empresaID).Error; err != nil {
		if err == gorm.ErrRecordNotFound || empresaID == 0 {
			return 0, errors.New("empresa não encontrada para este usuário")
		}
		return 0, err
	}

	return empresaID, nil
}

func (r *userRepository) GetScoreByUserID(userID uint) (float64, error) {
	var score float64
	query := `SELECT c.score FROM consumidor c WHERE c.user_id = @userID`
	err := r.db.Raw(query, sql.Named("userID", userID)).Scan(&score).Error
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (r *userRepository) UpdateScoreByUserID(userID uint, novoScore float64) error {
	query := `UPDATE consumidor SET score = @novoScore WHERE user_id = @userID`
	return r.db.Exec(query,
		sql.Named("novoScore", novoScore),
		sql.Named("userID", userID),
	).Error
}
