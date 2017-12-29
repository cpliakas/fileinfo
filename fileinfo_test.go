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

func ExampleFileinfo_Name() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	n := i.Name()
	fmt.Println(n)
	// Output: fixtures/image1.jpg
}

func ExampleFileinfo_Basename() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	n := i.Basename()
	fmt.Println(n)
	// Output: image1.jpg
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

func ExampleFileinfo_FirstBytes() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	bytes, _ := i.FirstBytes()
	fmt.Println(bytes)
	// Output: /9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjo=
}

func ExampleFileinfo_LastBytes() {
	i, err := fileinfo.New("fixtures/image1.jpg")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	bytes, _ := i.LastBytes()
	fmt.Println(bytes)
	// Output: q7/Hc/zf1oOcle9UZBk402VnwD6UK9l/ePvQqUQ//9k=
}
