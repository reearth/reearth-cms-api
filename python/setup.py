# coding: utf-8

"""
    ReEarth-CMS Integration API

    ReEarth-CMS Integration API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""

from setuptools import setup, find_packages  # noqa: H301

NAME = "reearthcmsapi"
VERSION = "0.0.3"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
    "certifi >= 14.5.14",
    "frozendict >= 2.3.4",
    "python-dateutil >= 2.7.0",
    "setuptools >= 21.0.0",
    "typing_extensions >= 4.3.0",
    "urllib3 >= 1.26.7 < 2.1.0",
]

def read_file(filename):
    with open(filename, 'r', encoding='utf-8') as f:
        return f.read()

setup(
    name=NAME,
    version=VERSION,
    description="ReEarth-CMS Integration API",
    long_description=read_file('README.md'),
    long_description_content_type='text/markdown',
    license='MIT',
    license_file='LICENSE',
    author="Eukarya Inc.",
    url="https://github.com/reearth/reearth-cms-api",
    keywords=["OpenAPI", "OpenAPI-Generator", "ReEarth-CMS Integration API"],
    python_requires=">=3.7",
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    setup_requires=['wheel']
)
