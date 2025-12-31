package obs

import (
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

type SentryObs struct {
	hub    *sentry.Hub
	config config.ObsConfig
}

func NewSentryObs(config config.ObsConfig) (*SentryObs, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:            config.SENTRY_DSN,
		Debug:          true,
		SendDefaultPII: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup sentry")
	}

	sentryHub := sentry.CurrentHub()
	if sentryHub == nil {
		return nil, errors.New("failed to get sentry hub")
	}

	return &SentryObs{
		hub:    sentryHub,
		config: config,
	}, nil
}

// Flush should be called when the application is shutting down to ensure all events are sent
func (s *SentryObs) Flush() {
	sentry.Flush(2 * time.Second)
}

func (s *SentryObs) CaptureError(err error, status_code int) *sentry.EventID {
	var event_id *sentry.EventID

	// Capture error with metadata
	s.hub.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelError)
		scope.SetExtra("status_code", status_code)

		// Extract metadata from error chain
		type metadataError interface {
			GetMetadata() map[string]interface{}
		}

		var mdErr metadataError
		if errors.As(err, &mdErr) {
			metadata := mdErr.GetMetadata()
			for key, value := range metadata {
				scope.SetExtra(key, value)
			}
		}

		event_id = s.hub.CaptureException(err)
	})

	return event_id
}
