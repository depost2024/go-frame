/*
 * @Date: 2022-04-24 09:56:04
 * @LastEditTime: 2022-04-29 18:32:17
 * @FilePath: /github.com/depost2024/go-frame/pkgs/options/options.go
 */

package options

import (
	"go.uber.org/fx"
)

type (
	Options struct {
		ConfigFile      string      // 配置文件路径
		FxOptions       []fx.Option // 需要注入的对象
		IsStart         bool        // true=>运行一次就退出
		UseHttp         bool        // 是否使用http
		UseGraceRestart bool        // 是否使用优雅重启
	}
	Option func(*Options) // 选项函数类型
)

// 注册需要提供的对象
func WithProviders(providers ...any) Option {
	return func(o *Options) {
		for _, provider := range providers {
			o.FxOptions = append(o.FxOptions, fx.Provide(provider))
		}
	}
}

// 注册需要被调用的函数
func WithInvokes(Invokes ...any) Option {
	return func(o *Options) {
		for _, Invoke := range Invokes {
			o.FxOptions = append(o.FxOptions, fx.Invoke(Invoke))
		}
	}
}

// 一次注册单个fx.Option
func WithFxOption(fxOptions ...fx.Option) Option {
	return func(o *Options) {
		o.FxOptions = append(o.FxOptions, fxOptions...)
	}
}

// 一次注册多个fx.Option
func WithFxOptions(fxOptions ...[]fx.Option) Option {
	return func(o *Options) {
		for _, op := range fxOptions {
			o.FxOptions = append(o.FxOptions, op...)
		}
	}
}
