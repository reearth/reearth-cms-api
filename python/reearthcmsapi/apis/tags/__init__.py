# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from reearthcmsapi.apis.tag_to_api import tag_to_api

import enum


class TagValues(str, enum.Enum):
    ASSETS = "Assets"
    ASSETS_COMMENTS = "Assets comments"
    ASSETS_PROJECT = "Assets project"
    ITEMS_PROJECT = "Items project"
    CSV = "CSV"
    FIELDS = "Fields"
    ITEMS = "Items"
    GEO_JSON = "GeoJSON"
    GROUPS = "Groups"
    ITEMS_COMMENTS = "Items comments"
    METADATA_SCHEMA = "MetadataSchema"
    JSON = "JSON"
    SCHEMA = "Schema"
    MODELS = "Models"
    PROJECTS = "Projects"
