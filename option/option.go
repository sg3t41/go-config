package option

import (
	"log"
	"os"
	"strings"

	"github.com/sg3t41/gocnf/enum"
)

// Optionはアプリケーションの実行モードと、モードに対応するファイルパスを保持する構造体
type Option struct {
	DefaultFilePath string
	RunMode         enum.RunMode
	ModeToFilePath  map[enum.RunMode]string
}

func NewOption() *Option {
	option := &Option{
		ModeToFilePath: make(map[enum.RunMode]string),
	}
	return option
}

func (opt *Option) SetDefaultFilePath(path string) *Option {
	opt.DefaultFilePath = path
	return opt
}

// SetModeToFilePathは指定したモードとファイルパスのペアを設定します
func (opt *Option) SetModeToFilePath(mode enum.RunMode, path string) *Option {
	opt.ModeToFilePath[mode] = path
	return opt
}

// SetRunModeはOptionのRunModeを設定します
func (opt *Option) SetRunMode(mode enum.RunMode) *Option {
	opt.RunMode = mode
	return opt
}

// LoadCurrentRunModeは環境変数から実行モードを読み込み、適切なモードをセットします
func (opt *Option) SetCurrentRunMode(key string) *Option {
	mode := os.Getenv(key)
	printMode := mode
	if printMode == "" {
		log.Printf("【INFO】実行モード($%s)が設定されていません。", key)
	} else {
		log.Printf("【INFO】実行モード [%s] で設定ファイルを読み込みます。\n", printMode)
	}
	switch {
	case strings.EqualFold(mode, enum.Development.String()):
		opt.RunMode = enum.Development
	case strings.EqualFold(mode, enum.Staging.String()):
		opt.RunMode = enum.Staging
	case strings.EqualFold(mode, enum.Production.String()):
		opt.RunMode = enum.Production
	default:
		log.Printf("【INFO】実行モードの設定値 [%s] は不正です。 既定値は[development], [staging], [production] のいずれかです。\n", mode)
		opt.RunMode = enum.Default
	}
	return opt
}
