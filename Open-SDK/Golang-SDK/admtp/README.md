# 中国行政区划 Golang SDK (`admtp`)

基于民政部公开数据的中国行政区划查询 SDK，支持省份、城市、区县的快速查询，名称大小写不敏感、名称变体（如“北京市”“内蒙古自治区”“宁夏省”）也可识别。

## 安装

- 使用方式：
  - `go get github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp`

## 快速开始

```go
package main

import (
    "fmt"
    admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func main() {
    d, err := admtp.NewDivision()
    if err != nil { panic(err) }

    p, ok := d.GetProvinceByName("北京市")
    fmt.Println(p.ID, ok) // 110000 true

    c, ok := d.GetCityByName("石家庄")
    fmt.Println(c.ID, ok) // 130100 true

    k, ok := d.GetCountyByName("海淀区")
    fmt.Println(k.ID, ok) // 110108 true

    fmt.Println(len(d.ListProvinces()))
    fmt.Println(len(d.GetCitiesByProvince("130000")))
    fmt.Println(len(d.GetCountiesByCity("130100")))
}
```

## 主要功能

- 根据 ID / 名称获取省、市、区县
- 名称大小写不敏感，支持中文与英文名
- 省份名称变体识别（北京市/自治区/特别行政区、自动补“省”后缀）
- 列表查询与层级关系（省→市→区县）
- 模糊搜索（中文/英文）

接口定义参考：`data/Open-SDK/Golang-SDK/admtp/division.go:1`

## 数据加载

- 包内默认加载 `admtp/data` 下的 JSON 文件：
  - `province.json`、`cities/*.json`、`counties/*/counties.json`
- 运行时路径解析：`data/Open-SDK/Golang-SDK/admtp/loader.go:1` 会优先使用包内 `admtp/data`；若不存在，回退到 Python SDK 的数据目录。
- 若希望 Go 包完全自包含，请将数据目录复制到 `data/Open-SDK/Golang-SDK/admtp/data` 下（已提供完整副本）。

## 许可证

[Apache License 2.0](https://github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties?tab=Apache-2.0-1-ov-file#readme)
