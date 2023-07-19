import typing_extensions

from reearthcmsapi.paths import PathValues
from reearthcmsapi.apis.paths.models_model_id import ModelsModelId
from reearthcmsapi.apis.paths.models_model_id_items import ModelsModelIdItems
from reearthcmsapi.apis.paths.projects_project_id_or_alias_models_model_id_or_key import ProjectsProjectIdOrAliasModelsModelIdOrKey
from reearthcmsapi.apis.paths.projects_project_id_or_alias_models_model_id_or_key_items import ProjectsProjectIdOrAliasModelsModelIdOrKeyItems
from reearthcmsapi.apis.paths.items_item_id import ItemsItemId
from reearthcmsapi.apis.paths.items_item_id_comments import ItemsItemIdComments
from reearthcmsapi.apis.paths.items_item_id_comments_comment_id import ItemsItemIdCommentsCommentId
from reearthcmsapi.apis.paths.projects_project_id_assets import ProjectsProjectIdAssets
from reearthcmsapi.apis.paths.assets_asset_id import AssetsAssetId
from reearthcmsapi.apis.paths.assets_asset_id_comments import AssetsAssetIdComments
from reearthcmsapi.apis.paths.assets_asset_id_comments_comment_id import AssetsAssetIdCommentsCommentId

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.MODELS_MODEL_ID: ModelsModelId,
        PathValues.MODELS_MODEL_ID_ITEMS: ModelsModelIdItems,
        PathValues.PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY: ProjectsProjectIdOrAliasModelsModelIdOrKey,
        PathValues.PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS: ProjectsProjectIdOrAliasModelsModelIdOrKeyItems,
        PathValues.ITEMS_ITEM_ID: ItemsItemId,
        PathValues.ITEMS_ITEM_ID_COMMENTS: ItemsItemIdComments,
        PathValues.ITEMS_ITEM_ID_COMMENTS_COMMENT_ID: ItemsItemIdCommentsCommentId,
        PathValues.PROJECTS_PROJECT_ID_ASSETS: ProjectsProjectIdAssets,
        PathValues.ASSETS_ASSET_ID: AssetsAssetId,
        PathValues.ASSETS_ASSET_ID_COMMENTS: AssetsAssetIdComments,
        PathValues.ASSETS_ASSET_ID_COMMENTS_COMMENT_ID: AssetsAssetIdCommentsCommentId,
    }
)

path_to_api = PathToApi(
    {
        PathValues.MODELS_MODEL_ID: ModelsModelId,
        PathValues.MODELS_MODEL_ID_ITEMS: ModelsModelIdItems,
        PathValues.PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY: ProjectsProjectIdOrAliasModelsModelIdOrKey,
        PathValues.PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS: ProjectsProjectIdOrAliasModelsModelIdOrKeyItems,
        PathValues.ITEMS_ITEM_ID: ItemsItemId,
        PathValues.ITEMS_ITEM_ID_COMMENTS: ItemsItemIdComments,
        PathValues.ITEMS_ITEM_ID_COMMENTS_COMMENT_ID: ItemsItemIdCommentsCommentId,
        PathValues.PROJECTS_PROJECT_ID_ASSETS: ProjectsProjectIdAssets,
        PathValues.ASSETS_ASSET_ID: AssetsAssetId,
        PathValues.ASSETS_ASSET_ID_COMMENTS: AssetsAssetIdComments,
        PathValues.ASSETS_ASSET_ID_COMMENTS_COMMENT_ID: AssetsAssetIdCommentsCommentId,
    }
)
