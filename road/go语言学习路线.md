```
Golang 云原生开发工程师学习图谱
│
├── 1. 计算机与前置基础 (学习资源:https://xiaolincoding.com/)
│   ├── Linux 操作系统
│   │   ├── 常用基础命令 (文件操作、权限控制、系统状态查看)
│   │   ├── 用户与进程管理
│   │   ├── Shell 脚本编程基础
│   │   └── 网络抓包与状态查看 (tcpdump, netstat, lsof)
│   ├── 计算机网络基础
│   │   ├── TCP/IP 五层模型
│   │   ├── TCP 三次握手与四次挥手、可靠传输机制
│   │   ├── HTTP/1.1、HTTP/2.0 与 HTTPS 原理
│   │   └── DNS 域名解析过程
│   └── 数据结构与算法 (基础)
│       ├── 数组、链表、栈、队列、哈希表
│       └── 常见排序与查找算法 (冒泡、快排、二分查找)
│
├── 2. Go 语言核心(学习资源:入门https://studygolang.com/,https://golang.halfiisland.com/guide.html 进阶https://draven.co/golang/, 视频:B站 UP主“幼麟实验室”)
│   ├── 基础语法
│   │   ├── 变量、常量与基本数据类型
│   │   ├── 控制流程 (if、for、switch)
│   │   ├── 函数 (多返回值、匿名函数、闭包、defer 延迟调用)
│   │   └── 错误处理 (error 接口规范、panic 与 recover)
│   ├── 复合数据类型与底层机制
│   │   ├── 数组 (Array) 与 切片 (Slice) 底层结构及扩容原理
│   │   ├── 字典 (Map) 哈希链表实现与并发不安全性
│   │   ├── 结构体 (Struct) 与内存对齐
│   │   └── 指针 (Pointer) 与内存寻址
│   ├── 面向对象与抽象编程
│   │   ├── 方法 (Method) 与接收者 (Receiver)
│   │   ├── 接口 (Interface) 定义与隐式实现机制
│   │   ├── 空接口 (any) 与类型断言 (Type Assertion)
│   │   └── 结构体嵌套 (Go 语言的组合机制)
│   ├── 并发编程核心 (核心重点)
│   │   ├── Goroutine 协程基础与生命周期
│   │   ├── Channel 通道通信机制 (无缓冲、有缓冲、死锁、select)
│   │   ├── 并发同步机制 (sync.Mutex、sync.WaitGroup、sync.Once、sync.Map)
│   │   └── Context 上下文管理与超时/取消控制
│   ├── 底层原理与进阶机制
│   │   ├── 并发调度模型 (GMP 模型: G协程、M线程、P处理器核心原理解析)
│   │   ├── 垃圾回收机制 (GC: 三色标记法与混合写屏障)
│   │   ├── 内存分配机制 (TCMalloc 思想、mspan 结构)
│   │   └── 反射 (reflect) 基础与应用场景
│   └── 工程化与标准库
│       ├── 模块管理 (Go Modules)
│       ├── 常用标准库 (fmt, io, net/http, os, time, encoding/json)
│       └── 测试驱动开发 (testing 包、pprof 性能分析调优)
│
├── 3. 后端研发基础与微服务架构(Gin框架,Gorm框架,go-zero框架文档)
│   ├── Web 框架与应用
│   │   ├── 路由框架应用 (Gin 或 Hertz)
│   │   ├── 中间件 (Middleware) 设计模式
│   │   ├── 鉴权机制 (JWT、OAuth2)
│   │   └── RESTful API 设计规范
│   ├── 数据库与持久化
│   │   ├── 关系型数据库 MySQL 原理 (B+树索引、事务隔离级别、锁)
│   │   ├── SQL 工具与 ORM 库 (GORM、sqlx)
│   │   └── Redis 缓存技术 (核心数据类型、RDB/AOF 持久化、缓存穿透/击穿/雪崩)
│   ├── 微服务通信基础
│   │   ├── RPC (远程过程调用) 原理概念
│   │   ├── Protobuf (Protocol Buffers) 语法与序列化机制
│   │   └── gRPC 框架应用 (一元调用、流式通信、拦截器)
│
├── 4. 容器化技术 (Docker生态)(学习资源:https://docker-practice.github.io/zh-cn/)
│   ├── 核心概念与实操命令
│   │   ├── 镜像 (Image)、容器 (Container)、仓库 (Registry) 体系
│   │   └── 容器生命周期管理指令 (run, start, stop, rm, logs, exec)
│   ├── 镜像构建与规范
│   │   ├── Dockerfile 指令详解 (FROM, RUN, CMD, ENTRYPOINT, ENV)
│   │   └── 镜像优化技术 (多阶段构建 Multi-stage build、体积瘦身)
│   ├── 容器持久化与网络
│   │   ├── 存储卷挂载 (Volumes、Bind mounts)
│   │   └── 容器网络模型 (Bridge、Host、None、Overlay 网络)
│   └── 容器底层隔离原理
│       ├── 视图隔离: Namespace (PID, Mount, Network, UTS, IPC, User)
│       ├── 资源限制: Cgroups (CPU、内存配额)
│       └── 文件系统: UnionFS (联合文件系统与写时复制 CoW)
│
├── 5. Kubernetes (K8s) 云原生平台核心(学习资源:入门https://kubernetes.io/zh-cn/ 进阶极客时间 - 张磊的《深入剖析 Kubernetes》)
│   ├── 集群架构与核心组件流转
│   │   ├── 控制平面 (Master): API Server (核心网关), etcd (状态存储), Scheduler (调度器), Controller Manager (控制器)
│   │   └── 工作节点 (Node): Kubelet (节点代理), Kube-proxy (网络代理), Containerd (容器运行时)
│   ├── 核心资源对象 (Workloads)
│   │   ├── 最小调度单元: Pod (生命周期、Init Container、探针 Probe)
│   │   ├── 无状态应用: Deployment, ReplicaSet
│   │   ├── 有状态应用: StatefulSet, Headless Service
│   │   └── 特定节点调度与批处理: DaemonSet, Job, CronJob
│   ├── 服务发现与负载均衡 (Network)
│   │   ├── 集群内网络: Service 机制 (ClusterIP, NodePort, LoadBalancer)
│   │   ├── Kube-proxy 代理模式 (iptables 与 IPVS 原理对比)
│   │   └── 外部流量接入: Ingress 与 Ingress Controller
│   ├── 存储卷与持久化架构 (Storage)
│   │   ├── 基础与宿主存储: emptyDir, hostPath
│   │   └── 持久化存储体系: PV (PersistentVolume), PVC (PersistentVolumeClaim), StorageClass 与 CSI
│   ├── 配置注入与调度策略
│   │   ├── 应用配置去中心化: ConfigMap, Secret
│   │   └── 高级调度规则: 节点亲和性 (Affinity), 污点 (Taint) 与 容忍度 (Toleration)
│   └── 安全与访问控制 (Security)
│       └── 基于角色的访问控制 RBAC (Role, ClusterRole, RoleBinding, ServiceAccount)
│
└── 6. 云原生生态进阶体系
    ├── 可观测性 (Observability)
    │   ├── 指标监控与告警: Prometheus 原理与 PromQL, Alertmanager
    │   ├── 数据可视化看板: Grafana
    │   └── 分布式日志收集: EFK / ELK 架构 (Elasticsearch, Fluentd/Filebeat, Kibana)
    ├── 持续集成与持续部署 (CI/CD)
    │   ├── 版本控制与工作流: Git 与 GitFlow
    │   ├── 流水线编排: GitLab CI 或 Jenkins
    │   └── GitOps 部署理念: ArgoCD 实战
    ├── 服务网格基础设施 (Service Mesh)
    │   └── Istio 架构基础 (Sidecar 代理模式、流量治理规则、服务双向 TLS 认证)
    └── 云原生自定义扩展编程 (Operator 范式)
        ├── K8s API 交互: client-go 工作原理与 Informer 机制
        ├── 资源模型扩展: CRD (Custom Resource Definition) 编写
        └── 控制器开发: Operator 模式原理、Kubebuilder / Operator SDK 框架实战开发
```