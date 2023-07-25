import typing_extensions

from reearthcmsapi.apis.tags import TagValues
from reearthcmsapi.apis.tags.assets_api import AssetsApi
from reearthcmsapi.apis.tags.assets_comments_api import AssetsCommentsApi
from reearthcmsapi.apis.tags.assets_project_api import AssetsProjectApi
from reearthcmsapi.apis.tags.items_api import ItemsApi
from reearthcmsapi.apis.tags.items_comments_api import ItemsCommentsApi
from reearthcmsapi.apis.tags.items_project_api import ItemsProjectApi
from reearthcmsapi.apis.tags.models_api import ModelsApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.ASSETS: AssetsApi,
        TagValues.ASSETS_COMMENTS: AssetsCommentsApi,
        TagValues.ASSETS_PROJECT: AssetsProjectApi,
        TagValues.ITEMS: ItemsApi,
        TagValues.ITEMS_COMMENTS: ItemsCommentsApi,
        TagValues.ITEMS_PROJECT: ItemsProjectApi,
        TagValues.MODELS: ModelsApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.ASSETS: AssetsApi,
        TagValues.ASSETS_COMMENTS: AssetsCommentsApi,
        TagValues.ASSETS_PROJECT: AssetsProjectApi,
        TagValues.ITEMS: ItemsApi,
        TagValues.ITEMS_COMMENTS: ItemsCommentsApi,
        TagValues.ITEMS_PROJECT: ItemsProjectApi,
        TagValues.MODELS: ModelsApi,
    }
)
