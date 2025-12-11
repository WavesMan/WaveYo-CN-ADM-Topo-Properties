from setuptools import setup
import os

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="admtp",
    version="1.0.1",
    author="Waves_Man",
    author_email="waveyo@waveyo.cn",
    description="基于民政部公开数据的中国行政区划查询SDK，支持省份、城市、区县的快速查询",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/tree/main/Open-SDK/Python-SDK/admtp",
    packages=["admtp"],
    package_dir={"admtp": "."},
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: Apache Software License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Programming Language :: Python :: 3.13",
        "Programming Language :: Python :: 3.14",
    ],
    python_requires=">=3.7",
    install_requires=[],
    package_data={
        'admtp': [
            'data/*.json',
            'data/cities/*.json',
            'data/counties/*/*.json',
        ],
    },
    include_package_data=True,
    extras_require={
        "dev": [
            "pytest>=6.0",
            "pytest-cov>=2.0",
        ],
    },
)
