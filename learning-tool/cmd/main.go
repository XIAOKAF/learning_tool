package main

import (
	"learning-tool/api"
	"learning-tool/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
	api.Server()
	api.Client()
}