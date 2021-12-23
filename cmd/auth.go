package main

import (
	"flag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
	"strconv"
)

func main() {
	port := flag.Int("p", 9011, "what's your port?")
	flag.Parse()
	app := iris.New()
	app.PartyFunc("/", func(child iris.Party) {
		child.Post("test", Test)
	})
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)
	_ = app.Run(iris.Addr(":" + strconv.Itoa(*port)))
}

func Test(ctx iris.Context){
	_, _ = ctx.Text("ok")
	return
}

func Test1() string{
	return "ok"
}
