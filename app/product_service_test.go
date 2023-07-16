package app_test

import (
	"github.com/golang/mock/gomock"
	"github.com/nobruin/hexagonal-arch-example/app"
	mockApp "github.com/nobruin/hexagonal-arch-example/app/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockApp.NewMockProductInterface(ctrl)
	persistence := mockApp.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	productService := app.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Get("ID")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockApp.NewMockProductInterface(ctrl)
	persistence := mockApp.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	productService := app.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Create("p1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockApp.NewMockProductInterface(ctrl)
	persistence := mockApp.NewMockProductPersistenceInterface(ctrl)

	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	productService := app.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result1, err1 := productService.Disable(product)
	require.Nil(t, err1)
	require.Equal(t, product, result1)
}
