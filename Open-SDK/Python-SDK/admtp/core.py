"""
核心功能实现
"""

from typing import Optional, List, Union
from .loader import load_provinces, load_cities, load_counties
from .utils import build_name_index, create_province_objects, create_city_objects, create_county_objects
from .models import Province, City, County

class Division:
    """行政区划查询核心类"""

    def __init__(self):
        """初始化，加载所有数据并构建索引"""
        self._provinces = []
        self._province_map = {}
        self._province_index = {}

        self._cities = []
        self._city_map = {}
        self._city_index = {}

        self._counties = []
        self._county_map = {}
        self._county_index = {}

        self._load_all_data()
        self._build_indexes()

    def _load_all_data(self):
        """加载所有行政区划数据"""
        # 加载省份数据
        province_data = load_provinces()
        self._provinces = create_province_objects(province_data)

        # 构建省份映射
        for province in self._provinces:
            self._province_map[province.id] = province

        # 加载所有城市的数据
        for province in self._provinces:
            try:
                city_data = load_cities(province.id)
                cities = create_city_objects(city_data)
                self._cities.extend(cities)

                # 构建城市映射
                for city in cities:
                    self._city_map[city.id] = city
            except FileNotFoundError:
                # 某些直辖市可能没有城市数据文件
                pass

        # 加载所有区县的数据
        for province in self._provinces:
            try:
                county_data = load_counties(province.id)
                counties = create_county_objects(county_data)
                self._counties.extend(counties)

                # 构建区县映射
                for county in counties:
                    self._county_map[county.id] = county
            except FileNotFoundError:
                pass

    def _build_indexes(self):
        """构建查询索引"""
        from .models import Province, City, County

        # 构建省份名称索引
        province_items = [{"id": p.id, "zh": p.zh, "en": p.en} for p in self._provinces]
        self._province_index = build_name_index(province_items, Province)

        # 构建城市名称索引
        city_items = [{"id": c.id, "zh": c.zh, "en": c.en} for c in self._cities]
        self._city_index = build_name_index(city_items, City)

        # 构建区县名称索引
        county_items = [{"id": c.id, "zh": c.zh, "en": c.en} for c in self._counties]
        self._county_index = build_name_index(county_items, County)

    def get_province_by_id(self, province_id: str) -> Optional[Province]:
        """
        根据ID获取省份信息

        Args:
            province_id: 省份ID

        Returns:
            Province: 省份对象，未找到返回None
        """
        return self._province_map.get(province_id)

    def get_province_by_name(self, name: str) -> Optional[Province]:
        """
        根据名称获取省份信息（大小写不敏感）

        Args:
            name: 省份名称

        Returns:
            Province: 省份对象，未找到返回None
        """
        province_id = self._province_index.get(name.lower())
        return self._province_map.get(province_id) if province_id else None

    def get_city_by_id(self, city_id: str) -> Optional[City]:
        """
        根据ID获取城市信息

        Args:
            city_id: 城市ID

        Returns:
            City: 城市对象，未找到返回None
        """
        return self._city_map.get(city_id)

    def get_city_by_name(self, name: str) -> Optional[City]:
        """
        根据名称获取城市信息（大小写不敏感）

        Args:
            name: 城市名称

        Returns:
            City: 城市对象，未找到返回None
        """
        city_id = self._city_index.get(name.lower())
        return self._city_map.get(city_id) if city_id else None

    def get_county_by_id(self, county_id: str) -> Optional[County]:
        """
        根据ID获取区县信息

        Args:
            county_id: 区县ID

        Returns:
            County: 区县对象，未找到返回None
        """
        return self._county_map.get(county_id)

    def get_county_by_name(self, name: str) -> Optional[County]:
        """
        根据名称获取区县信息（大小写不敏感）

        Args:
            name: 区县名称

        Returns:
            County: 区县对象，未找到返回None
        """
        county_id = self._county_index.get(name.lower())
        return self._county_map.get(county_id) if county_id else None

    def search_by_name(self, name: str) -> List[Union[Province, City, County]]:
        """
        根据名称模糊搜索行政区划

        Args:
            name: 搜索关键词

        Returns:
            List[Union[Province, City, County]]: 匹配的结果列表
        """
        results = []
        lower_name = name.lower()

        # 搜索省份
        for province in self._provinces:
            if (lower_name in province.zh.lower() or
                lower_name in province.en.lower()):
                results.append(province)

        # 搜索城市
        for city in self._cities:
            if (lower_name in city.zh.lower() or
                lower_name in city.en.lower()):
                results.append(city)

        # 搜索区县
        for county in self._counties:
            if (lower_name in county.zh.lower() or
                lower_name in county.en.lower()):
                results.append(county)

        return results

    def list_provinces(self) -> List[Province]:
        """获取所有省份列表"""
        return self._provinces.copy()

    def list_cities(self) -> List[City]:
        """获取所有城市列表"""
        return self._cities.copy()

    def list_counties(self) -> List[County]:
        """获取所有区县列表"""
        return self._counties.copy()

    def get_cities_by_province(self, province_id: str) -> List[City]:
        """
        根据省份ID获取该省的所有城市

        Args:
            province_id: 省份ID

        Returns:
            List[City]: 城市列表
        """
        cities = []
        for city in self._cities:
            if city.id.startswith(province_id[:2]):
                cities.append(city)
        return cities

    def get_counties_by_city(self, city_id: str) -> List[County]:
        """
        根据城市ID获取该市的所有区县

        Args:
            city_id: 城市ID

        Returns:
            List[County]: 区县列表
        """
        counties = []
        for county in self._counties:
            if county.id.startswith(city_id[:4]):
                counties.append(county)
        return counties
