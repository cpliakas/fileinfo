package fileinfo

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
)

// New returns a new Fileinfo populated with the named file opened for reading.
func New(fname string) (i *Fileinfo, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}

	i = &Fileinfo{f}
	return
}

// Fileinfo extracts information about the passed file.
type Fileinfo struct {
	File *os.File
}

// Close closes the file, rendering it unusable for I/O.
func (i *Fileinfo) Close() error {
	return i.File.Close()
}

// Hash returns the file's MD5 sum.
func (i *Fileinfo) Hash() (hash string, err error) {
	hasher := md5.New()

	b, err := ioutil.ReadAll(i.File)
	if err != nil {
		return
	}

	hasher.Write(b)
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}
