package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func UpdateProduct(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	uid := ctx.Params("uid")

	// converte json para struct
	productBody, err := product.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS090", "O formado dos dados envidados está incorreto."))
	}

	// verifica e compara o produto recebido
	if productBody.Uid != "" && uid != productBody.Uid {
		log.Error().Msgf("Não é possível atualizar o produto %v para o %v", uid, productBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS091", "Não é possível atualizar o produto "+uid+" para o "+productBody.Uid))
	}

	// carrega o produto da base de dados
	productDatabase, err := product.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar o produto %v.", productBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS092", "Erro ao tentar validar o produto."))
	}

	// injecta dos dados novos no lugar dos dados trazidos da base de dados
	productDatabase.Inject(productBody)
	productDatabase.UpdatedAt = date.NowUTC()
	productDatabase.Uid = uid

	// guarda as alterações do produto na base de dados
	err = product.Repository().Update(productDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do produto (%v)", productBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS093"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
