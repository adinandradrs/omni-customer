package main

import (
	"os"

	sbase "github.com/adinandradrs/codefun-go-service"
	config "github.com/adinandradrs/omni-customer/configuration"
	"github.com/adinandradrs/omni-customer/controller"
	"github.com/adinandradrs/omni-customer/repository"
	"github.com/adinandradrs/omni-customer/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//load os variables [start]
	dburl := os.Getenv("dburl")
	dbhost := os.Getenv("dbhost")
	dbport := os.Getenv("dbport")
	dbuser := os.Getenv("dbuser")
	dbpass := os.Getenv("dbpass")
	dbsch := os.Getenv("dbschema")

	//load base configuration [start]
	db := config.ConfigDatabase(dbhost, dburl, dbport, dbuser, dbpass, dbsch)

	//load dependency injection [start]
	srepo := sbase.NewBaseRepository(db)
	crepo := repository.NewCustomerRepository(*srepo)
	cact := service.NewCustomerActivation(crepo)
	creq := service.NewCustomerRegister(crepo)
	cctrl := controller.NewCustomerController(cact, creq)

	route := gin.Default()
	route.POST("/v1/register", cctrl.Register)
	route.POST("/v1/activate", cctrl.Activate)
	route.Run(":9000")
}
