package function

import (
	"github.com/pkg/errors"
	"time"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	for i := 0; i < attempts; i++ {
		if err := fn(); err != nil {
			if i == attempts-1 {
				return err
			}
			time.Sleep(sleep)
			sleep *= 2
			continue
		}
		return nil
	}
	return errors.New("maximum retries reached")
}
