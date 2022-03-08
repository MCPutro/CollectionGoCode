package table_test

import (
	modul2_testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pertama(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "test1",
			request:  "test1",
			expected: "hi,test1",
		},
		{
			name:     "test2",
			request:  "test2",
			expected: "hi,test2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := modul2_testing.SayHi(test.name)
			assert.Equal(t, result, test.expected)
		})
	}
}
