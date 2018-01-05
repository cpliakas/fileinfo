package fileinfo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/cpliakas/fileinfo"
)

func newFileinfo(t *testing.T, fname string) (i *fileinfo.Fileinfo) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if strings.HasSuffix(dir, "/fixtures") {
		os.Chdir("..")
	}

	i, err = fileinfo.New("./fixtures/" + fname)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestNewMissing(t *testing.T) {
	_, err := fileinfo.New("./fixtures/missing.txt")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestFileinfo_Name(t *testing.T) {
	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg"},
		{name: "text1.txt"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		ex := "./fixtures/" + test.name
		got := i.Name()
		if got != ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}

func TestFileinfo_Basename(t *testing.T) {
	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg"},
		{name: "text1.txt"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		ex := test.name
		got := i.Basename()
		if got != ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}

func TestFileinfo_Hash(t *testing.T) {
	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg", ex: "ace3598e21517c9db3e65621e7f578a4"},
		{name: "image2.jpg", ex: "1eb36d12024520174523a3b93dc86462"},
		{name: "image3.jpg", ex: "1eb36d12024520174523a3b93dc86462"},
		{name: "text1.txt", ex: "202cb962ac59075b964b07152d234b70"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		got, err := i.Hash()
		if err != nil {
			t.Error("error calculating hash", err)
		}
		if got != test.ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}

func TestFileinfo_FirstBytes(t *testing.T) {

	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg", ex: "/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjo="},
		{name: "image2.jpg", ex: "/9j/4AAQSkZJRgABAQEBLAEsAAD/7QCyUGhvdG9zaG8="},
		{name: "image3.jpg", ex: "/9j/4AAQSkZJRgABAQEBLAEsAAD/7QCyUGhvdG9zaG8="},
		{name: "text1.txt", ex: "MTIz"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		got, err := i.FirstBytes()
		if err != nil {
			t.Error("error reading file", err)
		}
		if got != test.ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}

func TestFileinfo_LastBytes(t *testing.T) {

	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg", ex: "q7/Hc/zf1oOcle9UZBk402VnwD6UK9l/ePvQqUQ//9k="},
		{name: "image2.jpg", ex: "jzfe2IIXMicACrcY4M13IdVK0clkNhbFkAyTafB//9k="},
		{name: "image3.jpg", ex: "jzfe2IIXMicACrcY4M13IdVK0clkNhbFkAyTafB//9k="},
		{name: "text1.txt", ex: "MTIz"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		got, err := i.LastBytes()
		if err != nil {
			t.Error("error reading file", err)
		}
		if got != test.ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}

func TestFileinfo_Type(t *testing.T) {

	tests := []struct {
		name string
		ex   string
	}{
		{name: "image1.jpg", ex: "image/jpeg"},
		{name: "text1.txt", ex: "text/plain; charset=utf-8"},
	}

	for _, test := range tests {
		i := newFileinfo(t, test.name)
		got, err := i.Type()
		if err != nil {
			t.Error("error reading file", err)
		}
		if got != test.ex {
			t.Error("expected", test.ex, "got", got)
		}
	}
}
