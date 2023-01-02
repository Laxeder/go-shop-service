package routes

import (
	"github.com/gofiber/fiber/v2"
	mid "github.com/laxeder/go-shop-service/pkg/middlewares"
	acc "github.com/laxeder/go-shop-service/pkg/routes/account"
	addr "github.com/laxeder/go-shop-service/pkg/routes/address"
	ctg "github.com/laxeder/go-shop-service/pkg/routes/category"
	fgt "github.com/laxeder/go-shop-service/pkg/routes/freight"
	prod "github.com/laxeder/go-shop-service/pkg/routes/product"
	"github.com/laxeder/go-shop-service/pkg/routes/redis"
	spc "github.com/laxeder/go-shop-service/pkg/routes/shopcart"
	usr "github.com/laxeder/go-shop-service/pkg/routes/user"
)

func ApiV1(app *fiber.App) {

	route := app.Group("/api/v1")

	//? ****************************** Rotas do servidor *****************************

	route.Get("/health", Health)
	route.Get("/redis/health", redis.RedisHealth)

	//? ****************************** Rotas do usuário ******************************

	route.Put("/user/password/:uuid", mid.JWT, usr.UpdateUserPassword)
	route.Put("/user/document/:uuid", mid.JWT, usr.UpdateUserDocument)
	route.Post("/user/login", usr.LoginUser)
	route.Patch("/user/:uuid", mid.JWT, usr.RestoreUser)
	route.Delete("/user/:uuid", mid.JWT, usr.DeleteUser)
	route.Put("/user/:uuid", mid.JWT, usr.UpdateUser)
	route.Get("/user/:uuid", mid.JWT, usr.ShowUser)
	route.Post("/user", usr.CreateUser)
	route.Get("/users", mid.JWT, usr.ListUsers)
	route.Patch("/address/:uid", mid.JWT, addr.RestoreAddress)

	//? ****************************** Rotas do endereço *****************************

	route.Delete("/address/:uid", mid.JWT, addr.DeleteAddress)
	route.Put("/address/:uid", mid.JWT, addr.UpdateAddress)
	route.Get("/address/:uid", mid.JWT, addr.ShowAddress)
	route.Post("/address", mid.JWT, addr.CreateAddress)
	route.Get("/adresses", mid.JWT, addr.ListAddress)

	//? ******************************** Rotas da conta ******************************

	route.Patch("/account/:uid", mid.JWT, acc.RestoreAccount)
	route.Delete("/account/:uid", mid.JWT, acc.DeleteAccount)
	route.Put("/account/:uid", mid.JWT, acc.UpdateAccount)
	route.Get("/account/:uid", mid.JWT, acc.ShowAccount)
	route.Post("/account", mid.JWT, acc.CreateAccount)
	route.Get("/accounts", mid.JWT, acc.ListAccounts)

	//? ******************************* Rotas do produto *****************************

	route.Patch("/product/:uid", mid.JWT, prod.RestoreProduct)
	route.Delete("/product/:uid", mid.JWT, prod.DeleteProduct)
	route.Put("/product/:uid", mid.JWT, prod.UpdateProduct)
	route.Get("/product/:uid", mid.JWT, mid.PermissionsCheck("getProduct"), prod.ShowProduct)
	route.Post("/product", mid.JWT, prod.CreateProduct)
	route.Get("/products", mid.JWT, prod.ListProducts)

	//? ****************************** Rotas da categoria ******************************

	route.Patch("/category/:code", mid.JWT, ctg.RestoreCategory)
	route.Delete("/category/:code", mid.JWT, ctg.DeleteCategory)
	route.Put("/category/:code", mid.JWT, ctg.UpdateCategory)
	route.Get("/category/:code", mid.JWT, ctg.ShowCategory)
	route.Post("/category", mid.JWT, ctg.CreateCategory)
	route.Get("/categories", mid.JWT, ctg.ListCategories)

	//? ************************* Rotas do carrinho de compras *************************

	route.Patch("/shopcart/:uuid", mid.JWT, spc.RestoreShopCart)
	route.Delete("/shopcart/:uuid", mid.JWT, spc.DeleteShopCart)
	route.Put("/shopcart/:uuid", mid.JWT, spc.UpdateShopCart)
	route.Post("/shopcart/:uuid/product", mid.JWT, spc.AddShopCartProduct)
	route.Delete("/shopcart/:uuid/product", mid.JWT, spc.RemoveShopCartProduct)

	route.Get("/shopcart/:uuid", mid.JWT, spc.ShowShopCart)
	route.Post("/shopcart", mid.JWT, spc.CreateShopCart)
	route.Get("/shopcarts", mid.JWT, spc.ListShopCarts)

	//? ******************************** Rotas do frete ********************************

	route.Patch("/freight/:uid", mid.JWT, fgt.RestoreFreight)
	route.Delete("/freight/:uid", mid.JWT, fgt.DeleteFreight)
	route.Put("/freight/:uid", mid.JWT, fgt.UpdateFreight)
	route.Get("/freight/:uid", mid.JWT, fgt.ShowFreight)
	route.Post("/freight", mid.JWT, fgt.CreateFreight)
	route.Get("/freights", mid.JWT, fgt.ListFreights)

}

func ErrorNotFound(app *fiber.App) {
	app.Use("/", NotFound)
}
