package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	logger := NewLogger(gin.Mode())
	logger.Debug("replaced zap's global loggers", zap.String("memberID", "1231"), zap.Any("member", ""))
	fmt.Println(gin.Mode())
	//defer logger.Sync()
	//sugar := logger.Sugar()
	//sugar.Debug("debug message")
	//sugar.Info("info message")
	//sugar.Error("error message")
	//sugar.Warn("warn message")
	//sugar.Panic("panic message")
	//sugar.Fatal("fatal message")
	//undo := zap.ReplaceGlobals(logger)
	//defer undo()

	//logger.Info("failed to fetch URL",
	//	zap.String("url", "http://abc.com/get"),
	//	zap.Int("attempt", 3),
	//	zap.Bool("enabled", true),
	//	zap.Any("", ""),
	//)
	//b := &a{}
	//b.ID = "123"
	//zap.L().Info("replaced zap's global loggers", zap.String("memberID", "123"), zap.Any("member", b))
	//name()
}

func name() {
	c := &a{}
	c.ID = "1231"
	name := "gggg"
	c.Name = &name
	zap.L().Debug("replaced zap's global loggers", zap.String("memberID", "1231"), zap.Any("member", c))
}

type a struct {
	ID   string  `json:"ID,omitempty"`
	Name *string `json:"name,omitempty"`
}

func NewLogger(mode string) *zap.Logger {
	stdPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { return lev >= zap.InfoLevel })
	//accessPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { return lev == zap.InfoLevel })
	//systemPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { return lev >= zap.WarnLevel })
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	if mode == "debug" || mode == "" {
		stdPriority = func(lev zapcore.Level) bool {
			return lev >= zap.DebugLevel
		}
	}

	//syncer := zapcore.AddSync(writer)
	//accessCore := zapcore.NewCore(jsonEnc, syncer, accessPriority)
	//systemCore := zapcore.NewCore(jsonEnc, syncer, systemPriority)
	stdCore := zapcore.NewCore(jsonEnc, zapcore.Lock(os.Stdout), stdPriority)

	core := zapcore.NewTee(stdCore)

	//_ = accessCore
	//_ = systemCore
	return zap.New(core).WithOptions(zap.AddCaller())
}
