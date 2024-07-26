package middleware

//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code int
//		var data interface{}
//
//		code = response.SUCCESS
//		token := c.Query("token")
//		if token == "" {
//			code = response.InvalidParams
//		} else {
//			claims, err := util.ParseToken(token)
//			if err != nil {
//				code = response.ErrorAuthCheckTokenFail
//			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
//				code = response.ErrorAuthCheckTokenTimeout
//			}
//		}
//
//		if code != response.SUCCESS {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"code": code,
//				"msg":  response.GetMsg(code),
//				"data": data,
//			})
//
//			c.Abort()
//			return
//		}
//
//		c.Next()
//	}
//}
