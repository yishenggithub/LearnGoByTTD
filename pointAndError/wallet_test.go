package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
} // func是可以写在里面的

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Error("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
