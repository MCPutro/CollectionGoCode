package sub_test

import (
	modul2_testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SubTest(t *testing.T) {
	t.Run("subtest1", func(t *testing.T) {
		result := modul2_testing.SayHi("kuro")
		assert.Equal(t, result, "hi,kuro")
	})

	t.Run("subtest2", func(t *testing.T) {
		result := modul2_testing.SayHi("kuro")
		assert.Equal(t, result, "hi,kuro")
	})

}
