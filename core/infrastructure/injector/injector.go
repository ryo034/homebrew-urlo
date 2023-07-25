package injector

import (
	"urlo/core/driver"
	"urlo/core/infrastructure"
	"urlo/core/interface/controller"
	"urlo/core/interface/gateway"
	"urlo/core/interface/presenter"
	"urlo/core/usecase"
)

type Injector interface {
	Controller() *controller.Controller
}

type injector struct {
	controller *controller.Controller
}

func NewInjector(filePath string, cmdExecutor infrastructure.CommandExecutor, promptExecutor infrastructure.PromptExecutor) Injector {
	d := driver.NewDriver(filePath, cmdExecutor, promptExecutor)
	adapter := gateway.NewGatewayAdapter()
	gw := gateway.NewGateway(d, adapter)
	p := presenter.NewPresenter()
	interactor := usecase.NewInteractor(gw, p)
	return &injector{controller.NewController(interactor)}
}

func (i *injector) Controller() *controller.Controller {
	return i.controller
}
