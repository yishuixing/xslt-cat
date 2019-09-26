package watch

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/yishuixing/xslt-cat/cat"
	"log"
	"os"
)

func NewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					//log.Println("modified file:", event.Name)
					cat.Cat(os.Args[1])
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	if len(os.Args) < 2 {
		fmt.Println("请输入检测的路径")
		os.Exit(-1)
	}
	fmt.Println("开始监控目录: " + os.Args[1])
	err = watcher.Add(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
