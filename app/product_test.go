package app_test

import (
	"testing"

	"github.com/nobruin/hexagonal-arch-example/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{
		Name:   "P1",
		Status: app.DISABLED,
		Price:  100,
	}
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
