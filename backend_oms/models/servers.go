package models

type Capacity struct {
	ID    int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name  string `gorm:"column:name;type:varchar(100);unique;not null"`
	Value int64  `gorm:"column:value;type:bigint(20);not null"`
	Unit  string `gorm:"column:unit;type:varchar(100);not null"`
}

type Networks struct {
	ID      int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name    string `gorm:"column:name;type:varchar(100);unique;not null"`
	Address string `gorm:"column:address;type:varchar(100);unique;not null"`
	IsUsed  bool   `gorm:"column:is_used;type:tinyint(1);not null;default:0"`
	CIDR    string `gorm:"column:cidr;type:varchar(100);not null"`
	Desc    string `gorm:"column:desc;type:text"`
}

// Tags - 标签
type Tags struct {
	ID   int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name string `gorm:"column:name;type:varchar(100);unique;not null"`
}

// ServersOS - 服务器操作系统
type ServersOS struct {
	ID   int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name string `gorm:"column:name;type:varchar(100);unique;not null"`
}

// Servers - 服务器
type Servers struct {
	ID       int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Name     string `gorm:"column:name;type:varchar(100);unique;not null"`
	Hostname string `gorm:"column:name;type:varchar(100);unique;not null"`
	Status   bool   `gorm:"column:status;type:tinyint(1);not null;default:0"`
	UserID   int64  `gorm:"column:user_id;type:bigint(20);not null"`
	GroupID  int64  `gorm:"column:group_id;type:bigint(20);not null"`
	OSId     int64  `gorm:"column:os_id;type:bigint(20);not null"`
	Desc     string `gorm:"column:desc;type:text"`
}

// ServerNetworks - 服务器网络
type ServersNetworks struct {
	ID        int64 `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	ServerID  int64 `gorm:"column:server_id;type:bigint(20);not null"`
	NetworkID int64 `gorm:"column:network_id;type:bigint(20);not null"`
}

// ServersTags - 服务器标签
type ServersTags struct {
	ID       int64 `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	ServerID int64 `gorm:"column:server_id;type:bigint(20);not null"`
	TagID    int64 `gorm:"column:tag_id;type:bigint(20);not null"`
}

type ServersCapacity struct {
	ID         int64 `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	ServerID   int64 `gorm:"column:server_id;type:bigint(20);not null"`
	CapacityID int64 `gorm:"column:capacity_id;type:bigint(20);not null"`
}

type Authenticates struct {
	ID       int64  `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	Username string `gorm:"column:username;type:varchar(100);unique;not null"`
	Password string `gorm:"column:password;type:varchar(100);not null"`
}

// AuthenticatesTags - 认证标签
type AuthenticatesTags struct {
	ID             int64 `gorm:"column:id;primaryKey;type:bigint(20);not null"`
	AuthenticateID int64 `gorm:"column:authenticate_id;type:bigint(20);not null"`
	TagID          int64 `gorm:"column:tag_id;type:bigint(20);not null"`
}

type RespServers struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Hostname string `db:"hostname" json:"hostname"`
	Status   bool   `db:"status" json:"status"`
	UserID   int64  `db:"user_id" json:"user_id"`
	GroupID  int64  `db:"group_id" json:"group_id"`
	OSId     int64  `db:"os_id" json:"os_id"`
	Desc     string `db:"desc" json:"desc"`
}

type RespServersList struct {
	Total   int64         `json:"total"`
	Servers []RespServers `json:"servers"`
}
