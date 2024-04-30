package record

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/logger"
)

func commonPath(v string) string {
	common := ""
	remaining := v

	for {
		i := strings.IndexAny(remaining, "\\/")
		if i < 0 {
			break
		}

		var part string
		part, remaining = remaining[:i+1], remaining[i+1:]

		if strings.Contains(part, "%") {
			break
		}

		common += part
	}

	if len(common) > 0 {
		common = common[:len(common)-1]
	}

	return common
}

// CleanerEntry is a cleaner entry.
type CleanerEntry struct {
	PathFormat  string
	Format      conf.RecordFormat
	DeleteAfter time.Duration
}

// Cleaner removes expired recording segments from disk.
type Cleaner struct {
	Entries []CleanerEntry
	Parent  logger.Writer

	ctx       context.Context
	ctxCancel func()

	done chan struct{}
}

// Initialize initializes a Cleaner.
func (c *Cleaner) Initialize() {
	c.ctx, c.ctxCancel = context.WithCancel(context.Background())
	c.done = make(chan struct{})

	go c.run()
}

// Close closes the Cleaner.
func (c *Cleaner) Close() {
	c.ctxCancel()
	<-c.done
}

// Log is the main logging function.
func (c *Cleaner) Log(level logger.Level, format string, args ...interface{}) {
	c.Parent.Log(level, "[record cleaner]"+format, args...)
}

func (c *Cleaner) run() {
	defer close(c.done)

	interval := 30 * 60 * time.Second
	for _, e := range c.Entries {
		if interval > (e.DeleteAfter / 2) {
			interval = e.DeleteAfter / 2
		}
	}

	c.doRun() //nolint:errcheck

	for {
		select {
		case <-time.After(interval):
			c.doRun()

		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Cleaner) doRun() {
	for _, e := range c.Entries {
		c.doRunEntry(&e) //nolint:errcheck
	}
}

func (c *Cleaner) doRunEntry(e *CleanerEntry) error {
	pathFormat := e.PathFormat

	switch e.Format {
	case conf.RecordFormatMPEGTS:
		pathFormat += ".ts"

	default:
		pathFormat += ".mp4"
	}

	// we have to convert to absolute paths
	// otherwise, commonPath and fpath inside Walk() won't have common elements
	pathFormat, _ = filepath.Abs(pathFormat)

	commonPath := commonPath(pathFormat)
	now := timeNow()

	filepath.Walk(commonPath, func(fpath string, info fs.FileInfo, err error) error { //nolint:errcheck
		if err != nil {
			return err
		}

		if !info.IsDir() {
			var pa path
			ok := pa.decode(pathFormat, fpath)
			if ok {
				if now.Sub(time.Time(pa)) > e.DeleteAfter {
					c.Log(logger.Debug, "removing %s", fpath)
					os.Remove(fpath)
				}
			}
		}

		return nil
	})

	filepath.Walk(commonPath, func(fpath string, info fs.FileInfo, err error) error { //nolint:errcheck
		if err != nil {
			return err
		}

		if info.IsDir() {
			os.Remove(fpath)
		}

		return nil
	})

	return nil
}
