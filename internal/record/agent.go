package record

import (
	"time"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/logger"
	"github.com/noranetworks/kvds/internal/stream"
)

// Agent writes recordings to disk.
type Agent struct {
	WriteQueueSize    int
	PathFormat        string
	Format            conf.RecordFormat
	PartDuration      time.Duration
	SegmentDuration   time.Duration
	PathName          string
	Stream            *stream.Stream
	OnSegmentCreate   OnSegmentFunc
	OnSegmentComplete OnSegmentFunc
	Parent            logger.Writer

	restartPause time.Duration

	currentInstance *agentInstance

	terminate chan struct{}
	done      chan struct{}
}

// Initialize initializes Agent.
func (w *Agent) Initialize() {
	if w.OnSegmentCreate == nil {
		w.OnSegmentCreate = func(string) {
		}
	}
	if w.OnSegmentComplete == nil {
		w.OnSegmentComplete = func(string) {
		}
	}
	if w.restartPause == 0 {
		w.restartPause = 2 * time.Second
	}

	w.terminate = make(chan struct{})
	w.done = make(chan struct{})

	w.currentInstance = &agentInstance{
		agent: w,
	}
	w.currentInstance.initialize()

	go w.run()
}

// Log is the main logging function.
func (w *Agent) Log(level logger.Level, format string, args ...interface{}) {
	w.Parent.Log(level, "[record] "+format, args...)
}

// Close closes the agent.
func (w *Agent) Close() {
	w.Log(logger.Info, "recording stopped")
	close(w.terminate)
	<-w.done
}

func (w *Agent) run() {
	defer close(w.done)

	for {
		select {
		case <-w.currentInstance.done:
			w.currentInstance.close()
		case <-w.terminate:
			w.currentInstance.close()
			return
		}

		select {
		case <-time.After(w.restartPause):
		case <-w.terminate:
			return
		}

		w.currentInstance = &agentInstance{
			agent: w,
		}
		w.currentInstance.initialize()
	}
}
