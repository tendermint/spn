package docs

import "embed"

// Docs represents the embedded doc file
//go:embed static
var Docs embed.FS
