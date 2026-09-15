// License: GPLv3 Copyright: 2023, Kovid Goyal, <kovid at kovidgoyal.net>

package transfer

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kovidgoyal/kitty/tools/tty"
)

const signatureProgressLogInterval = 2 * time.Second

func signature_timeout(opts *Options) time.Duration {
	if opts == nil || opts.SignatureTimeout <= 0 {
		return 30 * time.Second
	}
	return opts.SignatureTimeout
}

func should_log_signature_progress(last_logged_at, now time.Time, is_final bool) bool {
	return is_final || last_logged_at.IsZero() || now.Sub(last_logged_at) >= signatureProgressLogInterval
}

func transfer_debug_enabled() bool {
	val := strings.TrimSpace(strings.ToLower(os.Getenv("KITTY_TRANSFER_DEBUG")))
	return val != "" && val != "0" && val != "false" && val != "no" && val != "off"
}

func transfer_debugf(format string, args ...any) {
	if transfer_debug_enabled() {
		tty.DebugPrintln(fmt.Sprintf("[transfer-debug] "+format, args...))
	}
}
