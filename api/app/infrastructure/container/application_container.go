package container

import "github.com/TestBackendMeli/api/app/application/handler"

func GetQuasarAplicationContainer() handler.QuasarHandlerApplication{
	return  &handler.QuasarHandler{GetQuasarServicePort: GetQuasarService()}
}
