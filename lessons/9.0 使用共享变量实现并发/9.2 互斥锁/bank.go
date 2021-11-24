package main

import "sync"

//var (
//	sema    = make(chan struct{}, 1)
//	balance int
//)
//
//func Despoit(amount int) {
//	sema <- struct{}{}
//	balance = balance + amount
//	<-sema
//}
//
//func Balance() int {
//	sema <- struct{}{}
//	b := balance
//	<-sema
//	return b
//}

var (
	mu      sync.Mutex
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)

}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {

}
