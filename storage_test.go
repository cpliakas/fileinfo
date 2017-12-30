package fileinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cpliakas/fileinfo"
)

func TestStorage(t *testing.T) {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		t.Fatal(err)
	}

	f, err := ioutil.TempFile(os.TempDir(), "fileinfo-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	s, err := fileinfo.NewStorage(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	err = s.Save(i)
	if err != nil {
		t.Fatal(err)
	}
}
