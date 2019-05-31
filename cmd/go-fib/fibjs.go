package go_fib

import (
	"flag"
	"github.com/Zheaoli/golang-with-fib/config"
	"github.com/Zheaoli/golang-with-fib/engine"
	"time"
)

func main() {
	cpath := flag.String("c", "", "template file path")
	flag.Parse()
	config.InitConfig(*cpath)
	engine.InitFunction()
	execute()
}

func execute() {
	for {
		engine.MainEngine()
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

}
