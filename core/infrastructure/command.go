package infrastructure

import (
	"os/exec"
	"urlo/core/domain"
)

type CommandExecutor interface {
	Open(u domain.StrictUrl) error
}

type commandExecutor struct{}

func NewCommandExecutor() CommandExecutor {
	return &commandExecutor{}
}

func (c *commandExecutor) Open(u domain.StrictUrl) error {
	return exec.Command("open", u.String()).Start()
}
