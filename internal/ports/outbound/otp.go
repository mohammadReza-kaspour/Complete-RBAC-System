package outbound

import "context"

type Generator interface {
	Generate(ctx context.Context, length int) (string, error)
}

type Sender interface {
	Send(ctx context.Context, phone string, code string) error
}
