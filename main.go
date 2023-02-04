package main

import (
	"fmt"

	"log"
	"net"
	"os"

	"github.com/hasrulrhul/grpc-go-gin/model"
	userHandler "github.com/hasrulrhul/grpc-go-gin/user/handler"
	userRepo "github.com/hasrulrhul/grpc-go-gin/user/repo"
	userUsecase "github.com/hasrulrhul/grpc-go-gin/user/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var port = os.Getenv("APP_PORT")
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(
		model.UserDB{},
	)

	server := grpc.NewServer()

	userRepo := userRepo.CreateUserRepoImpl(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)

	userHandler.CreateUserHandler(server, userUsecase)

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Starting at Port: ", port)
	log.Fatal(server.Serve(conn))
}
