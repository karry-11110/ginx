package id

import (
	tests "github.com/karry-11110/ginx/test"
	"testing"

	"github.com/karry-11110/ginx/framework/contract"
	"github.com/karry-11110/ginx/framework/provider/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test ginx console log normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&config.GinxConfigProvider{})

		err := c.Bind(&GinxIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		So(xid, ShouldNotBeEmpty)
	})
}
