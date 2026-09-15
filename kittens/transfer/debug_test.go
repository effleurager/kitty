// License: GPLv3 Copyright: 2026, Kovid Goyal, <kovid at kovidgoyal.net>

package transfer

import (
	"testing"
	"time"
)

func TestTransferDebugEnabled(t *testing.T) {
	t.Setenv("KITTY_TRANSFER_DEBUG", "")
	if transfer_debug_enabled() {
		t.Fatal("debug should be disabled for empty value")
	}
	for _, val := range []string{"0", "false", "no", "off"} {
		t.Setenv("KITTY_TRANSFER_DEBUG", val)
		if transfer_debug_enabled() {
			t.Fatalf("debug should be disabled for %q", val)
		}
	}
	for _, val := range []string{"1", "true", "yes", "on"} {
		t.Setenv("KITTY_TRANSFER_DEBUG", val)
		if !transfer_debug_enabled() {
			t.Fatalf("debug should be enabled for %q", val)
		}
	}
}

func TestSignatureTimeoutDefault(t *testing.T) {
	if got := signature_timeout(nil); got != 30*time.Second {
		t.Fatalf("expected default timeout for nil opts, got %s", got)
	}
	if got := signature_timeout(&Options{}); got != 30*time.Second {
		t.Fatalf("expected default timeout for zero opts, got %s", got)
	}
	if got := signature_timeout(&Options{SignatureTimeout: 7 * time.Second}); got != 7*time.Second {
		t.Fatalf("expected configured timeout, got %s", got)
	}
}

func TestShouldLogSignatureProgress(t *testing.T) {
	now := time.Now()
	if !should_log_signature_progress(time.Time{}, now, false) {
		t.Fatal("should log when there has been no prior progress log")
	}
	if should_log_signature_progress(now.Add(-time.Second), now, false) {
		t.Fatal("should not log before the progress interval elapses")
	}
	if !should_log_signature_progress(now.Add(-3*time.Second), now, false) {
		t.Fatal("should log after the progress interval elapses")
	}
	if !should_log_signature_progress(now, now, true) {
		t.Fatal("should always log the final progress event")
	}
}
