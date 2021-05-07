package testutil

import (
	"os"

	"github.com/ganboonhong/gotp/cmd/app"
	pkgConfig "github.com/ganboonhong/gotp/pkg/config"
)

func SetupDB(suitename string) {
	config := pkgConfig.NewTestConfig(suitename)

	if err := app.InitApp(config); err != nil {
		panic(err)
	}
}

func TearDownDB(suitename string) {
	configDir := pkgConfig.NewTestConfig(suitename).Dir()
	os.RemoveAll(configDir)
}
