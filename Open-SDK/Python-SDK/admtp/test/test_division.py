import pytest
from admtp import Division
from admtp.models import Province, City, County

@pytest.fixture
def division():
    return Division()

def test_get_province_by_id(division):
    """测试根据ID获取省份"""
    province = division.get_province_by_id("110000")
    assert province is not None
    assert province.id == "110000"
    assert province.zh == "北京"
    assert province.en == "Beijing"

def test_get_province_by_name(division):
    """测试根据名称获取省份"""
    province = division.get_province_by_name("北京")
    assert province is not None
    assert province.id == "110000"

    # 测试大小写不敏感
    province = division.get_province_by_name("beijing")
    assert province is not None
    assert province.id == "110000"

def test_get_city_by_id(division):
    """测试根据ID获取城市"""
    city = division.get_city_by_id("130100")
    assert city is not None
    assert city.id == "130100"
    assert city.zh == "石家庄"
    assert city.en == "Shijiazhuang"

def test_get_county_by_id(division):
    """测试根据ID获取区县"""
    county = division.get_county_by_id("110101")
    assert county is not None
    assert county.id == "110101"
    assert county.zh == "东城区"
    assert county.en == "Dongchengqu"

def test_search_by_name(division):
    """测试模糊搜索"""
    results = division.search_by_name("北京")
    assert len(results) > 0

    # 应该包含北京市和东城区等
    province_found = any(isinstance(item, Province) and item.zh == "北京" for item in results)
    assert province_found

def test_list_provinces(division):
    """测试获取所有省份"""
    provinces = division.list_provinces()
    assert len(provinces) > 0
    assert isinstance(provinces[0], Province)

def test_get_cities_by_province(division):
    """测试获取省份下的城市"""
    cities = division.get_cities_by_province("130000")  # 河北省
    assert len(cities) > 0
    assert isinstance(cities[0], City)

def test_get_counties_by_city(division):
    """测试获取城市下的区县"""
    counties = division.get_counties_by_city("110100")  # 北京市
    assert len(counties) > 0
    assert isinstance(counties[0], County)
