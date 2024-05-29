package echo

func RegisterRouter() {
	server.e.GET("/", indexHandler)
	server.e.GET("/cards", cardsHandler)
	// 其他路由...
}
