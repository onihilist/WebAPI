package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/onihilist/WebAPI/pkg/databases"

	// Repositories
	MiscRepository "github.com/onihilist/WebAPI/pkg/repositories/misc"
	UserRepository "github.com/onihilist/WebAPI/pkg/repositories/user"

	// Services
	MiscService "github.com/onihilist/WebAPI/pkg/services/misc"
	UserService "github.com/onihilist/WebAPI/pkg/services/user"

	// Controllers
	MiscController "github.com/onihilist/WebAPI/pkg/controllers/misc"
	UserController "github.com/onihilist/WebAPI/pkg/controllers/user"
)

type App struct {
	UserController *UserController.UserController
	MiscController MiscController.MiscController
	DB             *sql.DB
}

func InitializeApp(db *sql.DB) *App {

	// Initialize repositories
	userRepo := UserRepository.NewUserRepository(db)
	miscRepo := MiscRepository.NewMiscRepository(db)

	// Initialize services
	userService := UserService.NewUserService(userRepo)
	miscService := MiscService.NewMiscService(miscRepo)

	// Initialize controllers
	userController := UserController.NewUserController(userService)
	MiscController := MiscController.NewMiscController(userController, miscService)

	return &App{
		UserController: userController,
		MiscController: MiscController,
		DB:             db,
	}
}

func SetupRouter() *gin.Engine {

	// gin.DisableConsoleColor()

	db := databases.DatabaseConnect()

	databases.DatabaseHealthCheck(db)

	app := InitializeApp(db)
	return LoadRoutes(app)

}
