package services

import (
	"context"

	"github.com/walterwong1001/gin_common_libs/pkg/page"
)

type Paginator[T, S any] interface {
	Pagination(ctx context.Context, p page.Paginator[T], filter S) error
}
