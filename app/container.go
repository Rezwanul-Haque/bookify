package container

import (
	"bookify/app/http/controllers"
	repoImpl "bookify/app/repository/impl"
	svcImpl "bookify/app/svc/impl"
	"bookify/infra/conn/db"
	"bookify/infra/logger"
	"context"
)

func Init(g interface{}, lc logger.LogClient) {
	basectx := context.Background()
	dbc := db.Client()

	// register all repos impl, services impl, controllers
	sysRepo := repoImpl.NewSystemRepository(basectx, dbc)
	userRepo := repoImpl.NewBookRepository(basectx, dbc)

	sysSvc := svcImpl.NewSystemService(sysRepo)
	userSvc := svcImpl.NewUsersService(basectx, lc, userRepo)

	controllers.NewSystemController(g, sysSvc)
	controllers.NewBookController(g, userSvc)

}
