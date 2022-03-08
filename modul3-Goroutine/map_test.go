package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func add2map(mapp *sync.Map, group *sync.WaitGroup, value int) {
	defer group.Done()

	group.Add(1)
	mapp.Store(value, value)
}

func Test_map(t *testing.T) {

	//tidak ada race condition, jika di akses oleh bnyak goroutine skaligus
	mapp := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go add2map(mapp, group, i)
	}

	group.Wait()

	mapp.Range(func(key, value interface{}) bool {
		fmt.Println(key, "-", value)
		return true //untuk melanjutkan looping, false untuk berhenti
	})

}
