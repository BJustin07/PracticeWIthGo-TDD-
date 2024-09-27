package pointersAndErrors

import (
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("expected an error but got none")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	checkWalletBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Deposit to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(500))
		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		checkWalletBalance(t, wallet, Bitcoin(500))
	})
	t.Run("Withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(500)}
		wallet.Withdraw(Bitcoin(250))
		checkWalletBalance(t, wallet, Bitcoin(250))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		checkWalletBalance(t, wallet, Bitcoin(20))
	})
}
