package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet.Balance(), 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		wallet.Withdraw(15)

		assertBalance(t, wallet.Balance(), 5)
	})

	t.Run("will throw error when insufficient funds", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(25)
		if err == nil {
			t.Error("expected error but didn't get one")
		}

		if err != InsufficientFundsError {
			t.Error("expected InsufficientFundsError but didn't get one")
		}

	})
}

func assertBalance(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
