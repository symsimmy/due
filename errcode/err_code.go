package errcode

const (
	Succeed                   = 0   // 成功
	Config_error              = 1   // 配置错误
	Player_offline            = 3   // 玩家不在线
	No_authority              = 4   // 权限不足
	Request_param_error       = 8   // 请求参数错误
	Invalid_login_message     = 11  // 注册信息格式错误
	Invalid_auth_api          = 12  // 注册接口请求失败
	No_bind                   = 13  // 网关服务器绑定失败
	No_dispatch               = 14  // 游戏服务器绑定失败
	Invalid_geteuserinfo_api  = 15  // 获取用户信息接口失败
	No_deliver                = 16  // 转发到游戏服务器失败
	Invalid_json_message      = 17  // json信息格式非法
	Leaderboard_max           = 18  // 排行榜已经到达底部
	Invalid_pb_message        = 20  // proto消息格式非法
	Multiple_accounts_kickoff = 21  // 账号多地登录踢用户下线
	Game_server_down_kickoff  = 22  // 游戏服务器下线踢用户下线
	Backward_client_version   = 25  // 客户端版本落后
	Advanced_client_version   = 26  // 客户端版本超前
	Ban_accounts_kickoff      = 27  // 账号违禁踢用户下线
	Bad_request               = 400 // 请求失败
	Unauthorized              = 401 // 未验证
	Forbidden                 = 403 // 被禁止
	Not_found                 = 404 // 未找到
	Internal_server_error     = 500 // 服务器内部错误
	Mongo_op_error            = 501 // mongo操作错误
	Redis_op_error            = 502 // redis操作错误
)
