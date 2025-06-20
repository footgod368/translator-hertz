package log

import (
	"os"

	"github.com/bytedance/gg/gslice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// func Init() {
// 	logPath := filepath.Join(filepath.Dir(os.Args[0]), "log", "app.log")
// 	f, _ := os.Create(logPath)
// 	defer f.Close()
// 	hlog.SetOutput(f)
// }

func Init() {
	if gslice.Contains(os.Args, "-debug") {
		hlog.SetLevel(hlog.LevelDebug)
	} else {
		hlog.SetLevel(hlog.LevelInfo)
	}
}
