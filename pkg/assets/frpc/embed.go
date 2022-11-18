package frpc

import (
	"Erfrp/pkg/assets"
	"embed"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
