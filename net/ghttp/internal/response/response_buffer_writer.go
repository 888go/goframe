// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//

package response

import (
	"bytes"
	"net/http"
)

// BufferWriter is the custom writer for http response with buffer.
type BufferWriter struct {
	*Writer               // The underlying BufferWriter.
	Status  int           // HTTP status.
	buffer  *bytes.Buffer // The output buffer.
}


// ff:
// writer:
func NewBufferWriter(writer http.ResponseWriter) *BufferWriter {
	return &BufferWriter{
		Writer: NewWriter(writer),
		buffer: bytes.NewBuffer(nil),
	}
}

// RawWriter returns the underlying BufferWriter.

// ff:
func (w *BufferWriter) RawWriter() http.ResponseWriter {
	return w.Writer
}

// Write implements the interface function of http.BufferWriter.Write.

// ff:
// data:
func (w *BufferWriter) Write(data []byte) (int, error) {
	return w.buffer.Write(data)
}

// WriteString writes string content to internal buffer.

// ff:
// data:
func (w *BufferWriter) WriteString(data string) (int, error) {
	return w.buffer.WriteString(data)
}

// Buffer returns the buffered content as []byte.

// ff:
func (w *BufferWriter) Buffer() []byte {
	return w.buffer.Bytes()
}

// BufferString returns the buffered content as string.

// ff:
func (w *BufferWriter) BufferString() string {
	return w.buffer.String()
}

// BufferLength returns the length of the buffered content.

// ff:
func (w *BufferWriter) BufferLength() int {
	return w.buffer.Len()
}

// SetBuffer overwrites the buffer with `data`.

// ff:
// data:
func (w *BufferWriter) SetBuffer(data []byte) {
	w.buffer.Reset()
	w.buffer.Write(data)
}

// ClearBuffer clears the response buffer.

// ff:
func (w *BufferWriter) ClearBuffer() {
	w.buffer.Reset()
}

// WriteHeader implements the interface of http.BufferWriter.WriteHeader.

// ff:
// status:
func (w *BufferWriter) WriteHeader(status int) {
	w.Status = status
}

// Flush outputs the buffer to clients and clears the buffer.

// ff:
func (w *BufferWriter) Flush() {
	if w.Writer.IsHijacked() {
		return
	}

	if w.Status != 0 && !w.Writer.IsHeaderWrote() {
		w.Writer.WriteHeader(w.Status)
	}
	// Default status text output.
	if w.Status != http.StatusOK && w.buffer.Len() == 0 {
		w.buffer.WriteString(http.StatusText(w.Status))
	}
	if w.buffer.Len() > 0 {
		_, _ = w.Writer.Write(w.buffer.Bytes())
		w.buffer.Reset()
		if flusher, ok := w.RawWriter().(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
