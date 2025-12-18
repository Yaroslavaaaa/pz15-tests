package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum_Table(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative number", 10, -5, 5},
		{"both zero", 0, 0, 0},
		{"both negative", -4, -6, -10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.a, tc.b)
			if got != tc.expected {
				t.Fatalf(
					"Sum(%d, %d) = %d; want %d",
					tc.a, tc.b, got, tc.expected,
				)
			}
		})
	}
}

func TestDivide_OkAndError(t *testing.T) {
	t.Run("ok division", func(t *testing.T) {
		got, err := Divide(10, 2)
		require.NoError(t, err)
		assert.Equal(t, 5, got)
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		assert.Error(t, err)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sum(123, 456)
	}
}

func TestPanicDivide(t *testing.T) {
	require.Panics(t, func() {
		PanicDivide(10, 0)
	}, "ожидаем панику при делении на 0")

	// Проверка обычного деления не паникит
	require.NotPanics(t, func() {
		PanicDivide(10, 2)
	})
}
