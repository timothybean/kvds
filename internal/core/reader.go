package core

import (
	"github.com/noranetworks/kvds/internal/asyncwriter"
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/stream"
)

// reader is an entity that can read a stream.
type reader interface {
	close()
	apiReaderDescribe() defs.APIPathSourceOrReader
}

func readerMediaInfo(r *asyncwriter.Writer, stream *stream.Stream) string {
	return mediaInfo(stream.MediasForReader(r))
}
