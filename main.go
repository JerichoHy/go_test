package main

import (
	"awesomeProject/model"
	"awesomeProject/route"
	"awesomeProject/utils"
)

func main() {
	model.InitDatabase()
	r := route.InitRouter()
	r.Run(utils.C.Server.Port)
}
