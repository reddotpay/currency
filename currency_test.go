package currency_test

import (
	"testing"

	"github.com/darylnwk/currency"
	"github.com/stretchr/testify/assert"
)

func TestCurrency_New(t *testing.T) {
	sgd, err := currency.New("SGD")
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal("SGD", sgd.AlphaCode)
	assert.Equal("Singapore dollar", sgd.Name)
	assert.Equal(2, sgd.Decimal)
}

func TestCurrency_New_LowerCase(t *testing.T) {
	sgd, err := currency.New("sgd")
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal("SGD", sgd.AlphaCode)
	assert.Equal("Singapore dollar", sgd.Name)
	assert.Equal(2, sgd.Decimal)
}

func TestCurrency_New_InvalidCurrency(t *testing.T) {
	sgd, err := currency.New("XXX")
	assert := assert.New(t)
	assert.Nil(sgd)
	assert.EqualError(err, "currency: code is not a valid ISO 4217 alphabetic code")
}
