// License: GPLv3 Copyright: 2023, Kovid Goyal, <kovid at kovidgoyal.net>

package transfer

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kovidgoyal/kitty/tools/tty"
)

func signature_timeout(opts *Options) time.Duration {
	if opts == nil || opts.SignatureTimeout <= 0 {
		return 30 * time.Second
	}
	return opts.SignatureTimeout
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
