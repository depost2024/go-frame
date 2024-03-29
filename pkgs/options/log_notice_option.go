package options

import (
	"github.com/depost2024/go-frame/internal/components/lognotice"
	lognoticeI "github.com/depost2024/go-frame/pkgs/lognotice"
	"go.uber.org/fx"
)

// 错误日志企业微信通知
func WithLogNotice() Option {
	return func(o *Options) {
		o.FxOptions = append(o.FxOptions, fx.Provide(lognotice.New), fx.Invoke(lognoticeI.Inject))
	}
}
