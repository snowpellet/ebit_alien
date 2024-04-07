package resourse

import (
	_ "embed"
)

//go:embed config.json
var Config []byte
