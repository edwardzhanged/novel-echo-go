package model

import (
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	gorm.Model
	Username     string        `gorm:"type:varchar(50);not null" json:"username" comment:"用户名"`
	Password     string        `gorm:"type:varchar(50);not null" json:"password" comment:"密码"`
	Name         string        `gorm:"type:varchar(100)" json:"name" comment:"真实姓名"`
	Sex          uint8         `gorm:"type:tinyint unsigned" json:"sex" comment:"性别"`
	Birth        time.Time     `gorm:"type:datetime" json:"birth" comment:"出身日期"`
	Email        string        `gorm:"type:varchar(100)" json:"email" comment:"邮箱"`
	Mobile       string        `gorm:"type:varchar(100)" json:"mobile" comment:"手机号"`
	Status       uint8         `gorm:"type:tinyint unsigned;not null;default:1" json:"status" comment:"状态"`
	SysUserRoles []SysUserRole `gorm:"foreignKey:UserID"`
}

type SysRole struct {
	gorm.Model
	RoleName     string        `gorm:"type:varchar(100);not null" json:"role_name" comment:"角色名称"`
	RoleSign     string        `gorm:"type:varchar(100)" json:"role_sign" comment:"角色标识"`
	Remark       string        `gorm:"type:varchar(100)" json:"remark" comment:"备注"`
	SysUserRoles []SysUserRole `gorm:"foreignKey:RoleID"`
	SysRoleMenus []SysRoleMenu `gorm:"foreignKey:RoleID"`
}

type SysUserRole struct {
	gorm.Model
	UserID uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"用户ID"`
	RoleID uint64 `gorm:"type:bigint unsigned;not null" json:"role_id" comment:"角色ID"`
}

type SysMenu struct {
	gorm.Model
	ParentID     uint64        `gorm:"type:bigint unsigned;not null;default:0" json:"parent_id" comment:"父菜单ID 一级菜单为0"`
	Name         string        `gorm:"type:varchar(50);not null" json:"name" comment:"菜单名称"`
	URL          string        `gorm:"type:varchar(200)" json:"url" comment:"菜单URL"`
	Type         uint8         `gorm:"type:tinyint unsigned;not null" json:"type" comment:"类型 0-目录 1-菜单"`
	Icon         string        `gorm:"type:varchar(50)" json:"icon" comment:"菜单图标"`
	Sort         uint8         `gorm:"type:tinyint unsigned" json:"sort" comment:"排序"`
	SysRoleMenus []SysRoleMenu `gorm:"foreignKey:MenuID"`
}

type SysRoleMenu struct {
	gorm.Model
	RoleID uint64 `gorm:"type:bigint unsigned;not null" json:"role_id" comment:"角色ID"`
	MenuID uint64 `gorm:"type:bigint unsigned;not null" json:"menu_id" comment:"菜单ID"`
}

type SysLog struct {
	gorm.Model
	UserID    uint64 `gorm:"type:bigint unsigned" json:"user_id" comment:"用户ID"`
	Username  string `gorm:"type:varchar(50)" json:"username" comment:"用户名"`
	Operation string `gorm:"type:varchar(50)" json:"operation" comment:"用户操作"`
	Time      uint32 `gorm:"type:int unsigned" json:"time" comment:"响应时间"`
	Method    string `gorm:"type:varchar(200)" json:"method" comment:"请求方法"`
	Params    string `gorm:"type:varchar(5000)" json:"params" comment:"请求参数"`
	IP        string `gorm:"type:varchar(64)" json:"ip" comment:"IP地址"`
}
