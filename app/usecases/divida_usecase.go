package usecases

import (
	"auth-system/app/domain"
	"auth-system/app/dto"
	"auth-system/pkg"
	"errors"
	"math"
)

type DividaUsecase interface {
	CriarDivida(divida domain.Divida, cpfConsumidor string, empresaID uint) (dto.DividaResponse, error)
	ListarDividas(id uint64) ([]domain.Divida, error)
	AtualizarEObterScore(id uint64) (float64, error)
}

type dividaUsecase struct {
	dividaRepo domain.DividaRepository
	userRepo   domain.UserRepository
}

func NewDividaUsecase(dividaRepo domain.DividaRepository, userRepo domain.UserRepository) DividaUsecase {
	return &dividaUsecase{
		dividaRepo: dividaRepo,
		userRepo:   userRepo,
	}
}

// Durante o desenvolvimento, surgiu a percepção de que seria interessante permitir
// que empresas registrem dívidas mesmo quando o consumidor ainda não está cadastrado
// na plataforma.
// Embora essa lógica não tenha sido implementada, houve a intenção de, em um segundo
// momento, viabilizar esse fluxo utilizando usuários "placeholder", permitindo maior
// flexibilidade e refletindo práticas comuns em sistemas similares. Este registro
// serve para sinalizar essa preocupação, ainda que o tempo limitado do teste técnico
// tenha impedido sua concretização.
func (u *dividaUsecase) CriarDivida(divida domain.Divida, cpfConsumidor string, userID uint) (dto.DividaResponse, error) {

	if err := validators.ValidateDividaInput(divida, cpfConsumidor); err != nil {
		return dto.DividaResponse{}, err
	}
	divida.NaturezaCobrancaID = 1 //Idealmente, seria possível editar esse campo em outros endpoints. A ideia é que haja mais informações sobre a dívida, aproximando-a de um objeto de um sistema real

	empresaID, err := u.userRepo.FindEmpresaIDByUserID(userID)
	if err != nil || empresaID == 0 {
		return dto.DividaResponse{}, errors.New("somente usuários de empresas podem adicionar dívidas")
	}

	consumidorUserID, err := u.userRepo.FindUserByCPF(cpfConsumidor)
	if err != nil {
		return dto.DividaResponse{}, errors.New("consumidor ainda não registrado")
	}

	isConsumidor, err := u.userRepo.IsConsumidor(consumidorUserID)
	if err != nil {
		return dto.DividaResponse{}, errors.New("erro ao verificar se o usuário é consumidor")
	}
	if !isConsumidor {
		return dto.DividaResponse{}, errors.New("somente consumidores podem ter dívidas")
	}

	createdDivida, err := u.dividaRepo.Create(divida, cpfConsumidor, empresaID)
	if err != nil {
		return dto.DividaResponse{}, err
	}
	dividaResponse := dto.DividaResponse{
		ID:         createdDivida.ID,
		Valor:      createdDivida.Valor,
		Vencimento: createdDivida.Vencimento,
	}
	return dividaResponse, nil
}

func (u *dividaUsecase) ListarDividas(userID uint64) ([]domain.Divida, error) {
	isConsumidor, err := u.userRepo.IsConsumidor(uint(userID))
	if err != nil {
		return []domain.Divida{}, errors.New("erro ao verificar se o usuário é consumidor")
	}
	if !isConsumidor {
		return []domain.Divida{}, errors.New("somente o autor da dívida pode verificá-la")
	}
	dividas, err := u.dividaRepo.ListByID(uint(userID))
	if err != nil {
		return []domain.Divida{}, err
	}

	return dividas, nil
}

func (u *dividaUsecase) AtualizarEObterScore(userID uint64) (float64, error) {
	isConsumidor, err := u.userRepo.IsConsumidor(uint(userID))
	if err != nil {
		return 0, errors.New("erro ao verificar se o usuário é consumidor")
	}
	if !isConsumidor {
		return 0, errors.New("somente consumidores podem ter score")
	}

	scoreAtual, err := u.userRepo.GetScoreByUserID(uint(userID))
	if err != nil {
		return 0, errors.New("erro ao buscar score do consumidor")
	}

	dividas, err := u.dividaRepo.ListByID(uint(userID))
	if err != nil {
		return 0, errors.New("erro ao listar dívidas do consumidor")
	}

	if len(dividas) == 0 {
		return scoreAtual, nil // Não atualiza se não houver dívidas
	}

	total := 0.0
	for _, d := range dividas {
		total += d.Valor
	}
	media := total / float64(len(dividas))

	scoreCalculado := 10000 / math.Sqrt(media+100)

	if math.Abs(scoreCalculado-scoreAtual) > 0.01 {
		err = u.userRepo.UpdateScoreByUserID(uint(userID), scoreCalculado)
		if err != nil {
			return 0, errors.New("erro ao atualizar score do consumidor")
		}
	}

	return scoreCalculado, nil
}
