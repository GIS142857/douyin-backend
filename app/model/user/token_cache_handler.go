package user

import (
	"douyin-backend/app/global/variable"
	"douyin-backend/app/service/users/token_cache_redis"
)

// 本文件专门处理 token 缓存到 redis 的相关逻辑

func (u *UserModel) ValidTokenCacheToRedis(uid int64) {
	tokenCacheRedisFactory := token_cache_redis.CreateUsersTokenCacheFactory(uid)
	if tokenCacheRedisFactory == nil {
		variable.ZapLog.Error("redis 连接失败，请检查配置")
		return
	}
	defer tokenCacheRedisFactory.ReleaseRedisConn()

	sql := `SELECT token, expires_at FROM tb_auth_access_tokens WHERE uid=? AND revoked=0 AND expires_at>UNIX_TIMESTAMP(NOW()) ORDER BY expires_at DESC , updated_at DESC LIMIT ?`
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, uid, maxOnlineUsers).Rows()
	defer func() {
		// 凡是获取原生结果集的查询，记得释放记录集
		_ = rows.Close()
	}()

	var tempToken string
	var expiresAt int64
	if err == nil && rows != nil {
		for i := 1; rows.Next(); i++ {
			err = rows.Scan(&tempToken, &expiresAt)
			if err == nil {
				tokenCacheRedisFactory.SetTokenCache(expiresAt, tempToken)
				// 因为每个用户的token是按照过期时间倒叙排列的，第一个是有效期最长的，将该用户的总键设置一个最大过期时间，到期则自动清理，避免不必要的数据残留
				if i == 1 {
					tokenCacheRedisFactory.SetUserTokenExpire(expiresAt)
				}
			}
		}
	}
	// 缓存结束之后删除超过系统设置最大在线数量的token
	tokenCacheRedisFactory.DelOverMaxOnlineCache()
}
