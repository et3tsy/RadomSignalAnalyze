## 设计说明:基于RabbitMQ作为消息中间件的信号分析系统

项目地址: https://github.com/et3tsy/RadomSignalAnalyze



### 项目需求

<a href="https://imgtu.com/i/LcXxa9"><img src="https://s1.ax1x.com/2022/04/21/LcXxa9.png" alt="LcXxa9.png" border="0" /></a>



### 微服务以及消息队列关系图

<a href="https://imgtu.com/i/LcXgv8"><img src="https://s1.ax1x.com/2022/04/21/LcXgv8.png" alt="LcXgv8.png" border="0" /></a>



### 设计思路：

`Create`是信号产生端，他将消息打包发给交换器 `signal`，交换器 `signal` 将消息发送给分析端 `analyze`和可视化端 `visualize`。

分析端 `analyze` 获得消息后，维护一个固定大小的队列(该大小可通过设置 `settings` 中 `config.yaml` 的参数 `analyze.size` 改变)，花费 $O(1)$ 的时间完成均值和方差的维护，并将结果发布到 `analyze_to_visualize` 中。

可视化端从 `signal_to_visualize` 获取信号数据，并使用动态开点线段树维护各信号出现的频率次数，在获取前端分段参数后，对值域分段，每段数据采用线段树区间询问的方式统计出现的次数。同时，从 `analyze_to_visualize` 获取结果集，向前端进一步完成传递。

前端使用 `Vue.js` 进行展示，单独跑一个服务器，用代理解决跨域问题。

### Create

```
|--messsageQueue 设置消息中间件
    |--rabbitMQ.go 设置rabbitMQ
|--logger 日志记录
    |--logger.go 设置日志记录
|--models 定义模型
    |--signal.go 定义信号结构体
|--random 产生正态分布的信号
    |--rand.go 随机获得信号 
|--settings
    |--config.yaml 配置环境属性
    |--settings.go 读取环境配置
|--create.go 信号产生端主函数
```



### Analyze

```
|--calculate 动态维护期望和方差
    |--caluculate.go 设置了一定容量的队列，在插入\弹出时，O(1)维护对应期望和方差
|--messsageQueue 设置消息中间件
    |--rabbitMQ.go 设置rabbitMQ
|--logger 日志记录
    |--logger.go 设置日志记录
|--queue 数据结构--队列
    |--queue.go 队列
|--models 定义模型
    |--signal.go 定义信号、信号分析结果结构体
|--settings
    |--config.yaml 配置环境属性
    |--settings.go 读取环境配置
|--analyze.go 信号分析端主函数
```



### Visualize

```
|--controller 控制层
    |--signal.go 封装路由相应方法，完成前端参数校验、往前端传递处理结果
|--ds 数据结构
    |--segmentTree.go 采用动态开点线段树统计。新信号，O(logN)插入。分段统计上区间查询，单次复杂度为O(logN)
|--messsageQueue 设置消息中间件
    |--rabbitMQ.go 设置rabbitMQ
|--logic 逻辑层
    |--statistics.go 数据统计
    |--time.go 时间转换,方便前端进行展示
|--logger 日志记录
    |--logger.go 设置日志记录
|--models 定义模型
    |--signal.go 定义信号、信号分析结果结构体
|--settings 设置
    |--config.yaml 配置环境属性
    |--settings.go 读取环境配置
|--routes 路由设置
    |--routes.yaml 将路由和路由方法绑定
|--visual.go 可视化形成端主函数
```



### 效果展示

```
https://www.bilibili.com/video/BV1zT4y1r73c/
```

