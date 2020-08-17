
https://github.com/go-sql-driver/mysql
https://github.com/siddontang/go-mysql



参考
https://xuchao918.github.io/2019/06/13/Go%E6%93%8D%E4%BD%9CMySql%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9A%84%E6%96%B9%E5%BC%8F/

---------------------------------------------------------------------------------------------------------------------  

在golang的开发过程中，当我们使用orm的时候，常常需要将数据库表对应到golang的一个struct，这些struct会携带orm对应的tag,就像下面的struct定义一样：
type InsInfo struct {
	Connections  string    `gorm:"column:connections"`
	CPU          int       `gorm:"column:cpu"`
	CreateTime   time.Time `gorm:"column:create_time"`
	Env          int       `gorm:"column:env"`
	ID           int64     `gorm:"column:id;primary_key"`
	IP           string    `gorm:"column:ip"`
	Organization string    `gorm:"column:organization"`
	Pass         string    `gorm:"column:pass"`
	Port         string    `gorm:"column:port"`
	RegionId     string    `gorm:"column:regionid"`
	ServerIP     string    `gorm:"column:server_ip"`
	Status       int       `gorm:"column:status"`
	Type         string    `gorm:"column:type"`
	UUID         string    `gorm:"column:uuid"`
}


帮助生成上面这种全是重复机械工作的字段映射配置文件的工具
https://github.com/hantmac/fuckdb/blob/master/README_zh.md


---------------------------------------------------------------------------------------------------------------------  
