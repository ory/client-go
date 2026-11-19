package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)
// Regression test: addFile must close the source file only after copying its
// contents into the multipart writer.
func TestAddFileCopiesContentBeforeClose(t *testing.T) {
	want := "ory-file-content"
	dir := t.TempDir()
	path := filepath.Join(dir, "upload.txt")
	if err := os.WriteFile(path, []byte(want), 0o600); err != nil {
		t.Fatal(err)
	}

	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	if err := addFile(w, "file", path); err != nil {
		t.Fatalf("addFile returned error: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, "http://example.invalid", &body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	if err := req.ParseMultipartForm(1 << 20); err != nil {
		t.Fatal(err)
	}
	file, header, err := req.FormFile("file")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	if header.Filename != "upload.txt" {
		t.Fatalf("unexpected filename: %q", header.Filename)
	}
	got, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("copied content mismatch: got %q, want %q", got, want)
	}
}
