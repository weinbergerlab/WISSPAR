package run

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type ErrFunc func() error

func Go(ctx context.Context, functions...ErrFunc) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, function := range functions {
		g.Go(function)
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
