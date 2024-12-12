package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

// 将 map 转换为 Go 结构体的库
type AppConfig struct {
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*MySQLConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"log"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host          string `mapstructure:"host"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	DB            string `mapstructure:"dbname"`
	Port          int    `mapstructure:"port"`
	MaxOpenConns  int    `mapstructure:"max_open_conns"`
	MaxIdleConns  int    `mapstructure:"max_idle_conns"`
	AutoMigration bool   `mapstructure:"auto_migration"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

func Init(filePath string) (err error) {
	// SetConfigFile 显式定义配置文件的路径、名称和扩展名。
	// Viper 将使用它，而不是检查任何配置路径。
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// WatchConfig 开始监视配置文件的更改。
	viper.WatchConfig()
	// OnConfigChange 设置配置文件更改时调用的事件处理程序。
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
}
