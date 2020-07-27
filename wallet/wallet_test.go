package wallet

import (
	"testing"
)

/**
golang的值传递：
在golang里面都是值传递，只不过有的是传的值的副本，有的是传的指针的副本。
普通类型：传值。
map、channel：本身初始化时返回的就是指针类型（例如：*hmap、*hchan）
slice：传的是指向底层数组的指针。
*/
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	//fmt.Printf("in test %p\n", &wallet)
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(100)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWallet2(t *testing.T) {
	// 提取子方法
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		// 增加t.Helper()表示当前为辅助函数  报错显示调用行
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	//assertError := func(t *testing.T, err error) {
	//	if err == nil {
	//		t.Errorf("wanted an error but do not get one")
	//	}
	//}
	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("do not get an error but wanted one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		//got := wallet.Balance()
		want := Bitcoin(20)
		//
		//if got != want {
		//	t.Errorf("got %s, want %s", got, want)
		//}
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		//got := wallet.Balance()
		want := Bitcoin(10)
		//
		//if got != want {
		//	t.Errorf("got %s, want %s", got, want)
		//}
		assertBalance(t, wallet, want)
		assertError(t, err, InsufficientFundsError)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		//assertError(t, err)
		//assertError(t, err, "cannot withdraw, insufficient funds")
		assertError(t, err, InsufficientFundsError)

		//if err == nil {
		//	t.Errorf("wanted an error but do not get one")
		//}
	})
}
