package initialize

func Run() {
	LoadConfig()
	Mysql()
	Redis()
	// go Cron()
	go SaleWatcher()
	Router()
}
