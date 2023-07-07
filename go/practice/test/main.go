package main

import (
	"fmt"

	_ "net/http/pprof"
	"os"
	"runtime"

	"golang.org/x/net/context"
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

type In interface {
	Print()
}

type str struct{}

func (str) Print() {}

type str1 struct{}

func (str1) Print() {}

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

func main() {
	for i := 0; i < 10; i++ {
		switch i {
		case 1, 2, 3, 4:
			if i == 1 {
				break
			}
			fmt.Println(i)
		}
	}
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
