package api

import (
	"net/http"
	"shop-product-service/internal/repositories"
	"shop-product-service/internal/schemas"
	"shop-product-service/internal/services"
	"shop-product-service/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
	"github.com/samber/lo"
)

func CreateProductHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.CreateProductSchema{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := cartpkg.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	productService := services.ProductService{DB: appState.DB}
	if err := productService.CreateProduct(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}

func GenerateRandomProductHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.GenerateProductSchema{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := cartpkg.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	generateProduct := schemas.CreateProductSchema{
		Name:        faker.Name(),
		Description: lo.RandomString(10, lo.AlphanumericCharset),
		Price:       lo.Sample([]float32{11, 22, 33, 44, 55, 66, 77, 88, 99}),
		Stock:       lo.Sample([]uint{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		ShopId:      payload.ShopId,
	}

	productRepository := repositories.ProductRepository{DB: appState.DB}
	if err := productRepository.CreateProduct(generateProduct); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, http.StatusCreated)
}
