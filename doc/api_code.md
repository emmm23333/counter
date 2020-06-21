## 请求参数
***
key:"rect"  
value:
```
{
    "x": 0,
    "y": 0,
    "width": 480,
    "height": 960
}
```
***
key:"file"  
value:binary
***

## 响应
```
{
    "code": 200，
    "msg": "" #code不为200时的错误信息
    "rects:[
        {
            "x": 0,
            "y": 0,
            "width": 480,
            "height": 960
        }
    ]
}
```

## 错误码 code
- 200: 成功
- 201: form里没有文件
- 202: 保存文件失败
- 203: form读取失败
- 204: form里没有rect
- 205: rect反序列化失败
- 206: 算法处理失败（错误信息在msg里）
- 207: license校验失败