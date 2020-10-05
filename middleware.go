package cli

import (
	"context"
	"fmt"
	"log"
	"strings"
)

const (
	MwBeforeAll = "beforeall"
	MwBeforeProcess = "before"
	MwAfterProcess = "after"
)

type Middleware interface {
	Run(ctx context.Context, cli CLI) error
}

var (
	beforeAllMiddleware []Middleware
	beforeMiddleware []Middleware
	afterMiddleware []Middleware
)

func AddMiddleware(position string, handler Middleware) {
	switch strings.ToLower(position) {
	case MwBeforeAll:
		beforeAllMiddleware = append(beforeAllMiddleware, handler)
	case MwBeforeProcess:
		beforeMiddleware = append(beforeMiddleware, handler)
	case MwAfterProcess:
		afterMiddleware = append(afterMiddleware, handler)
	default:
		log.Fatal(fmt.Sprintf("cli: unknow middleware position %s", position))
	}
}

func callMiddleware(ctx context.Context, list []Middleware) (err error) {
	for _, m := range list {
		err = m.Run(ctx, c)
		if err != nil {
			return
		}
	}
	return
}