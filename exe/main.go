package main

import (
	_ "embed"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed index.html
var page []byte

func main() {
	dir, err := os.MkdirTemp("", "katakoryfi_voli")
	if err != nil {
		return
	}
	p := filepath.Join(dir, "index.html")
	if os.WriteFile(p, page, 0o644) != nil {
		return
	}
	exec.Command("rundll32", "url.dll,FileProtocolHandler", p).Start()
}
