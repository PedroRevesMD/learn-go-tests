package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit Test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})
	t.Run("Withdrawal Test", func(t *testing.T) {
		wallet := Wallet{balance: Coin(20)}
		err := wallet.Withdrawal(Coin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Coin(10))
	})
	t.Run("Withdrawal Insufficient Funds Test", func(t *testing.T) {
		startingBalance := Coin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawal(Coin(100))
		assertError(t, err, "Cannot Withdrawal Insufficient Funds")
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Coin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("We want an error but didnt get one so far...")
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Errorf("We want an error but didnt get one so far...")
	}
	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}

}
