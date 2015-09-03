package filter

import (
	"github.com/go-gonzo/filter/match"
	"github.com/omeid/gonzo"
	"github.com/omeid/gonzo/context"
)

func Filter(filter func(gonzo.File) bool) gonzo.Stage {
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {

		for {
			select {
			case file, ok := <-in:
				if !ok {
					return nil
				}
				if filter(file) {
					out <- file
				} else {
					err := file.Close()
					if err != nil {
						return err
					}
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

// Pick only files that match at least one of the provided patterns.
func Pick(patterns ...string) gonzo.Stage {
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {

		//Check patterns.
		err := match.Good(patterns...)
		if err != nil {
			return err
		}

		stage := Filter(
			func(file gonzo.File) bool {
				return match.Any(file.FileInfo().Name(), patterns...)
			},
		)

		return stage(ctx, in, out)
	}
}

// Drop files that match at least one of te patterns.
func Drop(patterns ...string) gonzo.Stage {
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {

		//Check patterns.
		err := match.Good(patterns...)
		if err != nil {
			return err
		}

		stage := Filter(func(file gonzo.File) bool { return !match.Any(file.FileInfo().Name(), patterns...) })

		return stage(ctx, in, out)
	}
}
