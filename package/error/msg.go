package error

var MsgFlags = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "系统错误",
	NOPERMISSION:   "没有访问权限",
	NEED_LOGIN:     "需要登录",
	INVALID_PARAMS: "请求参数错误",
	NOROUTE:        "找不到该路由",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_EXIST_USER:     "用户已存在",
	ERROR_NOT_EXIST_USER: "用户不存在",
	ERROR_DISABLE_USER:   "用户已被禁止，请联系管理员",
	ERROR_PASSWORD_USER:  "密码加密失败",

	ERROR_EXIST_ROLE:     "角色已存在",
	ERROR_NOT_EXIST_ROLE: "角色不存在",

	ERROR_EXIST_MENU:     "菜单已存在",
	ERROR_NOT_EXIST_MENU: "菜单不存在",

	ERROR_EXIST_CATEGORY:     "分类已存在",
	ERROR_NOT_EXIST_CATEGORY: "分类不存在",

	ERROR_EXIST_PRODUCT:     "商品已存在",
	ERROR_NOT_EXIST_PRODUCT: "商品不存在",

	ERROR_SQL_INSERT_FAIL: "数据插入失败",
	ERROR_SQL_DELETE_FAIL: "数据删除失败",
	ERROR_SQL_UPDATE_FAIL: "数据修改失败",

	ERROR_UPLOAD_IMAGE_FAIL:       "上传图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "检查图片失败",
	ERROR_UPLOAD_IMAGE_EXT_ERR:    "图片格式错误",
	ERROR_UPLOAD_IMAGE_SIZE_ERR:   "图片大小错误",
	ERROR_NOT_EXIST_IMAGE:         "图片不存在",
	ERROR_DEL_IMAGE_FAIL:          "删除图片失败",
	ERROR_MOVE_IMAGE_FAIL:         "删除图片失败",
}

//获取信息码
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
