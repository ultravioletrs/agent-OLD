// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

//go:build !test
// +build !test

package api

import (
	"context"
	"fmt"
	"time"

	log "github.com/mainflux/mainflux/logger"
	"github.com/ultravioletrs/agent/agent"
)

var _ agent.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger log.Logger
	svc    agent.Service
}

// LoggingMiddleware adds logging facilities to the core service.
func LoggingMiddleware(svc agent.Service, logger log.Logger) agent.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) Run(ctx context.Context, cmp agent.Computation) (response string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method Run for computation %s took %s to complete", cmp.ID, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.Run(ctx, cmp)
}

func (lm *loggingMiddleware) Algo(ctx context.Context, algorithm []byte) (response string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method Algo took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors", message))
	}(time.Now())

	return lm.svc.Algo(ctx, algorithm)
}

func (lm *loggingMiddleware) Data(ctx context.Context, dataset []byte) (response string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method Data took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors", message))
	}(time.Now())

	return lm.svc.Data(ctx, dataset)
}

func (lm *loggingMiddleware) Result(ctx context.Context) (response []byte, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method Result took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors", message))
	}(time.Now())

	return lm.svc.Result(ctx)
}

func (lm *loggingMiddleware) Attestation(ctx context.Context) (response []byte, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method Attestation took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors", message))
	}(time.Now())

	return lm.svc.Attestation(ctx)
}
