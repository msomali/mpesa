/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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

func (f *FormatAdapter) C2BSingleAsync(ctx context.Context, request models.PushRequest) (models.C2BSingleStageAsyncResponse, error) {
	fmt.Printf("c2b from formatter")
	return f.Next.C2BSingleAsync(ctx, request)
}

func (l *LoggerAdapter) C2BSingleAsync(ctx context.Context, request models.PushRequest) (models.C2BSingleStageAsyncResponse, error) {
	l.Logger.Printf("making push payyyyy")
	return l.Next.C2BSingleAsync(ctx, request)
}

func (f *FormatAdapter) SessionID(ctx context.Context) (response models.SessionResponse, err error) {
	fmt.Printf("just acknowledging the hustle in getting the session ID\n")
	return f.Next.SessionID(ctx)
}

func (l *LoggerAdapter) SessionID(ctx context.Context) (response models.SessionResponse, err error) {
	l.Logger.Printf("getting session id from hommies")
	return l.Next.SessionID(ctx)
}

func Adapt(service Service, adapters ...Adapter) Service {
	for _, adapter := range adapters {
		service = adapter(service)
	}
	return service
}
