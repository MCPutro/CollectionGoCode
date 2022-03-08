package bencmark

import (
	modul2_testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"testing"
)

func BenchmarkCoba(b *testing.B) {
	for i := 0; i < b.N; i++ {
		modul2_testing.SayHi("asd")
	}
}
