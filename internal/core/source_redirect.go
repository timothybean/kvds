package core

import (
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/logger"
)

// sourceRedirect is a source that redirects to another one.
type sourceRedirect struct{}

func (*sourceRedirect) Log(logger.Level, string, ...interface{}) {
}

// APISourceDescribe implements source.
func (*sourceRedirect) APISourceDescribe() defs.APIPathSourceOrReader {
	return defs.APIPathSourceOrReader{
		Type: "redirect",
		ID:   "",
	}
}
