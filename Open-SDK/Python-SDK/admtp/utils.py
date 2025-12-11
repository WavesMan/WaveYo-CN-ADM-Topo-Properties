"""
工具函数
"""

from typing import Dict, List, Any
from .models import Province, City, County

def build_name_index(items: List[Any], item_type: type) -> Dict[str, str]:
    """
    构建名称索引，支持大小写不敏感查询

    Args:
        items: 行政区划数据列表
        item_type: 数据类型 (Province, City, County)

    Returns:
        Dict[str, str]: 名称到ID的映射
    """
    index = {}

    for item_data in items:
        item_id = item_data["id"]
        zh_name = item_data["zh"]
        en_name = item_data["en"]

        # 添加标准名称
        index[zh_name.lower()] = item_id
        index[en_name.lower()] = item_id

        # 处理特殊变体
        if item_type == Province:
            # 省份特殊处理
            if zh_name == "北京":
                index["北京市".lower()] = item_id
            elif zh_name == "天津":
                index["天津市".lower()] = item_id
            elif zh_name == "上海":
                index["上海市".lower()] = item_id
            elif zh_name == "重庆":
                index["重庆市".lower()] = item_id
            elif zh_name == "香港":
                index["香港特别行政区".lower()] = item_id
            elif zh_name == "澳门":
                index["澳门特别行政区".lower()] = item_id
            elif zh_name == "西藏":
                index["西藏自治区".lower()] = item_id
            elif zh_name == "内蒙古":
                index["内蒙古自治区".lower()] = item_id
            elif not zh_name.endswith(("省", "市", "自治区", "特别行政区")):
                # 为没有后缀的添加常见后缀变体
                index[(zh_name + "省").lower()] = item_id

    return index

def create_province_objects(province_data: List[Dict[str, Any]]) -> List[Province]:
    """创建省份对象列表"""
    return [Province(**data) for data in province_data]

def create_city_objects(city_data: List[Dict[str, Any]]) -> List[City]:
    """创建城市对象列表"""
    return [City(**data) for data in city_data]

def create_county_objects(county_data: List[Dict[str, Any]]) -> List[County]:
    """创建区县对象列表"""
    return [County(**data) for data in county_data]
