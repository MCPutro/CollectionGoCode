package modul1

import (
	"fmt"
	"sort"
	"testing"
)

type murid struct {
	nama  string
	umur  int
	kelas int
}

type dataMurid []murid

func (d dataMurid) Len() int {
	return len(d)
}

func (d dataMurid) Less(i, j int) bool {
	return d[i].nama > d[j].nama
}

func (d dataMurid) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func TestSort(t *testing.T) {
	murid2 := []murid{
		{nama: "a", umur: 5},
		{nama: "a", umur: 5},
		{nama: "a", umur: 5},
	}

	fmt.Println(dataMurid(murid2).Len())

	sort.Sort(dataMurid(murid2))
	fmt.Println(murid2)

}
