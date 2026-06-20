package ports

import "context"

type InboxRepository interface {
	JaProcessado(
		ctx context.Context,
		eventID string,
	) (bool, error)

	MarcarComoProcessado(
		ctx context.Context,
		eventID string,
	) error
}
