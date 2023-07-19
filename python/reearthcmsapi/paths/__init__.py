# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from reearthcmsapi.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    MODELS_MODEL_ID = "/models/{modelId}"
    MODELS_MODEL_ID_ITEMS = "/models/{modelId}/items"
    PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY = "/projects/{projectIdOrAlias}/models/{modelIdOrKey}"
    PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS = "/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items"
    ITEMS_ITEM_ID = "/items/{itemId}"
    ITEMS_ITEM_ID_COMMENTS = "/items/{itemId}/comments"
    ITEMS_ITEM_ID_COMMENTS_COMMENT_ID = "/items/{itemId}/comments/{commentId}"
    PROJECTS_PROJECT_ID_ASSETS = "/projects/{projectId}/assets"
    ASSETS_ASSET_ID = "/assets/{assetId}"
    ASSETS_ASSET_ID_COMMENTS = "/assets/{assetId}/comments"
    ASSETS_ASSET_ID_COMMENTS_COMMENT_ID = "/assets/{assetId}/comments/{commentId}"
