package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/contrib/log/zap/v2"
	"os"
	"strings"

	"sunflower-blog-svc/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/natefinch/lumberjack"

	_ "go.uber.org/automaxprocs"
	uberzap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "sunflower-blog-svc.blog"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	outputLogger, err := newLogger(bc.Log)
	if err != nil {
		panic(err)
	}

	logger := log.With(outputLogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(bc.Server, bc.Data, &bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newLogger(c *conf.Log) (log.Logger, error) {
	zapConf := uberzap.NewProductionConfig()
	zapConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var logger log.Logger

	if c.GetMode() == "file" {
		fmt.Println("Log Mode: file")
		// 创建日志目录
		if err := os.MkdirAll(c.GetPath(), 0o755); err != nil {
			return nil, err
		}

		// 配置日志输出到文件
		zapConf.OutputPaths = []string{c.GetPath() + "/app.log"}
		zapConf.ErrorOutputPaths = []string{c.GetPath() + "/error.log"}

		// 设置日志分割
		// 注意：需要引入 "gopkg.in/natefinch/lumberjack.v2"
		logRotate := &lumberjack.Logger{
			Filename:   c.GetPath() + "/app.log",
			MaxSize:    int(c.MaxSize),    // 每个日志文件最大尺寸，单位 MB
			MaxBackups: int(c.MaxBackups), // 保留的旧文件最大数量
			MaxAge:     int(c.MaxBackups), // 保留的最大天数
			Compress:   c.Compress,        // 是否压缩
		}

		// 创建自定义 core
		var encoder zapcore.Encoder
		if c.GetEncoding() == "console" {
			encoder = zapcore.NewConsoleEncoder(zapConf.EncoderConfig)
		}

		if c.GetEncoding() == "json" {
			encoder = zapcore.NewJSONEncoder(zapConf.EncoderConfig)
		}
		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(logRotate),
			uberzap.NewAtomicLevelAt(getZapLevel(c.GetLevel())),
		)

		logger = zap.NewLogger(uberzap.New(core, uberzap.AddCaller(), uberzap.AddCallerSkip(2)))

		return logger, nil
	}

	// 默认输出到控制台
	zapConf.OutputPaths = []string{"stdout"}
	zapConf.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := zapConf.Build(uberzap.AddCaller(), uberzap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	return zap.NewLogger(zapLogger), nil
}

// getZapLevel 转换日志级别
func getZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
