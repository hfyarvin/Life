# GIN
### router
- router.GET("/someGet", getting)
- router.POST("/somePost", posting)
- router.PUT("/somePut", putting)
- router.DELETE("/someDelete", deleting)
- router.PATCH("/somePatch", patching)
- router.HEAD("/someHead", head)
- router.OPTIONS("/someOptions", options)

### 获取参数
- 路由参数--Context.Param()
- url参数--DefaultQuery 或 Query
- 表单参数--PostForm 或 DefaultForm  c.Request.PostForm