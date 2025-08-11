package wallet

import "github.com/labstack/echo/v4"

type walletService struct {
	repository walletRepository
}

func NewWalletService(r walletRepository) walletService {
	return walletService{
		repository: r,
	}
}

func (s walletService) create(c echo.Context, req createWalletReq) (wallet, error) {
	w, err := s.repository.create(c, req)
	if err != nil {
		return wallet{}, err
	}

	walletItem := wallet{
		ID:          w.ID,
		UserID:      w.UserID,
		Name:        w.Name,
		Description: w.Description,
		Currency:    w.Currency,
		CreatedAt:   w.CreatedAt.Time.String(),
		UpdatedAt:   w.UpdatedAt.Time.String(),
	}

	return walletItem, nil
}
