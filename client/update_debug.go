// +build !release,!autoupdate

package client

import (
	"github.com/zk-123/ngrok/client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
