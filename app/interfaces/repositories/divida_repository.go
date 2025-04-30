package repositories

import (
	"auth-system/app/domain"
	"database/sql"
	"gorm.io/gorm"
	"strconv"
)

type dividaRepository struct {
	db *gorm.DB
}

func NewDividaRepository(db *gorm.DB) domain.DividaRepository {
	return &dividaRepository{db: db}
}

func (r *dividaRepository) Create(divida domain.Divida, cpf string, empresaID uint) (domain.Divida, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Cria a dívida
		if err := tx.Create(&divida).Error; err != nil {
			return err
		}

		var consumidorID string
		query := `
        SELECT c.id 
        FROM consumidor c 
        INNER JOIN "user" u ON c.user_id = u.id 
        WHERE u.cpf = @cpf
        `
		err := tx.Raw(query, sql.Named("cpf", cpf)).Scan(&consumidorID).Error
		if err != nil {
			return err
		}

		consumidorIDUint, err := strconv.Atoi(consumidorID)
		if err != nil {
			return err
		}

		associacao := domain.DividaConsumidorEmpresa{
			DividaID:     divida.ID,
			ConsumidorID: uint(consumidorIDUint),
			EmpresaID:    empresaID,
		}

		if err := tx.Create(&associacao).Error; err != nil {
			return err
		}

		return nil
	})
	return divida, err
}

func (r *dividaRepository) ListByID(id uint) ([]domain.Divida, error) {
	var dividas []domain.Divida

	query := `
    SELECT d.*
    FROM divida d
    INNER JOIN divida_consumidor_empresa dce ON d.id = dce.divida_id
    INNER JOIN consumidor c ON dce.consumidor_id = c.id
    INNER JOIN "user" u ON c.user_id = u.id
    WHERE u.id = @id
    `

	err := r.db.Raw(query, sql.Named("id", id)).Scan(&dividas).Error
	if err != nil {
		return nil, err
	}

	return dividas, nil
}
