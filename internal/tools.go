// +build tools

package internal

import (
	// These tools required by ginkgo
	_ "github.com/go-task/slim-sprig"
	_ "github.com/google/pprof/profile"

	// These tools required to run counterfeiter
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
