package logger

// Config Logger Config
type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

type Logger interface {
	InitLogger()
}

// Application logger
type appLogger struct {
	level    string
	devMode  bool
	encoding string
}

// NewAppLogger App Logger constructor
func NewAppLogger(cfg *Config) *appLogger {
	return &appLogger{level: cfg.LogLevel, devMode: cfg.DevMode, encoding: cfg.Encoder}
}

func (l *appLogger) InitLogger() {
	
}
