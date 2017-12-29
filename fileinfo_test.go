package fileinfo_test

import (
	"testing"

	"github.com/cpliakas/fileinfo"
)

func TestFileinfoHash(t *testing.T) {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		t.Fatal(err)
	}

	got, err := i.Hash()
	if err != nil {
		t.Fatal(err)
	}

	ex := "ace3598e21517c9db3e65621e7f578a4"
	if got != ex {
		t.Fatalf("expected %s, got %s", ex, got)
	}
}
