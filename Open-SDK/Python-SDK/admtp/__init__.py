"""
中国行政区划数据 Python SDK
提供省份、城市、区县的查询功能
"""

from .core import Division
from .models import Province, City, County

__version__ = "1.0.0"
__author__ = "Waves_Man"

__all__ = ["Division", "Province", "City", "County"]
