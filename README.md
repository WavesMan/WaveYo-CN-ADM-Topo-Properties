# WaveYo 中国行政区划数据集

作者：Waves_Man（WaveYo）

LICENSE:[Apache License 2.0](LICENSE.md)

## 数据结构
- `province.json`：省级区划列表（`id`、`zh`、`en`）。
- `cities/<province_id>.json`：对应省的地级市列表。
- `counties/<province_id>/counties.json`：对应省的县区级列表（含直辖市的区县）。

## 命名规则
- 中文：去除地级“市”后缀，例如“合肥市”→“合肥”。
- 英文：基于拼音首字母大写；若以`shi`结尾则移除，例如“Shijiazhuangshi”→“Shijiazhuang”。
- 特例：`香港/Hong Kong`、`澳门/Macao`、`西藏/Tibet`、`内蒙古/Inner Mongolia`、`西安/Xi'an`等按内置映射处理。

## 快捷跳转
- [省级行政区划(含自治区、特别行政区、直辖市、台湾省)](province.json)

- 北京市（直辖市，无市级区划）
  - [县区行政区划](counties/110000/counties.json)
- 天津市（直辖市，无市级区划）
  - [县区行政区划](counties/120000/counties.json)
- 河北省
  - [市级行政区划](cities/130000.json)
  - [县区行政区划](counties/130000/counties.json)
- 山西省
  - [市级行政区划](cities/140000.json)
  - [县区行政区划](counties/140000/counties.json)
- 内蒙古自治区
  - [市级行政区划](cities/150000.json)
  - [县区行政区划](counties/150000/counties.json)
- 辽宁省
  - [市级行政区划](cities/210000.json)
  - [县区行政区划](counties/210000/counties.json)
- 吉林省
  - [市级行政区划](cities/220000.json)
  - [县区行政区划](counties/220000/counties.json)
- 黑龙江省
  - [市级行政区划](cities/230000.json)
  - [县区行政区划](counties/230000/counties.json)
- 上海市（直辖市，无市级区划）
  - [县区行政区划](counties/310000/counties.json)
- 江苏省
  - [市级行政区划](cities/320000.json)
  - [县区行政区划](counties/320000/counties.json)
- 浙江省
  - [市级行政区划](cities/330000.json)
  - [县区行政区划](counties/330000/counties.json)
- 安徽省
  - [市级行政区划](cities/340000.json)
  - [县区行政区划](counties/340000/counties.json)
- 福建省
  - [市级行政区划](cities/350000.json)
  - [县区行政区划](counties/350000/counties.json)
- 江西省
  - [市级行政区划](cities/360000.json)
  - [县区行政区划](counties/360000/counties.json)
- 山东省
  - [市级行政区划](cities/370000.json)
  - [县区行政区划](counties/370000/counties.json)
- 河南省
  - [市级行政区划](cities/410000.json)
  - [县区行政区划](counties/410000/counties.json)
- 湖北省
  - [市级行政区划](cities/420000.json)
  - [县区行政区划](counties/420000/counties.json)
- 湖南省
  - [市级行政区划](cities/430000.json)
  - [县区行政区划](counties/430000/counties.json)
- 广东省
  - [市级行政区划](cities/440000.json)
  - [县区行政区划](counties/440000/counties.json)
- 广西壮族自治区
  - [市级行政区划](cities/450000.json)
  - [县区行政区划](counties/450000/counties.json)
- 海南省
  - [市级行政区划](cities/460000.json)
  - [县区行政区划](counties/460000/counties.json)
- 重庆市（直辖市，无市级区划）
  - [县区行政区划](counties/500000/counties.json)
- 四川省
  - [市级行政区划](cities/510000.json)
  - [县区行政区划](counties/510000/counties.json)
- 贵州省
  - [市级行政区划](cities/520000.json)
  - [县区行政区划](counties/520000/counties.json)
- 云南省
  - [市级行政区划](cities/530000.json)
  - [县区行政区划](counties/530000/counties.json)
- 西藏自治区
  - [市级行政区划](cities/540000.json)
  - [县区行政区划](counties/540000/counties.json)
- 陕西省
  - [市级行政区划](cities/610000.json)
  - [县区行政区划](counties/610000/counties.json)
- 甘肃省
  - [市级行政区划](cities/620000.json)
  - [县区行政区划](counties/620000/counties.json)
- 青海省
  - [市级行政区划](cities/630000.json)
  - [县区行政区划](counties/630000/counties.json)
- 宁夏回族自治区
  - [市级行政区划](cities/640000.json)
  - [县区行政区划](counties/640000/counties.json)
- 新疆维吾尔自治区
  - [市级行政区划](cities/650000.json)
  - [县区行政区划](counties/650000/counties.json)
- 台湾省
  - 数据缺失
- 香港特别行政区
  - 数据缺失
- 澳门特别行政区
  - 数据缺失

## 数据来源与抓取
- 来源：民政部区划公开数据 `http://xzqh.mca.gov.cn/map`

## 示例字段
```
{
  "type": "Polygon",
  "properties": {
    "NAME": "瑶海区",
    "QUHUADAIMA": "340102",
    "FillColor": "#ee8b94",
    "strZTValue": ""
  }
}
```
