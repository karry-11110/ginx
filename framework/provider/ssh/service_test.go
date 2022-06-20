package ssh

import (
    "github.com/karry-11110/ginx/framework/provider/config"
    "github.com/karry-11110/ginx/framework/provider/log"
    tests "github.com/karry-11110/ginx/test"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestGinxSSHService_Load(t *testing.T) {
    container := tests.InitBaseContainer()
    container.Bind(&config.GinxConfigProvider{})
    container.Bind(&log.GinxLogServiceProvider{})

    Convey("test get client", t, func() {
        ginxRedis, err := NewGinxSSH(container)
        So(err, ShouldBeNil)
        service, ok := ginxRedis.(*GinxSSH)
        So(ok, ShouldBeTrue)
        client, err := service.GetClient(WithConfigPath("ssh.web-01"))
        So(err, ShouldBeNil)
        So(client, ShouldNotBeNil)
        session, err := client.NewSession()
        So(err, ShouldBeNil)
        out, err := session.Output("pwd")
        So(err, ShouldBeNil)
        So(out, ShouldNotBeNil)
        session.Close()
    })
}
