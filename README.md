# counter

## 服务功能
客户端传递图片文件和目标区域，服务调用钢管检测算法，返回每根钢管的坐标位置 

## 目录结构
* cmd cobra cli库，提供server和client两种模式
* common 全局变量如日志句柄
* doc 需求文档、api文档
* libobject 钢管检测库、模型、头文件
* service 服务实现

## 编译方法
* 推荐环境：centos8 + go1.14
* 在项目目录下执行: go build 生成counter

## 配置文件(.conf.json)
```
{
    "log": {
        "format": "console", #控制台格式的日志
        "stdout": true, #是否输出到stdout
        "path": "logs/tool.log", #日志位置
        "level": "debug", #日志级别
        "max": 10, #日志最大大小(M)
        "maxAge": 30, #日志最大保存天数
        "localtime": true #采用localtime
    },
    "http": {
        "port": ":8081", #服务监听端口
        "fileKey": "file", #formdata里文件的key
        "uri": "/recognize" #识别钢管功能的请求接口
    },
    "algo": {
        "modelPath": "libobject/models", #算法模型路径
        "tag": "gangguan" #算法标签
    },
    "license": "" #程序license，如果检验失败，会在2020-07-21 12:00:00后无法使用
}
```

## 运行方法
#### 直接运行
* 服务运行
```
  counter --config=.conf.json run
```
* 客户端运行(调试用)
```
  counter --config=.conf.json upload --file=xxx.jpg
```
  默认配置文件为.conf.json



