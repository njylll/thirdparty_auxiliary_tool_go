package redis

import (
	"context"
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
	gredis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	DBNUM_5    = 5
	DBNUM_10   = 10
	MAXRETRIES = 3
)

var DB5 *CRedisClient
var DB10 *CRedisClient

type CRedisClient struct {
	Client gredis.UniversalClient
}

func Setup() {
	DB5 = new(CRedisClient)
	DB10 = new(CRedisClient)

	var err error
	logrus.Info("开始初始化redis")
	DB5.Client, err = newRedisClient(DBNUM_5)
	if err != nil {
		logrus.Fatalf("初始化redisDB5失败, err:%v", err)
		return
	}

	DB10.Client, err = newRedisClient(DBNUM_10)
	if err != nil {
		logrus.Fatalf("初始化redisDB10失败, err:%v", err)
		return
	}
	logrus.Info("初始化redis成功")

}

func newRedisClient(dbNum int) (gredis.UniversalClient, error) {
	//创建redis客户端
	client := gredis.NewUniversalClient(&gredis.UniversalOptions{
		Addrs:      setting.RedisSetting.Addrs,
		MaxRetries: MAXRETRIES,
		DB:         dbNum,
		Password:   setting.RedisSetting.Password,
		MasterName: setting.RedisSetting.Master,
	})

	redisAddr := strings.Join(setting.RedisSetting.Addrs, ",")
	//测试链接
	result, err := client.Ping(context.Background()).Result()
	if err != nil || result != "PONG" {
		return nil, fmt.Errorf("init redis failure.[detail]url: %s, dbNum: %d, err:%v", redisAddr, dbNum, err)
	}
	return client, nil
}

// hmSet适配
func (c *CRedisClient) hmSet(ctx context.Context, keyStr []interface{}) *gredis.StatusCmd {
	cmdSlice := []interface{}{"hmset"}
	cmdSlice = append(cmdSlice, keyStr...)
	cmder := gredis.NewStatusCmd(ctx, cmdSlice...)
	_ = c.Client.Process(ctx, cmder)
	return cmder
}

// hGet适配
func (c *CRedisClient) hGet(ctx context.Context, keyStr []interface{}) *gredis.StringCmd {
	cmdSlice := []interface{}{"hget"}
	cmdSlice = append(cmdSlice, keyStr...)
	cmder := gredis.NewStringCmd(ctx, cmdSlice...)
	_ = c.Client.Process(ctx, cmder)
	return cmder
}

// hmGet适配
func (c *CRedisClient) hmGet(ctx context.Context, keyStr []interface{}) *gredis.SliceCmd {
	cmdSlice := []interface{}{"hmget"}
	cmdSlice = append(cmdSlice, keyStr...)
	cmder := gredis.NewSliceCmd(ctx, cmdSlice...)
	_ = c.Client.Process(ctx, cmder)
	return cmder
}

func (c *CRedisClient) HmGetString(ctx context.Context, str []interface{}) ([]string, error) {
	var value []string
	if c.Client == nil {
		err := fmt.Errorf("%s", "in HmGetString redis conn is nil")
		return value, err
	}

	result, err := c.hmGet(ctx, str).Result()
	if err != nil {
		return nil, fmt.Errorf("HmGetString 获取redis失败,err: %v", err)
	}
	for i, v := range result {
		//缺值返回错误
		if v == nil {
			return nil, fmt.Errorf("请求的第%d个值不存在", i)
		}
		value = append(value, v.(string))
	}
	return value, err
}
