# counter

## 服务功能
客户端传递图片文件和目标区域，服务调用钢管检测算法，返回每根钢管的坐标位置  
详细需求以及api文档见doc目录
## 目录结构
* cmd cobra cli库，提供server和client两种模式
* common 全局变量如日志句柄
* doc 需求文档、api文档
* libobject 钢管检测库、模型、头文件
* service 服务实现

## 运行方法
#### 直接运行
* 服务运行
```
  counter --config=.conf.json run
```
* 客户端运行
```
  counter --config=.conf.json client --file=xxx.jpg
```
  默认配置文件为.conf.json
#### docker运行服务
```
   todo...
```


