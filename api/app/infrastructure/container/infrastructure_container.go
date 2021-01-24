package container

import (
	"database/sql"
	"fmt"
	"github.com/jhavierc/TestBackendMeli/api/app/infrastructure/controller"
	"github.com/jhavierc/TestBackendMeli/goutils/logger"
)

func GetQuasarController() *controller.GetQuasarController{
	return  &controller.GetQuasarController{GetQuasarApplication: GetQuasarAplicationContainer()}
}

//not used bd in this project
func GetReadConnectionClient() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", "root", "1234", "localhost:3306", "bd_test")
	clientDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Printf("Error BD: %v", err)
	}
	err = clientDB.Ping()
	if err != nil {
		fmt.Printf("Error conectando a BD: %v", err)

	}
	logger.Info(fmt.Sprint("=====>Conectado correctamente a la BD"))
	return clientDB
}
