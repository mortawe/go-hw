package calc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const EPS = 0.0001 //погрешность

type testInput struct {
	expr     string
	expected float64
}

func TestCalculateOK(t *testing.T) {
	test := []testInput{
		{
			expr:     "2",
			expected: 2.0,
		},
		{
			expr:     "2.5",
			expected: 2.5,
		},
		{
			expr:     "-7/2",
			expected: -3.5,
		},
		{
			expr:     "-(3.5)",
			expected: -3.5,
		},
		{
			expr:     "2 + 3",
			expected: 5.0,
		},
		{
			expr:     "2 + 3 * 5",
			expected: 17.0,
		},
		{
			expr:     "(2 + 3) * 5",
			expected: 25.0,
		},
		{
			expr:     "(2 * (3 + 2)) * 5",
			expected: 50.0,
		},
		{
			expr:     "(2 * (3 + 2)) * 5 + 5 / (3 - 1)",
			expected: 52.5,
		},
		{
			expr:     "1/2",
			expected: 0.5,
		},
		{
			expr:     "-(5 + 4)",
			expected: -9.0,
		},
		{
			expr:     "5 - (-(5.5 + 4))",
			expected: 14.5,
		},
		{
			expr:     "5 * (-(5 + 4))",
			expected: -45.0,
		},
		{
			expr:     "2/6*3",
			expected: 1.0,
		},
		{
			expr:     "(4 + 5 * (7 - 3)) - 2",
			expected: 22.0,
		},
		{
			expr:     "4+5+7/2",
			expected: 12.5,
		},
		{
			expr:     "-4*5",
			expected: -20.0,
		},
	}

	for r := range test {
		res, err := Calculate(test[r].expr)
		if err != nil {
			assert.Fail(t, "возникла ошибочка %s", err)
		}
		assert.Equal(t, math.Abs(res-test[r].expected) < EPS, true)
	}
}

func TestCalculateWrongInput(t *testing.T) {
	test := []string{
		"2/0",
		"(50",
		"50)",
		"(50--2)",
		"(50-)-2)",
		"(53+test+54w2)",
		"(-)1",
		"*1",
		"1/",
		"(5+8) * 9 - )1/2( ",
	}

	for r := range test {
		_, err := Calculate(test[r])
		assert.NotNil(t, err)
	}

}
