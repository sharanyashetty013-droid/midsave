package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// When MIDSAVE_RUN_MAIN is set, the test binary acts as the midsave CLI so
// tests can run it as a real process with its own arguments and exit code.
func TestMain(m *testing.M) {
	if os.Getenv("MIDSAVE_RUN_MAIN") == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func runMidsave(t *testing.T, args ...string) (string, error) {
	t.Helper()
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), "MIDSAVE_RUN_MAIN=1")
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func TestVersionFlag(t *testing.T) {
	out, err := runMidsave(t, "--version")
	if err != nil {
		t.Fatalf("midsave --version failed: %v\n%s", err, out)
	}
	want := "midsave " + currentVersion() + "\n"
	if out != want {
		t.Fatalf("midsave --version printed %q, want %q", out, want)
	}
}

func TestCurrentVersionUsesBuildTimeValue(t *testing.T) {
	old := version
	t.Cleanup(func() { version = old })

	version = "0.1.0"
	if got := currentVersion(); got != "0.1.0" {
		t.Fatalf("currentVersion() = %q, want %q", got, "0.1.0")
	}
}

func TestMissingArgsStillShowsUsage(t *testing.T) {
	out, err := runMidsave(t)
	if err == nil {
		t.Fatalf("midsave with no args should fail, got success:\n%s", out)
	}
	if !strings.Contains(out, "usage: midsave") {
		t.Fatalf("expected usage message, got:\n%s", out)
	}
}
