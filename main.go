package main

import (
	"Gin_OJ/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("oj.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := router.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
