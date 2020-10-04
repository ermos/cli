package help

import (
	"context"
	"fmt"
	"github.com/ermos/gocli"
)

type Handler struct {}

func (Handler) Description(cli gocli.CLI) string {
	return fmt.Sprintf("show all commands and their utility")
}

func (Handler) Run(ctx context.Context, cli gocli.CLI) error {
	return nil
}