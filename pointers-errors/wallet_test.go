package main

import (
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(0)

		utils.AssertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw insuffient funds", func(t *testing.T) {
		startingBalance := Bitcoin(0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(1))

		utils.AssertError(t, err, ErrInsuficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
