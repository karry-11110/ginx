package redis

import (
	"context"
	"github.com/karry-11110/ginx/framework/provider/config"
	"github.com/karry-11110/ginx/framework/provider/log"
	tests "github.com/karry-11110/ginx/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGinxService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.GinxConfigProvider{})
	container.Bind(&log.GinxLogServiceProvider{})

	Convey("test get client", t, func() {
		ginxRedis, err := NewGinxRedis(container)
		So(err, ShouldBeNil)
		service, ok := ginxRedis.(*GinxRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
