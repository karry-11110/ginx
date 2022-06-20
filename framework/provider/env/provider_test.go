package env

import (
    "testing"

    "github.com/karry-11110/ginx/framework"
    "github.com/karry-11110/ginx/framework/contract"
    "github.com/karry-11110/ginx/framework/provider/app"
    tests "github.com/karry-11110/ginx/test"

    . "github.com/smartystreets/goconvey/convey"
)

func TestGinxEnvProvider(t *testing.T) {
    Convey("test ginx env normal case", t, func() {
        basePath := tests.BasePath
        c := framework.NewGinxContainer()
        sp := &app.GinxAppProvider{BaseFolder: basePath}

        err := c.Bind(sp)
        So(err, ShouldBeNil)

        sp2 := &GinxEnvProvider{}
        err = c.Bind(sp2)
        So(err, ShouldBeNil)

        envServ := c.MustMake(contract.EnvKey).(contract.Env)
        So(envServ.AppEnv(), ShouldEqual, "development")
        // So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
        // So(envServ.AppDebug(), ShouldBeTrue)
    })
}
