package config

import (
	"path/filepath"
	"testing"

	"github.com/karry-11110/ginx/framework/contract"
	tests "github.com/karry-11110/ginx/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGinxConfig_GetInt(t *testing.T) {
	container := tests.InitBaseContainer()

	Convey("test ginx env normal case", t, func() {
		appService := container.MustMake(contract.AppKey).(contract.App)
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		folder := filepath.Join(appService.ConfigFolder(), envService.AppEnv())

		serv, err := NewGinxConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*GinxConfig)
		timeout := conf.GetString("database.default.timeout")
		So(timeout, ShouldEqual, "10s")
	})
}
