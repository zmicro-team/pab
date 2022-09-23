package plog

import (
	"io"
	"log"
)

type Logger interface {
	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...any)
	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...any)
	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...any)
	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...any)
	// DPanicf uses fmt.Sprintf to log a templated message. In development, the
	// logger then panics. (See DPanicLevel for details.)
	DPanicf(template string, args ...any)
	// Panicf uses fmt.Sprintf to log a templated message, then panics.
	Panicf(template string, args ...any)
	// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
	Fatalf(template string, args ...any)
}

type StdLogger struct {
	l *log.Logger
}

var _ Logger = (*StdLogger)(nil)

func NewStdLogger(out io.Writer) *StdLogger {
	return &StdLogger{
		log.New(out, "openpab: ", log.LstdFlags),
	}
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *StdLogger) Debugf(template string, args ...any) {
	s.l.Printf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *StdLogger) Infof(template string, args ...any) {
	s.l.Printf(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *StdLogger) Warnf(template string, args ...any) {
	s.l.Printf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *StdLogger) Errorf(template string, args ...any) {
	s.l.Printf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *StdLogger) DPanicf(template string, args ...any) {
	s.l.Printf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *StdLogger) Panicf(template string, args ...any) {
	s.l.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *StdLogger) Fatalf(template string, args ...any) {
	s.l.Fatalf(template, args...)
}
