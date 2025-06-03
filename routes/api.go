package routes

import "github.com/aldysp34/deeptech-test/controllers"

func (a *Routes) ListRouter() {
	apiRouter := a.Router.PathPrefix("/api").Subrouter()

	controllers.NewAdminRouter(apiRouter)
	controllers.NewCategoryRouter(apiRouter)
	controllers.NewProductRouter(apiRouter)
	controllers.NewTransactionRouter(apiRouter)
}
