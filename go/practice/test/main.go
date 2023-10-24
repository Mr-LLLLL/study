package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"sync"

	_ "net/http/pprof"
	"os"
	"runtime"

	"mytest/pool"

	"gitlab.cd.anpro/go/common/file"
)

const (
	MB = 1024 * 1024
	GB = 1024 * MB
)

type Resp struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"errMsg"`
}

type Res struct {
	Name string
}

type EngineTaskUpdateBO struct {
	ID           int `yaml:"taskId"`
	TaskProgress int `yaml:"taskProgress"`
}

type LiveNotifyReq struct {
	EngineTaskUpdateBOList []EngineTaskUpdateBO `yaml:"engineScaTaskUpdateBOList"`
}

type valuemap struct {
	ctx context.Context
	res string
}

type str struct{}

func (str) Print() {}

type str1 struct{}

func (str1) Print() {}

type In interface {
	Print()
}

func test() {
	t := "test"
	defer fmt.Println(t)
	t = "hello"
	defer fmt.Println(t)
}

func link() {
	os.Link(os.Args[1], os.Args[2])
	os.Remove(os.Args[1])
}

type ProgramMem struct {
	HeapAlloc  uint64
	StackAlloc uint64
}

func GetProgramMem() ProgramMem {
	ms := new(runtime.MemStats)
	runtime.ReadMemStats(ms)

	return ProgramMem{
		HeapAlloc:  ms.HeapAlloc / MB,
		StackAlloc: ms.StackSys / MB,
	}
}

func recursive(s string, dep int) {
	if dep > 1000 {
		return
	}

	fmt.Println(GetProgramMem())
	recursive(s, dep+1)
}

type S struct {
	Key []struct {
		Key string
	}
}

func initFun() func() int {
	s := 5
	return func() int {
		return s
	}
}

var getS = initFun()

type cmdLimiter chan struct{}

var (
	aCmdLimiter cmdLimiter
	aCmdOnce    sync.Once
)

func getCmdLimiter() cmdLimiter {
	aCmdOnce.Do(func() {
		aCmdLimiter = make(chan struct{}, 1)
	})

	return aCmdLimiter
}

func (c cmdLimiter) limit() {
	c <- struct{}{}
}

func (c cmdLimiter) free() {
	<-c
}

type person struct {
	name string
}

func gitErrHandle(msg []byte, err error) error {
	if err == nil {
		return nil
	}

	m := err.Error()
	switch {
	case strings.Contains(m, "Warning") || strings.Contains(m, "warning"):
		return nil
	case bytes.Contains(msg, []byte("Access denied. The provided password or token is incorrect or your account has 2FA enabled and you must use a personal access token instead of a password")):
		return fmt.Errorf("")
	}

	return fmt.Errorf("%s", msg)
}

func main() {
	fmt.Println("ksjdfks")
	pool.Test()
}

func test111(path string) ([]byte, error) {
	os.MkdirAll(path, 0777)

	cmd := exec.CommandContext(context.Background(), "git", "clone", "-q", "-b", "dev", "--single-branch", "-j", "10", "--recursive", "--no-tags", "https://111:1111@gitee.com/acgist/taoyao.git", "./")
	cmd.Dir = path

	outbuf := bytes.Buffer{}
	errbuf := bytes.Buffer{}
	cmd.Stdout = &outbuf
	cmd.Stderr = &outbuf

	err := cmd.Run()
	if err != nil {
		// if err, return stdout and stderr
		return append(outbuf.Bytes(), errbuf.Bytes()...), err
	}

	if errbuf.Len() != 0 {
		return outbuf.Bytes(), fmt.Errorf("%s", errbuf.String())
	}

	return outbuf.Bytes(), nil
}

func unzipAllZipInDir(ctx context.Context, path string, unzipOpt *file.UnzipOpt) error {
	if unzipOpt == nil {
		return nil
	}

	files, err := file.GetAllFileInDir(path, file.GetVcsFeatureDirName()...)
	if err != nil {
		return err
	}

	zipFiles := make([]string, 0)
	for _, v := range files {
		if file.IsZipOrArchiveFile(v) {
			zipFiles = append(zipFiles, v)

			_, err := os.Open(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Recover(ctx ...context.Context) {
	if err := recover(); err != nil {
		if len(ctx) > 0 && ctx[0] != nil {
			fmt.Println("hello")
		} else {
			fmt.Println("word")
		}
	}
}
func SetValueByTag(s any, tag string) error {
	ref := reflect.ValueOf(s)
	if ref.Kind() != reflect.Pointer {
		return fmt.Errorf("s must a pointer")
	}
	ref = ref.Elem()

	err := assignTo(ref, tag)
	if err != nil {
		return err
	}

	return nil
}
func assignTo(ref reflect.Value, tag string) error {
	for i := 0; i < ref.NumField(); i++ {
		v := ref.Field(i)

		if v.Kind() == reflect.Pointer {
			v.Set(reflect.New(v.Type().Elem()))
		}

		if v.Kind() == reflect.Struct {
			err := assignTo(v, tag)
			if err != nil {
				return err
			}
			continue
		}

		tagValue := ref.Type().Field(i).Tag.Get(tag)
		if tagValue == "" {
			continue
		}

	}

	return nil
}

func add(i, j int) int {
	return i + j
}

type config struct {
	Mode   string `yaml:"mode" default:"release"` // opt:1-release 2-test 3-debug
	TmpDir string `yaml:"tmp_dir" default:"/tmp/sca"`

	HttpServer httpServer `yaml:"http_server"`
	HttpClient httpClient `yaml:"http_client"`
	Log        log1       `yaml:"log"`
	Hdfs       hdfs       `yaml:"hdfs"`

	WebTask webTask `yaml:"web_task"`
}

type log1 struct {
	Level          string `yaml:"level" default:"info"` // opt:1-trace 2-debug 3-info 4-warn 日志级别,只输出配置及其以上日志
	Directory      string `yaml:"directory" default:"/var/xm/sca/log"`
	SavedDays      int    `yaml:"saved_days" default:"15"`           // 日志存储时间
	MaxSizeForEach int    `yaml:"max_size_for_each" default:"40000"` // 单条日志最大限制，单位：Byte，默认40KB
}

type httpServer struct {
	ListenPort   int `yaml:"listen_port" default:"5000"`
	ReadTimeout  int `yaml:"read_timeout" default:"300"`  // http服务端读取所有数据超时时间，300s
	WriteTimeout int `yaml:"write_timeout" default:"300"` // http服务端写入所有数据超时时间，300s
}

type httpClient struct {
	SkipHttpsVerify bool `yaml:"skip_https_verify" default:"false"`
	Timeout         int  `yaml:"timeout" default:"600"` // http客户端连接超时时间，600s
}

type webTask struct {
	Host             string `yaml:"host" default:"http://127.0.0.1"`
	Port             int    `yaml:"port" default:"8090"`
	ParallelCount    int    `yaml:"parallel_count" default:"5"`
	HeartBeatInteral int    `yaml:"heart_beat_interal" default:"10"` // 10 second
	ListPullInteral  int    `yaml:"list_pull_interal" default:"10"`  // 10 second
}

type hdfs struct {
	Host string `yaml:"host" default:"http://10.1.2.32"`
	Port int    `yaml:"port" default:"9870"`
}
