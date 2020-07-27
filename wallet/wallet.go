package wallet

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

/**
类型别名和类型再定义：
类型别名：type myInt = int 和原始类型几乎相同，两个类型之间的值可以相互赋值，具有相同的方法集。
类型再定义：type Bitcoin int 类型再定义，两者是不同的类型，相互之间不可赋值，具有不同的方法集。
*/
type Bitcoin int

type Wallet struct {
	//balance int
	balance Bitcoin
}

//func (w *Wallet) Deposit(amount int) {
//	//fmt.Printf("in deposit %p\n", &w)
//	w.balance = amount
//}
//
//func (w *Wallet) Balance() int {
//	return w.balance
//}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// 返回一个新的error
		//return errors.New("oh no")
		//return errors.New("cannot withdraw, insufficient funds")
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
