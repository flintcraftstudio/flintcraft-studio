//go:build mage

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/sh"
)

const tailwindVersion = "v3.4.17"

// InstallTailwind downloads the Tailwind standalone CLI for the current platform
func InstallTailwind() error {
	binary := tailwindBinaryPath()
	if _, err := os.Stat(binary); err == nil {
		fmt.Println("Tailwind already installed, skipping.")
		return nil
	}

	url := tailwindDownloadURL()
	fmt.Printf("Downloading Tailwind %s from %s\n", tailwindVersion, url)

	if err := sh.Run("curl", "-sLo", binary, url); err != nil {
		return err
	}
	return sh.Run("chmod", "+x", binary)
}

// BuildCSS compiles Tailwind CSS
func BuildCSS() error {
	return sh.Run(
		tailwindBinaryPath(),
		"-c", "./tailwind/tailwind.config.js",
		"-i", "./tailwind/input.css",
		"-o", "./web/static/css/site.css",
		"--minify",
	)
}

// BuildChiroCSS compiles the Alpine Spine chiropractor demo's stylesheet.
// Each demo vertical carries its own palette/type via a per-demo Tailwind
// config + input, compiled through the same standalone CLI as the marketing
// site. To add a vertical, copy this target and the tailwind/demos/<name>.*
// files, then swap the palette tokens.
func BuildChiroCSS() error {
	return sh.Run(
		tailwindBinaryPath(),
		"-c", "./tailwind/demos/chiropractor.config.js",
		"-i", "./tailwind/demos/chiropractor.css",
		"-o", "./demos/chiropractor/static/css/site.css",
		"--minify",
	)
}

// BuildLawCSS compiles the Granite Peak Trial Lawyers demo's stylesheet.
// Second demo vertical; same per-demo pipeline as BuildChiroCSS — its own
// Tailwind config + input carry the Granite Peak palette/type.
func BuildLawCSS() error {
	return sh.Run(
		tailwindBinaryPath(),
		"-c", "./tailwind/demos/law.config.js",
		"-i", "./tailwind/demos/law.css",
		"-o", "./demos/law/static/css/site.css",
		"--minify",
	)
}

// GenerateTempl runs templ generate
func GenerateTempl() error {
	return sh.Run("templ", "generate")
}

// BuildGo compiles the Go binary
func BuildGo() error {
	if err := GenerateTempl(); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/server", "./cmd/server")
}

// Build runs a full production build
func Build() error {
	if err := BuildCSS(); err != nil {
		return err
	}
	if err := BuildChiroCSS(); err != nil {
		return err
	}
	if err := BuildLawCSS(); err != nil {
		return err
	}
	return BuildGo()
}

// Dev generates templates, builds CSS, then starts the server
func Dev() error {
	if err := GenerateTempl(); err != nil {
		return err
	}
	if err := BuildCSS(); err != nil {
		return err
	}
	if err := BuildChiroCSS(); err != nil {
		return err
	}
	if err := BuildLawCSS(); err != nil {
		return err
	}
	return sh.RunV("go", "run", "./cmd/server")
}

// templCLI is the templ generator matching go.mod. A newer installed CLI
// generates code the pinned runtime can't compile (see CLAUDE.md).
const templCLI = "github.com/a-h/templ/cmd/templ@v0.3.1020"

// Watch is the live-reload dev loop. Open http://localhost:7331 (templ's
// proxy), not :8080. Only one proxy can hold a port, so when another
// project's watch has 7331, set PROXY_PORT (e.g. PROXY_PORT=7332).
// Tailwind rebuilds the site's and every demo's site.css on save; templ
// regenerates on .templ saves. Go edits restart the server; a site.css
// rebuild is a watched file too, so it restarts and reloads the browser
// with the new styles. Text-only .templ edits reload without a restart.
func Watch(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		proxyPort = "7331"
	}

	// One watcher per stylesheet: the marketing site, then each demo.
	// --watch=always keeps Tailwind running when stdin isn't a terminal.
	sheets := [][3]string{
		{"./tailwind/tailwind.config.js", "./tailwind/input.css", "./web/static/css/site.css"},
		{"./tailwind/demos/chiropractor.config.js", "./tailwind/demos/chiropractor.css", "./demos/chiropractor/static/css/site.css"},
		{"./tailwind/demos/law.config.js", "./tailwind/demos/law.css", "./demos/law/static/css/site.css"},
	}
	for _, s := range sheets {
		tw := exec.CommandContext(ctx, tailwindBinaryPath(), "-c", s[0], "-i", s[1], "-o", s[2], "--watch=always")
		tw.Stdout, tw.Stderr = os.Stdout, os.Stderr
		if err := tw.Start(); err != nil {
			return err
		}
		defer func() { _ = tw.Process.Kill() }()
	}

	return sh.RunV("go", "run", templCLI, "generate", "--watch",
		"--proxy=http://localhost:"+port,
		"--proxyport="+proxyPort,
		"--open-browser=false",
		"--cmd=go run ./cmd/server",
		`--watch-pattern=(.+\.go$)|(.+\.templ$)|(.+_templ\.txt$)|(static/css/site\.css$)`,
	)
}

// Run starts the server
func Run() error {
	return sh.Run("./bin/server")
}

func tailwindBinaryPath() string {
	if runtime.GOOS == "windows" {
		return "./tailwind/tailwindcss.exe"
	}
	return "./tailwind/tailwindcss"
}

func tailwindDownloadURL() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	osName := map[string]string{
		"darwin":  "macos",
		"linux":   "linux",
		"windows": "windows",
	}[os]

	archName := map[string]string{
		"amd64": "x64",
		"arm64": "arm64",
	}[arch]

	ext := ""
	if os == "windows" {
		ext = ".exe"
	}

	return fmt.Sprintf(
		"https://github.com/tailwindlabs/tailwindcss/releases/download/%s/tailwindcss-%s-%s%s",
		tailwindVersion, osName, archName, ext,
	)
}
