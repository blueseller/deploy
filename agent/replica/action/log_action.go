package action

import (
	"context"
	"io/ioutil"

	"github.com/blueseller/deploy.git/logger"
)

type LogAction struct {
}

type DirFile struct {
	IsDir bool
	Name  string
}

func NewLogAction() *LogAction {
	return new(LogAction)
}

func (l *LogAction) GetLogCatalogList(ctx context.Context, basePath string) []*DirFile {
	// 获取basePath路径下的所有文件目录 和 文件
	// ls -l
	basePath = "/home/liukai02/go/src/github.com/blueseller/deploy.git"

	fileInfoList, err := ioutil.ReadDir(basePath)
	if err != nil {
		logger.GetContextLogger(ctx).Errorf("read dir by path is error: %+v", err)
	}
	list := make([]*DirFile, 0, len(fileInfoList))
	for _, v := range fileInfoList {
		dirFile := new(DirFile)
		dirFile.Name = v.Name()
		if v.IsDir() {
			dirFile.IsDir = true
		}
		list = append(list, dirFile)
	}
	return list
}

func (l *LogAction) StartLogTail2000(ctx context.Context, path string) []byte {
	// 获取某个文件的最后2000行
	// tail -2000
	return []byte("")
}

func (l *LogAction) StartLogTailf(ctx context.Context, path string) []byte {
	// 获取某个文件APPEND的文本
	// tail -f
	return []byte("")
}
