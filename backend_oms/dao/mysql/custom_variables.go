package mysql

import "github.com/oms/models"

// 定义初始化数据
var (
	InitContentTypesData = []models.ModelContentType{
		{AppLabel: "admin", Model: "logentry"},
		{AppLabel: "auth", Model: "group"},
		{AppLabel: "auth", Model: "permission"},
		{AppLabel: "auth", Model: "user"},
		{AppLabel: "contenttypes", Model: "contenttype"},
		{AppLabel: "sessions", Model: "session"},
		{AppLabel: "omsenvconfigs", Model: "env"},
	}

	InitLogEntry = []models.AuthPermissionRequest{
		{CodeName: "add_logentry", Name: "增加日志条目"},
		{CodeName: "change_logentry", Name: "修改日志条目"},
		{CodeName: "delete_logentry", Name: "删除日志条目"},
		{CodeName: "view_logentry", Name: "查看日志条目"},
	}

	InitGroups = []models.AuthPermissionRequest{
		{CodeName: "add_group", Name: "增加用户组"},
		{CodeName: "change_group", Name: "修改用户组"},
		{CodeName: "delete_group", Name: "删除用户组"},
		{CodeName: "view_group", Name: "查看用户组"},
	}

	InitPermissions = []models.AuthPermissionRequest{
		{CodeName: "add_permission", Name: "增加权限"},
		{CodeName: "change_permission", Name: "修改权限"},
		{CodeName: "delete_permission", Name: "删除权限"},
		{CodeName: "view_permission", Name: "查看权限"},
	}

	InitUser = []models.AuthPermissionRequest{
		{CodeName: "add_user", Name: "增加用户"},
		{CodeName: "change_user", Name: "修改用用户"},
		{CodeName: "delete_user", Name: "删除用户"},
		{CodeName: "view_user", Name: "查看用户"},
	}

	InitContentType = []models.AuthPermissionRequest{
		{CodeName: "add_contenttype", Name: "增加内容类型"},
		{CodeName: "change_contenttype", Name: "修改内容类型"},
		{CodeName: "delete_contenttype", Name: "删除内容类型"},
		{CodeName: "view_contenttype", Name: "查看内容类型"},
	}

	InitSession = []models.AuthPermissionRequest{
		{CodeName: "add_session", Name: "增加session"},
		{CodeName: "change_session", Name: "修改session"},
		{CodeName: "delete_session", Name: "删除session"},
		{CodeName: "view_session", Name: "查看session"},
	}
	InitEnv = []models.AuthPermissionRequest{
		{CodeName: "add_env", Name: "增加环境"},
		{CodeName: "change_env", Name: "修改环境"},
		{CodeName: "delete_env", Name: "删除环境"},
		{CodeName: "view_env", Name: "查看环境"},
	}
)
