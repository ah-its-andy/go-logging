package logging

import (
	"fmt"

	"github.com/xcltapestry/xclpkg/clcolor"
)

type ConsoleProvider struct {
}

func (provider *ConsoleProvider) Log(level Level, msg string) {
	fmt.Printf("%s: %s", clcolor.Black(level.String()), msg)
}
