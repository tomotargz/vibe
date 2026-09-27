package cli

import (
	"bytes"
	"strings"
	"testing"
)

// executeRoot runs the root command with args and returns its combined output.
func executeRoot(args ...string) (string, error) {
	// SetArgs(nil) makes Cobra fall back to os.Args, so always pass a non-nil slice.
	if args == nil {
		args = []string{}
	}

	var out bytes.Buffer
	cmd := NewRootCmd()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return out.String(), err
}

func TestRootCmd_NoArgsShowsUsage(t *testing.T) {
	out, err := executeRoot()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	if !strings.Contains(out, "vibe") {
		t.Errorf("output does not contain %q:\n%s", "vibe", out)
	}
}

func TestRootCmd_UnknownArgReturnsError(t *testing.T) {
	if _, err := executeRoot("foo"); err == nil {
		t.Fatalf("Execute() returned nil error, want error for unknown arg")
	}
}
