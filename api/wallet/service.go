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

	return w, nil
}
