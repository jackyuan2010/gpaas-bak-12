package main

import (
	"fmt"
	gpaascontext "github.com/jackyuan2010/gpaas/server/context"
	"github.com/jackyuan2010/gpaas/server/core"
	"github.com/jackyuan2010/gpaas/server/utils"
)

func main() {
	fmt.Println("gpaas app starting....")
	gpaascontext.InitAppContext()

	jwtUtil := utils.NewJWTUtil()
	claims := jwtUtil.CreateClaims("13311221122", "jackyuan2010")

	fmt.Println(claims)

	// config.Viper()

	// initDB()
	router := core.Routers()
	router.Run(":8081")
}
