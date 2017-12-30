package fileinfo

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

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

	i = &Fileinfo{file: f, name: fname, stat: stat}
	return
}

// Fileinfo extracts information about the passed file.
type Fileinfo struct {
	file *os.File
	name string
	stat os.FileInfo
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

// Type returns the file type as detected from it's magic bits.
func (i *Fileinfo) Type() (typ string, err error) {
	b := make([]byte, 261)
	_, err = i.file.Read(b)
	if err != nil {
		return
	}

	kind, unknown := filetype.Match(b)
	if unknown != nil {
		err = fmt.Errorf("unknown file type: %s", unknown)
		return
	}

	typ = kind.MIME.Value
	return
}

// Hash returns the file's MD5 sum.
func (i *Fileinfo) Hash() (hash string, err error) {
	hasher := md5.New()

	b, err := ioutil.ReadAll(i.file)
	if err != nil {
		return
	}

	hasher.Write(b)
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}

// FirstBytes returns the first 32 bytes of a file, base64 encoded.
func (i *Fileinfo) FirstBytes() (bytes string, err error) {
	buf := make([]byte, 32)
	_, err = i.file.ReadAt(buf, 0)
	if err == nil {
		bytes = encode(buf)
	}
	return
}

// LastBytes returns the last 32 bytes of a file, base64 encoded.
func (i *Fileinfo) LastBytes() (bytes string, err error) {
	buf := make([]byte, 32)
	_, err = i.file.ReadAt(buf, i.Size()-32)
	if err == nil {
		bytes = encode(buf)
	}
	return
}

// encode is a utility function that base64 encodes the slice of bytes.
func encode(buf []byte) string {
	return base64.StdEncoding.EncodeToString(buf)
}
