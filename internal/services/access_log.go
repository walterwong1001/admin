package services

import (
	"context"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/weitien/admin/internal/machine"
	"github.com/weitien/admin/internal/models"
	"github.com/weitien/admin/internal/repositories"
)

type AccessLogService interface {
	Log(ctx context.Context, metric map[string]any) error
}

type accessLogServiceImpl struct {
	repository repositories.AccessLogRepository
}

func NewAccessLogService() AccessLogService {
	return &accessLogServiceImpl{repositories.NewAccessLogRepository()}
}

// Log Append implements AccessLogService.
func (s *accessLogServiceImpl) Log(ctx context.Context, metric map[string]any) error {
	accLog := &models.AccessLog{}
	err := mapstructure.Decode(metric, accLog)
	if err != nil {
		log.Println("Decode error", err)
	}

	accLog.ID = machine.NextID()

	return s.repository.Append(ctx, repositories.GetDB(), accLog)
}
