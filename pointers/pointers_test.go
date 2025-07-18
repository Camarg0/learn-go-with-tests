package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(1))

		assertBalance(t, wallet, Bitcoin(1))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		AssertError(t, err, ErrorInsufficientFunds)
		assertBalance(t, wallet, wallet.balance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("want no error but did get one")
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("want an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
