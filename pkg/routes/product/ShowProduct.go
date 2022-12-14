package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// mostra os dados de um produto
func ShowProduct(ctx *fiber.Ctx) error {
	var log = logger.New()

	uid := ctx.Params("uid")

	// carega um produto da base de dados
	productData, err := product.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS089"))
	}

	return response.Ctx(ctx).Result(response.Success(200, productData))

}
