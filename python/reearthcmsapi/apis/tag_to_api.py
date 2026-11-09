import typing_extensions

from reearthcmsapi.apis.tags import TagValues
from reearthcmsapi.apis.tags.assets_api import AssetsApi
from reearthcmsapi.apis.tags.assets_comments_api import AssetsCommentsApi
from reearthcmsapi.apis.tags.assets_project_api import AssetsProjectApi
from reearthcmsapi.apis.tags.items_project_api import ItemsProjectApi
from reearthcmsapi.apis.tags.csv_api import CSVApi
from reearthcmsapi.apis.tags.fields_api import FieldsApi
from reearthcmsapi.apis.tags.items_api import ItemsApi
from reearthcmsapi.apis.tags.geo_json_api import GeoJSONApi
from reearthcmsapi.apis.tags.groups_api import GroupsApi
from reearthcmsapi.apis.tags.items_comments_api import ItemsCommentsApi
from reearthcmsapi.apis.tags.metadata_schema_api import MetadataSchemaApi
from reearthcmsapi.apis.tags.json_api import JSONApi
from reearthcmsapi.apis.tags.schema_api import SchemaApi
from reearthcmsapi.apis.tags.models_api import ModelsApi
from reearthcmsapi.apis.tags.projects_api import ProjectsApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.ASSETS: AssetsApi,
        TagValues.ASSETS_COMMENTS: AssetsCommentsApi,
        TagValues.ASSETS_PROJECT: AssetsProjectApi,
        TagValues.ITEMS_PROJECT: ItemsProjectApi,
        TagValues.CSV: CSVApi,
        TagValues.FIELDS: FieldsApi,
        TagValues.ITEMS: ItemsApi,
        TagValues.GEO_JSON: GeoJSONApi,
        TagValues.GROUPS: GroupsApi,
        TagValues.ITEMS_COMMENTS: ItemsCommentsApi,
        TagValues.METADATA_SCHEMA: MetadataSchemaApi,
        TagValues.JSON: JSONApi,
        TagValues.SCHEMA: SchemaApi,
        TagValues.MODELS: ModelsApi,
        TagValues.PROJECTS: ProjectsApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.ASSETS: AssetsApi,
        TagValues.ASSETS_COMMENTS: AssetsCommentsApi,
        TagValues.ASSETS_PROJECT: AssetsProjectApi,
        TagValues.ITEMS_PROJECT: ItemsProjectApi,
        TagValues.CSV: CSVApi,
        TagValues.FIELDS: FieldsApi,
        TagValues.ITEMS: ItemsApi,
        TagValues.GEO_JSON: GeoJSONApi,
        TagValues.GROUPS: GroupsApi,
        TagValues.ITEMS_COMMENTS: ItemsCommentsApi,
        TagValues.METADATA_SCHEMA: MetadataSchemaApi,
        TagValues.JSON: JSONApi,
        TagValues.SCHEMA: SchemaApi,
        TagValues.MODELS: ModelsApi,
        TagValues.PROJECTS: ProjectsApi,
    }
)
