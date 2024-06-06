package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but did not want one")
	}
}

func assertError(t *testing.T, got error, expected error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but it did not get one")
	}

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
