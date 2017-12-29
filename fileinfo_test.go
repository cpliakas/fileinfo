package fileinfo_test

import (
	"fmt"
	"testing"

	"github.com/cpliakas/fileinfo"
)

func TestNewMissing(t *testing.T) {
	_, err := fileinfo.New("fixtures/missing.txt")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func ExampleFileinfo_Hash() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	h, _ := i.Hash()
	fmt.Println(h)
	// Output: ace3598e21517c9db3e65621e7f578a4
}

func ExampleFileinfo_Type() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	h, _ := i.Type()
	fmt.Println(h)
	// Output: image/jpeg
}
