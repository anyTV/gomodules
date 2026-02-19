package logger

import (
	"context"
	"fmt"
)

type ctxLogKey int

const loggerKey ctxLogKey = 0

type option struct {
	LogCtx string
	Level  levelType
}

type OptionFunc func(*option)

// using idgen.Generate() here will cause cyclic import
func WithLogIdCtx(id, name string) OptionFunc {
	return WithCtx(fmt.Sprintf("(log_id:%s) %s", id, name))
}

func WithCtx(ctx string) OptionFunc {
	return func(opt *option) {
		opt.LogCtx = ctx
	}
}

func WithLevel(level levelType) OptionFunc {
	return func(opt *option) {
		opt.Level = level
	}
}

func NewWithContext(ctx context.Context, opts ...OptionFunc) (context.Context, LogStruct) {
	opt := option{
		Level: GetLevel(),
	}

	for _, fn := range opts {
		fn(&opt)
	}

	log := New(opt.LogCtx, opt.Level)

	return context.WithValue(
		ctx,
		loggerKey,
		log,
	), log
}

func FromContext(ctx context.Context, defaultCtx string) LogStruct {
	log := ctx.Value(loggerKey)
	if log == nil {
		return New(defaultCtx, GetLevel())
	} else {
		return log.(LogStruct)
	}
}
