package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"
)

// runMain is set when the test binary runs again as demou itself.
const runMain = "DEMOU_TEST_RUN_MAIN"

func TestMain(m *testing.M) {
	if os.Getenv(runMain) == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

// demou runs main with args in a process of its own, as the shell would.
func demou(t *testing.T, args ...string) (stdout, stderr string, code int) {
	t.Helper()
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), runMain+"=1")
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		code = exit.ExitCode()
	} else if err != nil {
		t.Fatal(err)
	}
	return out.String(), errOut.String(), code
}

// TestVersion: demou version prints the version, -v and --version too (tdp E1).
func TestVersion(t *testing.T) {
	for _, arg := range []string{"version", "-v", "--version"} {
		out, errOut, code := demou(t, arg)
		if out != "dev\n" || errOut != "" || code != 0 {
			t.Errorf("demou %s: stdout %q, stderr %q, exit %d; want stdout \"dev\\n\", exit 0", arg, out, errOut, code)
		}
	}
}

// TestHelp: demou help prints the command line's usage, -h and --help too
// (tdp E1).
func TestHelp(t *testing.T) {
	for _, arg := range []string{"help", "-h", "--help"} {
		out, errOut, code := demou(t, arg)
		if out != usage || errOut != "" || code != 0 {
			t.Errorf("demou %s: stdout %q, stderr %q, exit %d; want the usage, exit 0", arg, out, errOut, code)
		}
	}
}

// TestUnknownCommand: anything else says so with the usage on stderr, exit 2.
func TestUnknownCommand(t *testing.T) {
	out, errOut, code := demou(t, "shelf")
	want := "demou: unknown command \"shelf\"\n\n" + usage
	if out != "" || errOut != want || code != 2 {
		t.Errorf("demou shelf: stdout %q, stderr %q, exit %d; want stderr %q, exit 2", out, errOut, code, want)
	}
}
