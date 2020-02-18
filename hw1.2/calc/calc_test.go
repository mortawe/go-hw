package calc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)
const EPS = 0.0001 //погрешность

func TestCalculateOK(t *testing.T) {
	res, err := Calculate("2")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 2.0)

	res, err = Calculate("2.5")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 2.5)

	res, err = Calculate("-7/2")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, -3.5, res)

	res, err = Calculate("-(3.5)")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, -3.5, res)

	res, err = Calculate("2 + 3")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 5.0)

	res, err = Calculate("2 + 3 * 5")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 17.0)

	res, err = Calculate("(2 + 3) * 5")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 25.0)

	res, err = Calculate("(2 * (3 + 2)) * 5")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 50.0)

	res, err = Calculate("(2 * (3 + 2)) * 5 + 5 / (3 - 1)")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 52.5)

	res, err = Calculate("1/2")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 0.5)

	res, err = Calculate("-(5 + 4)")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, -9.0)

	res, err = Calculate("5 - (-(5.5 + 4))")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, 14.5)

	res, err = Calculate("5 * (-(5 + 4))")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, res, -45.0)

	res, err = Calculate("2/6*3")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, math.Abs(res - 1) < EPS, true)

	res, err = Calculate("(4 + 5 * (7 - 3)) - 2")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, 22.0, res)

	res, err = Calculate("4+5+7/2")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, 12.5, res)

	res, err = Calculate("-4*5")
	if err != nil {
		assert.Fail(t, "возникла ошибочка %s", err)
	}
	assert.Equal(t, -20.0, res)

}

func TestCalculateWrongInput(t *testing.T) {

	_, err := Calculate("2/0")
	assert.NotNil(t, err)

	_, err = Calculate("(50")
	assert.NotNil(t, err)

	_, err = Calculate("50)")
	assert.NotNil(t, err)

	_, err = Calculate("(50--2)")
	assert.NotNil(t, err)

	_, err = Calculate("(50-)-2)")
	assert.NotNil(t, err)

	_, err = Calculate("(53+test+54w2)")
	assert.NotNil(t, err)

	_, err = Calculate("(-)1")
	assert.NotNil(t, err)

	_, err = Calculate("*1")
	assert.NotNil(t, err)

	_, err = Calculate("1/")
	assert.NotNil(t, err)

	_, err = Calculate("(5+8) * 9 - )1/2( ")
	assert.NotNil(t, err)
}
