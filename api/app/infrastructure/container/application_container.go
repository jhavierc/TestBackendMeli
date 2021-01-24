package container

import "github.com/jhavierc/TestBackendMeli/api/app/application/handler"

func GetQuasarAplicationContainer() handler.QuasarHandlerApplication{
	return  &handler.QuasarHandler{GetQuasarServicePort: GetQuasarService()}
}
