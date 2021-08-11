package mpesa

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/mpesa/pkg/models"
	"log"
)

var _ Service = (*LoggerAdapter)(nil)
var _ Service = (*FormatAdapter)(nil)

type (
	Adapter func(Service) Service

	LoggerAdapter struct {
		Logger *log.Logger
		Next   Service
	}

	FormatAdapter struct {
		Next Service
	}
)

func (f *FormatAdapter) SessionID(ctx context.Context, platform Platform, market Market) (response models.SessionResponse, err error) {
	fmt.Printf("just acknowledging the hustle in getting the session ID\n")
	return f.Next.SessionID(ctx,platform,market)
}

func (l *LoggerAdapter) SessionID(ctx context.Context, platform Platform, market Market) (response models.SessionResponse, err error) {
	l.Logger.Printf("getting session id from hommies")
	return l.Next.SessionID(ctx,platform,market)
}

func Adapt(service Service, adapters ...Adapter) Service{
	for _, adapter := range adapters {
		service = adapter(service)
	}
	return service
}