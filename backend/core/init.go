package core

import (
	"fmt"
	"os"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// DatabaseCfg 数据库配置
type DatabaseCfg struct {
	Type     string
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Charset  string
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// JWTSetting JWT配置
type JWTSetting struct {
	Secret string
	Expire int // 小时
}

// ServerConfig 服务器配置
type ServerConfig struct {
	RunMode string
	Port    int
}

// 全局配置变量
var (
	DbConfig   DatabaseCfg
	RedisCfg   RedisConfig
	JwtSetting JWTSetting
	ServerCfg  ServerConfig
)

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	// 读取配置
	host, _ := beego.AppConfig.String("db.host")
	port, _ := beego.AppConfig.Int("db.port")
	database, _ := beego.AppConfig.String("db.database")
	username, _ := beego.AppConfig.String("db.username")
	password, _ := beego.AppConfig.String("db.password")
	charset, _ := beego.AppConfig.String("db.charset")

	// 设置配置
	DbConfig = DatabaseCfg{
		Type:     "mysql",
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
		Charset:  charset,
	}

	// 构建数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 创建默认连接
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		return fmt.Errorf("注册数据库失败: %v", err)
	}

	// 设置连接池参数
	if db, err := orm.GetDB(); err == nil {
		// 最大空闲连接数
		db.SetMaxIdleConns(10)
		// 最大打开连接数
		db.SetMaxOpenConns(100)
		// 连接最大存活时间
		db.SetConnMaxLifetime(time.Hour)
	}

	// 自动建表（开发模式）
	runMode, _ := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		orm.RunSyncdb("default", true, true)
	}

	fmt.Println("✅ 数据库连接成功")
	return nil
}

// InitRedis 初始化Redis连接
func InitRedis() error {
	// 读取配置
	host, _ := beego.AppConfig.String("redis.host")
	port, _ := beego.AppConfig.Int("redis.port")
	password, _ := beego.AppConfig.String("redis.password")
	db, _ := beego.AppConfig.Int("redis.db")

	RedisCfg = RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}

	// Redis连接信息
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("✅ Redis连接配置: %s (DB: %d)\n", addr, db)
	
	return nil
}

// InitConfig 初始化所有配置
func InitConfig() error {
	// 读取服务器配置
	runMode, _ := beego.AppConfig.String("runmode")
	port, _ := beego.AppConfig.Int("httpport")

	ServerCfg = ServerConfig{
		RunMode: runMode,
		Port:    port,
	}

	// 读取JWT配置
	jwtSecret, _ := beego.AppConfig.String("jwt.secret")
	jwtExpire, _ := beego.AppConfig.Int("jwt.expire")

	JwtSetting = JWTSetting{
		Secret: jwtSecret,
		Expire: jwtExpire,
	}

	// 确保必要的配置存在
	if jwtSecret == "" {
		JwtSetting.Secret = "caelum-default-secret"
	}
	if jwtExpire == 0 {
		JwtSetting.Expire = 24
	}

	// 创建必要的目录
	if err := os.MkdirAll("logs", 0755); err != nil {
		fmt.Printf("⚠️ 创建日志目录失败: %v\n", err)
	}

	fmt.Println("✅ 配置初始化完成")
	return nil
}