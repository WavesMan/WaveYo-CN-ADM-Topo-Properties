import json
import os
import sys
from typing import Dict, List, Any

def get_data_directory() -> str:
    """获取数据目录路径"""
    # 获取当前模块的路径
    current_dir = os.path.dirname(os.path.abspath(__file__))
    # 数据目录应该在 admtp/data 下
    data_dir = os.path.join(current_dir, 'data')
    
    if os.path.exists(data_dir):
        return data_dir
    
    # 如果上面的路径不存在，尝试其他可能的路径
    parent_dir = os.path.dirname(current_dir)
    data_dir = os.path.join(parent_dir, 'data')
    
    if os.path.exists(data_dir):
        return data_dir
    
    raise FileNotFoundError(f"无法找到数据目录: {data_dir}")

def load_json_file(file_path: str) -> Dict[str, Any]:
    """加载JSON文件"""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            return json.load(f)
    except FileNotFoundError:
        raise FileNotFoundError(f"数据文件未找到: {file_path}")
    except json.JSONDecodeError as e:
        raise ValueError(f"JSON解析错误: {file_path}, {str(e)}")

def load_provinces() -> List[Dict[str, Any]]:
    """加载省份数据"""
    data_dir = get_data_directory()
    province_file = os.path.join(data_dir, "province.json")
    data = load_json_file(province_file)
    return data.get("province", [])

def load_cities(province_id: str) -> List[Dict[str, Any]]:
    """加载城市数据"""
    data_dir = get_data_directory()
    cities_file = os.path.join(data_dir, "cities", f"{province_id}.json")
    data = load_json_file(cities_file)
    return data.get("cities", [])

def load_counties(province_id: str) -> List[Dict[str, Any]]:
    """加载区县数据"""
    data_dir = get_data_directory()
    counties_file = os.path.join(data_dir, "counties", province_id, "counties.json")
    data = load_json_file(counties_file)
    return data.get("counties", [])
