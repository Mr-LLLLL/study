# 设计模式

 设计模式(design pattern)是对面向对象设计中反复出现的问题的解决方案。这个术语是在1990年代由Erich Gamma等人从建筑设计领域引入到计算机科学中来的。这个术语的含义还存有争议。算法不是设计模式，因为算法致力于解决问题而非设计问题。设计模式通常描述了一组相互紧密作用的类与对象。设计模式提供一种讨论软件设计的公共语言，使得熟练设计者的设计经验可以被初学者和其他设计者掌握。设计模式还为软件重构提供了目标

## 1.Singleton(单例模式)

- 定义：一个类只允许创建一个对象（或者叫实例），那这个类就是一个单例类

- 用处：从业务概念上，有些数据在系统中只应该保存一份，就比较适合设计为单例类

- 唯一性：在多线程（或者协程）保证其唯一性

### 饿汉式

``` go
package singleton

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
    singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
    return singleton
}
```

### 懒汉式

``` go
package singleton

import "sync"

var (
    lazySingleton *Singleton
    once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
    if lazySingleton == nil {
        once.Do(func() {
            lazySingleton = &Singleton{}
        })
    }
    return lazySingleton
}
```

## 2.Factory Method（工厂方法）和Abstract Factory(抽象工厂）

- 封装变化： 创建逻辑有可能变化，封装成工厂类之后，创建逻辑的变更对调用者透明

- 代码复用： 创建代码抽离到独立的工厂类之后可以复用

- 隔离复杂性： 封装复杂的创建逻辑，调用者无需了解如何创建对象

- 控制复杂度： 将创建代码抽离出来，让原本的函数或类职责更单一，代码更简洁

### 简单工厂

- 实现简单，有较多的if else分之，适用于改动不频繁的方法
- 由于 Go 本身是没有构造函数的，一般而言我们采用 NewName 的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式

``` go
package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
    Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
    panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
    panic("implement me")
}

// NewIRuleConfigParser NewIRuleConfigParser
func NewIRuleConfigParser(t string) IRuleConfigParser {
    switch t {
    case "json":
        return jsonRuleConfigParser{}
    case "yaml":
        return yamlRuleConfigParser{}
    }
    return nil
}
```

### 工厂方法

- 当对象的创建逻辑比较复杂，不只是简单的new以下就可以，而是要组合其它类对象，做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂

``` go
// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
    CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
    return yamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
    return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
    switch t {
    case "json":
        return jsonRuleConfigParserFactory{}
    case "yaml":
        return yamlRuleConfigParserFactory{}
    }
    return nil
}
```

### 抽象工厂

- 可以让一个工厂负责创建多个不同类型的对象，而不是只创建一种类型对象，这样就可以有效地减少工厂类的个数

``` go
package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
    Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct{}

// Parse Parse
func (j jsonRuleConfigParser) Parse(data []byte) {
    panic("implement me")
}

// ISystemConfigParser ISystemConfigParser
type ISystemConfigParser interface {
    ParseSystem(data []byte)
}

// jsonSystemConfigParser jsonSystemConfigParser
type jsonSystemConfigParser struct{}

// Parse Parse
func (j jsonSystemConfigParser) ParseSystem(data []byte) {
    panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
    CreateRuleParser() IRuleConfigParser
    CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
    return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
    return jsonSystemConfigParser{}
}
```

## Builder(建造者模式)

应用场景:

- 类中的属性比较多

- 类的属性之间有一定的依赖关系，或者是约束条件

- 存在必选和非必选的属性

- 希望创建不可变的对象

``` go
package builder

import "fmt"

// ResourcePoolConfigOption option
type ResourcePoolConfigOption struct {
    maxTotal int
    maxIdle  int
    minIdle  int
}

// ResourcePoolConfigOptFunc to set option
type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

// NewResourcePoolConfig NewResourcePoolConfig
func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
    if name == "" {
        return nil, fmt.Errorf("name can not be empty")
    }

    option := &ResourcePoolConfigOption{
        maxTotal: 10,
        maxIdle:  9,
        minIdle:  1,
    }

    for _, opt := range opts {
        opt(option)
    }

    if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
        return nil, fmt.Errorf("args err, option: %v", option)
    }

    if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
        return nil, fmt.Errorf("args err, option: %v", option)
    }

    return &ResourcePoolConfig{
        name:     name,
        maxTotal: option.maxTotal,
        maxIdle:  option.maxIdle,
        minIdle:  option.minIdle,
    }, nil
}
```

## Prototype（原型模式）

- 定义： 利用已有对象（原型）进行复制（拷贝）的方式来创建新对象，以达到节省创建时间的目的

- 使用场景： 对象的创建成本比较大，并且同一个类的不同对象之间的差别不大（大部分字段相同）

``` go
package prototype

import (
    "encoding/json"
    "time"
)

// Keyword 搜索关键字
type Keyword struct {
    word      string
    visit     int
    UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
    var newKeyword Keyword
    b, _ := json.Marshal(k)
    json.Unmarshal(b, &newKeyword)
    return &newKeyword
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
    newKeywords := Keywords{}

    for k, v := range words {
        // 这里是浅拷贝，直接拷贝了地址
        newKeywords[k] = v
    }

    // 替换掉需要更新的字段，这里用的是深拷贝
    for _, word := range updatedWords {
        newKeywords[word.word] = word.Clone()
    }

    return newKeywords
}
```

## 创建型模式总结

- 单例模式： 如果有些数据只应该保存一份，那么就比较适合单例

- 工厂模式： 用于创建同一类型的不同类的对象

  - 简单工厂： 适用于类型改动不频繁的情况
  - 工厂方法： 适用于对象创建逻辑比较复杂的情况
  - 抽象工厂： 适用于同时创建有关联的多个不同类型的对象

- 建造者模式： 适用于类的属性比较多，存在必选和非必选或其他较为复杂的约束条件的情况

- 原型模式： 适用于对象的创建成本较大，并且创建出的对象直接的差别较小的情况(一般用于前端)

## Proxy(代理模式)

- 在不改变原始类代码的情况下，通过引入代理类来给原始类附加功能

``` go
package proxy

import (
    "log"
    "time"
)

// IUser IUser
type IUser interface {
    Login(username, password string) error
}

// User 用户
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
    // 不实现细节
    return nil
}

// UserProxy 代理类
type UserProxy struct {
    user *User
}

// NewUserProxy NewUserProxy
func NewUserProxy(user *User) *UserProxy {
    return &UserProxy{
        user: user,
    }
}

// Login 登录，和 user 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
    // before 这里可能会有一些统计的逻辑
    start := time.Now()

    // 这里是原有的业务逻辑
    if err := p.user.Login(username, password); err != nil {
        return err
    }

    // after 这里可能也有一些监控统计的逻辑
    log.Printf("user login cost time: %s", time.Now().Sub(start))

    return nil
}
```

## Bridge(桥接模式)

- 将抽象和实现解耦，让他们可以独立变化

- 一个类存在两个（或多个）独立变化的维度，我们通过组合的方式，让这两个（或多个）维度可以独立进行拓展

``` go
package bridge

// IMsgSender IMsgSender
type IMsgSender interface {
    Send(msg string) error
}

// EmailMsgSender 发送邮件
// 可能还有 电话、短信等各种实现
type EmailMsgSender struct {
    emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
    return &EmailMsgSender{emails: emails}
}

// Send Send
func (s *EmailMsgSender) Send(msg string) error {
    // 这里去发送消息
    return nil
}

// INotification 通知接口
type INotification interface {
    Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
    sender IMsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
    return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
    return n.sender.Send(msg)
}
```

## Decorator(装饰器模式)

- 装饰器模式主要解决继承关系过于复杂的问题，通过组合来替代继承，他主要的作用是给原始类添加增强功能

- 和代理模式的区别

  - 代码形式上几乎没有区别
  - 代理模式主要给原始类添加无关的功能
  - 装饰器主要给原始类增强功能，添加的功能都是有关联的

``` go
package decorator

// IDrawer IDraw
type IDraw interface {
    Draw() string
}

// Square 正方形
type Square struct{}

// Draw Draw
func (s Square) Draw() string {
    return "this is a square"
}

// ColorSquare 有颜色的正方形
type ColorSquare struct {
    square IDraw
    color  string
}

// NewColorSquare NewColorSquare
func NewColorSquare(square IDraw, color string) ColorSquare {
    return ColorSquare{color: color, square: square}
}

// Draw Draw
func (c ColorSquare) Draw() string {
    return c.square.Draw() + ", color is " + c.color
}
```

## Adapter(适配器模式)

- 顾名思义，这个模式就是用来做适配的，它将不兼容的接口转换为可兼容的接口，让原本由于接口不兼容而不能一起工作的类可以一起工作

- 和其他模式的区别

  - 代理模式： 在不改变原始类接口的条件下，为原始类定义一个代理类，主要目的是控制访问，而且加强功能
  - 桥接模式： 目的是将接口部分和实现部分分离，从而让他们可以较为容易，也相对独立地加以改变
  - 装饰器： 在不改变原始类接口的情况下，对原始类功能进行增强
  - 适配器： 是一种事后的补救策略，适配器提供跟原始类不同的接口，而代理，装饰器提供的都是跟原始类相同的接口

``` go
package adapter

import "fmt"

// ICreateServer 创建云主机
type ICreateServer interface {
    CreateServer(cpu, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance 启动实例
func (c *AWSClient) RunInstance(cpu, mem float64) error {
    fmt.Printf("aws client run success, cpu： %f, mem: %f", cpu, mem)
    return nil
}

// AwsClientAdapter 适配器
type AwsClientAdapter struct {
    Client AWSClient
}

// CreateServer 启动实例
func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
    a.Client.RunInstance(cpu, mem)
    return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
    fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
    return nil
}

// AliyunClientAdapter 适配器
type AliyunClientAdapter struct {
    Client AliyunClient
}

// CreateServer 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
    a.Client.CreateServer(int(cpu), int(mem))
    return nil
}
```

## Facade(门面模式)

- 门面模式为子系统提供一组统一的接口，定义一组高层接口让子系统更易用
- 将几个细粒度的接口包装成一个接口

``` go
package facade

// IUser 用户接口
type IUser interface {
    Login(phone int, code int) (*User, error)
    Register(phone int, code int) (*User, error)
}

// IUserFacade 门面模式
type IUserFacade interface {
    LoginOrRegister(phone int, code int) error
}

// User 用户
type User struct {
    Name string
}

// UserService UserService
type UserService struct {}

// Login 登录
func (u UserService) Login(phone int, code int) (*User, error) {
    // 校验操作 ...
    return &User{Name: "test login"}, nil
}

// Register 注册
func (u UserService) Register(phone int, code int) (*User, error) {
    // 校验操作 ...
    // 创建用户
    return &User{Name: "test register"}, nil
}

// LoginOrRegister 登录或注册
func (u UserService)LoginOrRegister(phone int, code int) (*User, error) {
    user, err := u.Login(phone, code)
    if err != nil {
        return nil, err
    }

    if user != nil {
        return user, nil
    }

    return u.Register(phone, code)
}
```

## Composite(组合模式)

- 将一组对象组织成数形结构，以表示“部分-整体”的层次接口，组合让调用者可以统一单个对象和组合对象的处理逻辑

``` go
package composite

// IOrganization 组织接口，都实现统计人数的功能
type IOrganization interface {
    Count() int
}

// Employee 员工
type Employee struct {
    Name string
}

// Count 人数统计
func (Employee) Count() int {
    return 1
}

// Department 部门
type Department struct {
    Name string

    SubOrganizations []IOrganization
}

// Count 人数统计
func (d Department) Count() int {
    c := 0
    for _, org := range d.SubOrganizations {
        c += org.Count()
    }
    return c
}

// AddSub 添加子节点
func (d *Department) AddSub(org IOrganization) {
    d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization 构建组织架构 demo
func NewOrganization() IOrganization {
    root := &Department{Name: "root"}
    for i := 0; i < 10; i++ {
        root.AddSub(&Employee{})
        root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
    }
    return root
}
```

## Flyweight(享元模式)

- 享元模式的意图是复用对象，节省内存，前提是享元对象是不可变对象

- 对比
  
  - 单例模式
    - 单例模式： 一个类只能创建一个对象
    - 享元模式： 一个类可以创建多个对象，对象被多处引用，为了对象复用，节省内存
    - 多例模式： 为了限制对象的个数

  - 缓存
    - 缓存： 为了提高访问效率
    - 享元： 为了复用

  - 对象池
    - 对象池： 为了节省时间，例如连接池中创建关闭连接的时间
    - 享元： 主要为了节省空间，在整个生命周期被所有共享者共同使用

``` go
package flyweight

var units = map[int]*ChessPieceUnit{
    1: {
        ID:    1,
        Name:  "車",
        Color: "red",
    },
    2: {
        ID:    2,
        Name:  "炮",
        Color: "red",
    },
    // ... 其他棋子
}

// ChessPieceUnit 棋子享元
type ChessPieceUnit struct {
    ID    uint
    Name  string
    Color string
}

// NewChessPieceUnit 工厂
func NewChessPieceUnit(id int) *ChessPieceUnit {
    return units[id]
}

// ChessPiece 棋子
type ChessPiece struct {
    Unit *ChessPieceUnit
    X    int
    Y    int
}

// ChessBoard 棋局
type ChessBoard struct {
    chessPieces map[int]*ChessPiece
}

// NewChessBoard 初始化棋盘
func NewChessBoard() *ChessBoard {
    board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
    for id := range units {
        board.chessPieces[id] = &ChessPiece{
            Unit: NewChessPieceUnit(id),
            X:    0,
            Y:    0,
        }
    }
    return board
}

// Move 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
    c.chessPieces[id].X = x
    c.chessPieces[id].Y = y
}
```

<center>
  <font size=5>
    谢谢大家的参会，如上述中有错误,不足或者不解之处，请提出，让我们相互学习，共同成长
  </font>
</center>

