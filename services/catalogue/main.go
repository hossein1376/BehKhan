package main

// @title           Catalogue
// @version         0.1.0
// @description     BehKhan's catalogue microservice.
// @contact.name    Hossein Yazdani
// @contact.url     https://GodlyNice.ir
// @license.name    MIT license
// @license.url     https://opensource.org/license/mit/
// @host            localhost:8000
// @BasePath        /api/v1/catalogue
func main() {
	go doGrpc()
	doHttp()
}
