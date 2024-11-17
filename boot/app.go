package boot

import "golang.org/x/net/context"

type Service interface {
	OnStart(ctx context.Context)
	OnDelete(ctx context.Context)
}

type App struct {
	services []Service
}
