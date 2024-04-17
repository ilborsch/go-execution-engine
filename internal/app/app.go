package app

import (
	"github.com/ilborsch/go-execution-engine/internal/code/parser"
	"log/slog"
	"time"
)

type Executor struct {
	log          *slog.Logger
	rateLimit    time.Duration
	scriptParser *parser.Parser
	// TODO: continue implementing executor app
}
