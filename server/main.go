package main

import (
	"fmt"
	"github.com/jackyuan2010/gpaas/server/appcontext"
	"github.com/jackyuan2010/gpaas/server/utils"
)

func main() {
	fmt.Println("gpaas app starting....")
	appcontext.InitAppContext()

	jwtUtil := utils.NewJWTUtil()
	claims := jwtUtil.CreateClaims("13311221122", "jackyuan2010")

	fmt.Println(claims)

	// config.Viper()

	// initDB()
	router := appcontext.Routers()
	router.Run(":8081")
}
