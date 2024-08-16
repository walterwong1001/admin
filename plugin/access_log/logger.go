package accesslog

import (
	"context"
	"fmt"
)

type Logger interface {
	Log(ctx context.Context, metric map[string]any)
}

type ConsoleLog struct{}

func (l *ConsoleLog) Log(ctx context.Context, metric map[string]any) {
	for k, v := range metric {
		fmt.Printf("%s: %v \n", k, v)
	}
}

type LogAppender interface {
	Append(ctx context.Context, metric map[string]any) error
}

type DBLog struct {
	Appender LogAppender
}

func (l *DBLog) Log(ctx context.Context, metric map[string]any) {
	l.Appender.Append(ctx, metric)
}
