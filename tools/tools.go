//go:build tools
// +build tools

package tools

import (
	_ "cirello.io/openapigen"
	_ "github.com/go-swagger/go-swagger"
	_ "github.com/grpc-ecosystem/grpc-health-probe"
	_ "github.com/gunk/gunk"
	_ "github.com/gunk/opt/openapiv2"
	_ "github.com/kenshaw/git-buildnumber"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/gofumpt"
	_ "sigs.k8s.io/kustomize/kustomize/v4"
)
