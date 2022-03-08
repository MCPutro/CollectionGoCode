package bencmark

import (
	modul2_testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"testing"
)

func BenchmarkSubCoba(b *testing.B) {
	b.Run("sub1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			modul2_testing.SayHi("asd")
		}
	})

	b.Run("sub2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			modul2_testing.SayHi("asd")
		}
	})

}
