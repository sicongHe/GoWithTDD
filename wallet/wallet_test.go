package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(0))
		assertNotError(t, err)
	})
	t.Run("withdraw bitcoin to < 0", func(t *testing.T) {
		startingWallet := Bitcoin(10)
		wallet := Wallet{startingWallet}
		err := wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, startingWallet)
		assertError(t, err, InsufficientFundsError.Error())
	})
}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %s, Want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("need an err but got nil")
	}
	if got.Error() != want {
		t.Errorf("Got %s, Want %s", got.Error(), want)
	}
}
func assertNotError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("Want a nil but not")
	}
}
