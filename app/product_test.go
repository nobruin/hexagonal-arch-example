package app_test

import (
	"testing"

	"github.com/nobruin/hexagonal-arch-example/app"
	uuid "github.com/satori/go.uuid"
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
}

func TestProduct_Enable_Error(t *testing.T) {
	product := app.Product{
		Name:   "P1",
		Status: app.DISABLED,
		Price:  0,
	}

	err := product.Enable()
	require.Equal(t, app.ENABLED_ERROR_MESSAGE, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{
		Name:   "P1",
		Status: app.ENABLED,
		Price:  0,
	}
	err := product.Disable()
	require.Nil(t, err)
}

func TestProduct_Disable_Error(t *testing.T) {
	product := app.Product{
		Name:   "P1",
		Status: app.ENABLED,
		Price:  100,
	}

	err := product.Disable()
	require.Equal(t, app.DISABLED_ERROR_MESSAGE, err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, app.STATUS_INVALID_ERROR_MESSAGE, err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, app.ISVALID_PRICE_ERROR_MESSAGE, err.Error())
}
