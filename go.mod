module good_gathering

go 1.12

replace github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 => github.com/ugorji/go v1.1.4

require (
	github.com/fsnotify/fsnotify v1.5.4
	github.com/gin-gonic/gin v1.8.1
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c
	github.com/spf13/cast v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/unrolled/secure v1.12.0
	go.uber.org/zap v1.17.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
