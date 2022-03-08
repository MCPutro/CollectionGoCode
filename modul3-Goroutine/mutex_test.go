package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_sample_race_condition(t *testing.T) {
	//beberapa goroutine mengakases data/variable yg sama yg bersamaan,
	//dan di dalam goroutine itu data/variable yg di akses di manipulasi

	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

func Test_solution1_race_condition(t *testing.T) {
	//menggunakan mutex,
	//secara kasar akan seperti antrian goroutine
	//hanya 1 goroutine yg di perbolehkan untuk melakukan lock

	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)

}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (a *BankAccount) addBalance(amount int) {
	a.RWMutex.Lock() //lock func write
	a.Balance += amount
	a.RWMutex.Unlock() //unlock func write
}

func (a *BankAccount) getBalance() int {
	a.RWMutex.RLock() //lock func read
	balance := a.Balance
	a.RWMutex.RUnlock() //unlock func read
	return balance
}

func Test_RWMutex(t *testing.T) {
	realAccount := BankAccount{Balance: 0}
	for i := 1; i <= 100; i++ {
		go func() {
			for i := 1; i <= 100; i++ {
				realAccount.addBalance(1)
				fmt.Println(realAccount.getBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(realAccount)
}
