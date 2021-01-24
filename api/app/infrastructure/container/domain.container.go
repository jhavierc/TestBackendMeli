package container

import "github.com/jhavierc/TestBackendMeli/api/app/domain/service"

//llama los respositorios
func GetQuasarService() service.GetQuasarServicePort{
	return &service.GetQuasarContentService{}
}
