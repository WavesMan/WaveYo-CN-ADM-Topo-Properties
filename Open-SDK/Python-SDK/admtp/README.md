# 中国行政区划 Python SDK

基于民政部公开数据的中国行政区划查询SDK，支持省份、城市、区县的快速查询。

## 安装

```bash
# pip
pip install admtp
# uv pyproject
uv add admtp
# uv pip
uv pip install admtp
```

## 快速开始

```python 
from admtp import Division
初始化
division = Division()
根据ID查询
province = division.get_province_by_id("110000") 
print(province.zh) # 北京
根据名称查询（大小写不敏感）
city = division.get_city_by_name("石家庄") 
print(city.en) # Shijiazhuang
模糊搜索
results = division.search_by_name("西安") 
for item in results: 
print(f"{item.zh} ({item.en})")
```

## 主要功能

- ✅ 根据ID查询省市区县信息
- ✅ 根据名称查询（支持大小写不敏感）
- ✅ 模糊搜索行政区划
- ✅ 获取层级关系（省->市->区县）
- ✅ 完整的数据覆盖（包含港澳台地区）

## SDK文档

### Division 类

#### `get_province_by_id(province_id)`
根据省份ID获取省份信息

#### `get_province_by_name(name)`
根据省份名称获取省份信息（大小写不敏感）

#### `get_city_by_id(city_id)`
根据城市ID获取城市信息

#### `get_city_by_name(name)`
根据城市名称获取城市信息（大小写不敏感）

#### `get_county_by_id(county_id)`
根据区县ID获取区县信息

#### `get_county_by_name(name)`
根据区县名称获取区县信息（大小写不敏感）

#### `search_by_name(name)`
根据名称模糊搜索行政区划

#### `list_provinces()`
获取所有省份列表

#### `list_cities()`
获取所有城市列表

#### `list_counties()`
获取所有区县列表

#### `get_cities_by_province(province_id)`
根据省份ID获取该省的所有城市

#### `get_counties_by_city(city_id)`
根据城市ID获取该市的所有区县

## 许可证

[Apache License 2.0](https://github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties?tab=Apache-2.0-1-ov-file#readme)
