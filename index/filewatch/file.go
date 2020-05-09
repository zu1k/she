package filewatch

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zu1k/she/common/tools"

	"github.com/zu1k/she/persistence"

	"github.com/fsnotify/fsnotify"
	C "github.com/zu1k/she/constant"
	"github.com/zu1k/she/index/fullline"
)

type Watch struct {
	watch *fsnotify.Watcher
}

func DoWatch() {
	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir(C.Path.OriginDir())
	select {}
}

//监控目录
func (w *Watch) watchDir(dir string) {
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
						} else {
							fmt.Println(ev.Name)
							sources, err := persistence.GetSourceSByName(tools.Path2Name(ev.Name))
							if err != nil {
								fullline.ParseAndIndex(ev.Name)
							} else {
								found := false
								for _, source := range sources {
									if source.OriginFile == ev.Name {
										found = true
										break
									}
								}
								if !found {
									fullline.ParseAndIndex(ev.Name)
								}
							}
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.watch.Remove(ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
