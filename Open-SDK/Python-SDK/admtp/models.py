"""
数据模型定义
"""

class DivisionBase:
    """行政区划基础类"""
    
    def __init__(self, id, zh, en):
        self.id = id
        self.zh = zh
        self.en = en
    
    def __repr__(self):
        return f"<{self.__class__.__name__} id={self.id} zh='{self.zh}' en='{self.en}'>"
    
    def to_dict(self):
        """转换为字典格式"""
        return {
            "id": self.id,
            "zh": self.zh,
            "en": self.en
        }

class Province(DivisionBase):
    """省份/直辖市/自治区/特别行政区"""
    pass

class City(DivisionBase):
    """地级市/自治州/盟"""
    pass

class County(DivisionBase):
    """区/县/县级市/旗/自治旗/林区/特区"""
    pass
