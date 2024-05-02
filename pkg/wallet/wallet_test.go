package wallet

import (
	"testing"

	"github.com/ainizoda/go-lessons-wallet/pkg/types"
)

func TestPay_Success(t *testing.T) {
	service := Service{}
	account, err := service.RegisterAccount("+992930000001")

	if err != nil {
		t.Error(err)
		return
	}
	account.Balance = 1000
	payment, err := service.Pay(account.ID, 250, "food")

	if err != nil {
		t.Error(err)
		return
	}
	if payment.Status != types.StatusInProgress {
		t.Errorf(`expected payment status to be "%s", got "%s"`, types.StatusInProgress, payment.Status)
	}
	if account.Balance != 750 {
		t.Errorf("expected account balance to be %d, got %d", 750, account.Balance)
	}
}

func TestReject_Success(t *testing.T) {
	service := Service{}
	account, err := service.RegisterAccount("+992930000001")

	if err != nil {
		t.Error(err)
		return
	}
	account.Balance = 1000
	payment, err := service.Pay(account.ID, 250, "food")
	if err != nil {
		t.Error(err)
		return
	}
	err = service.Reject(payment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if payment.Status != types.StatusFail {
		t.Errorf(`expected payment status to be "%s", got "%s"`, types.StatusFail, payment.Status)
	}
	if account.Balance != 1000 {
		t.Errorf("expected account balance to be %d, got %d", 1000, account.Balance)
	}
}
