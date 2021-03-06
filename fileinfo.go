package fileinfo

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"

	"gopkg.in/h2non/filetype.v1"
)

// New returns a new Fileinfo populated with the named file opened for reading.
func New(fname string) (i *Fileinfo, err error) {
	stat, err := os.Stat(fname)
	if err != nil {
		return
	}

	f, err := os.Open(fname)
	if err != nil {
		return
	}

	i = &Fileinfo{
		bufferSize: bufferSize(stat.Size()),
		file:       f,
		name:       fname,
		offset:     offset(stat.Size()),
		stat:       stat,
	}
	return
}

// Fileinfo extracts information about the passed file.
type Fileinfo struct {
	bufferSize int64
	file       *os.File
	name       string
	offset     int64
	stat       os.FileInfo
}

// Close closes the file, rendering it unusable for I/O.
func (i *Fileinfo) Close() error {
	return i.file.Close()
}

// Name returns the file name.
func (i *Fileinfo) Name() string {
	return i.name
}

// Basename returns the basename of the file.
func (i *Fileinfo) Basename() string {
	return i.stat.Name()
}

// Size returns the size of the file in bytes.
func (i *Fileinfo) Size() int64 {
	return i.stat.Size()
}

// Type returns the mime type.
func (i *Fileinfo) Type() (typ string, err error) {

	// Get the MIME type from the file's magic bits.
	kind, unknown := filetype.MatchReader(i.file)
	if unknown != nil {
		err = fmt.Errorf("unknown file type: %s", unknown)
		return
	}

	// Fall back to the file's extension.
	typ = kind.MIME.Value
	if typ == "" {
		ext := filepath.Ext(i.name)
		typ = mime.TypeByExtension(ext)
	}

	return
}

// Hash returns the file's MD5 sum.
func (i *Fileinfo) Hash() (hash string, err error) {
	h := md5.New()
	_, err = io.Copy(h, i.file)
	if err == nil {
		hash = hex.EncodeToString(h.Sum(nil))
	}
	return
}

// FirstBytes returns the first 32 bytes of a file, base64 encoded.
func (i *Fileinfo) FirstBytes() (bytes string, err error) {
	buf := make([]byte, i.bufferSize)
	_, err = i.file.ReadAt(buf, 0)
	if err == nil {
		bytes = encode(buf)
	}
	return
}

// LastBytes returns the last 32 bytes of a file, base64 encoded.
func (i *Fileinfo) LastBytes() (bytes string, err error) {
	buf := make([]byte, i.bufferSize)
	_, err = i.file.ReadAt(buf, i.offset)
	if err == nil {
		bytes = encode(buf)
	}
	return
}

// bufferSize calculated how many bytes are read when extracting the first and
// last bytes of the file.
func bufferSize(size int64) (bufsize int64) {
	bufsize = size
	if bufsize > 32 {
		bufsize = 32
	}
	return
}

// offset returns the file position to start at when reading the last bytes of
// the file.
func offset(size int64) (off int64) {
	off = size - 32
	if off < 0 {
		off = 0
	}
	return
}

// encode is a utility function that base64 encodes the slice of bytes.
func encode(buf []byte) string {
	return base64.StdEncoding.EncodeToString(buf)
}
