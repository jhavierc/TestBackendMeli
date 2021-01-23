package container

import "github.com/TestBackendMeli/api/app/domain/service"

//llama los respositorios
func GetQuasarService() service.GetQuasarServicePort{
	return &service.GetQuasarContentService{}
}
