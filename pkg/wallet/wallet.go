package wallet

import "github.com/ainizoda/go-lessons-wallet/pkg/types"

type Service struct {
	netAccountID int64
	accounts     []types.Account
	payments     []types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return
		}
	}
	s.netAccountID++
	s.accounts = append(s.accounts, types.Account{
		ID:      s.netAccountID,
		Phone:   phone,
		Balance: 0,
	})
}
