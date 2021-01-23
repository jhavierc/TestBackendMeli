package container

import "github.com/TestBackendMeli/api/app/application/handler"

func GetQuasarContainer() handler.QuasarHandlerContentApplication{
	return  &handler.GetGlobalContentQuasarService{GetQuasarContentServicePort: GetQuasarService()}
}
