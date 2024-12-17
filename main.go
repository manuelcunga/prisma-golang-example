package main

import (
	"log"
	"prisma-golang/infra/http/server"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	srv := server.NewServer(engine)

	if err := srv.StartServer(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
