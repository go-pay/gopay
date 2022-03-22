package xlog

import (
	"context"
	"testing"

	"github.com/go-pay/gopay/pkg/errgroup"
)

func TestLog(t *testing.T) {
	var eg errgroup.Group

	// default log
	eg.Go(func(ctx context.Context) error {
		Info(White, "白色 info", Reset, WhiteBright, "高亮 info", Reset, "恢复默认颜色", WhiteDelLine, "删除线", Reset, WhiteUnderLine, "下划线", Reset, WhiteBevel, "斜体 info", Reset, WhiteBg, "背景", Reset)
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Debug(Cyan, "青色 debug", Reset, CyanBright, "高亮 debug", Reset, "恢复默认颜色", CyanDelLine, "删除线", Reset, CyanUnderLine, "下划线", Reset, CyanBevel, "斜体 debug", Reset, CyanBg, "背景", Reset)
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Warn(Yellow, "黄色 warning", Reset, YellowBright, "高亮 warning", Reset, "恢复默认颜色", YellowDelLine, "删除线", Reset, YellowUnderLine, "下划线", Reset, YellowBevel, "斜体 warning", Reset, YellowBg, "背景", Reset)
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Error(Red, "红色 error", Reset, RedBright, "高亮 error", Reset, "恢复默认颜色", RedDelLine, "删除线", Reset, RedUnderLine, "下划线", Reset, RedBevel, "斜体 error", Reset, RedBg, "背景", Reset)
		return nil
	})

	// color log
	eg.Go(func(ctx context.Context) error {
		Color(White).Info("color log info")
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Color(Cyan).Debug("color log debug")
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Color(Yellow).Warn("color log warn")
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		Color(Red).Error("color log error")
		return nil
	})
	_ = eg.Wait()
}
