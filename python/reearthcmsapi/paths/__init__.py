# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from reearthcmsapi.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    WORKSPACE_ID_OR_ALIAS_PROJECTS = "/{workspaceIdOrAlias}/projects"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_ASSET_ID = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_ASSET_ID_COMMENTS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}/comments"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_ASSET_ID_COMMENTS_COMMENT_ID = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}/comments/{commentId}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_ASSET_ID_PUBLISH = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}/publish"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_ASSET_ID_UNPUBLISH = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{assetId}/unpublish"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_UUID1_UUID2_FILENAME = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/{uuid1}/{uuid2}/{filename}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_ASSETS_UPLOADS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/assets/uploads"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_GROUPS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_GROUPS_GROUP_ID_OR_KEY = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups/{groupIdOrKey}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_COPY = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/copy"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_IMPORT = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/import"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_CSV = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items.csv"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_GEOJSON = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items.geojson"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_ITEM_ID = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_ITEM_ID_COMMENTS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}/comments"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_ITEM_ID_COMMENTS_COMMENT_ID = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}/comments/{commentId}"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_ITEM_ID_PUBLISH = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}/publish"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_ITEMS_FILTER = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/filter"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_METADATA_SCHEMA_JSON = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/metadata_schema.json"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_MODELS_MODEL_ID_OR_KEY_SCHEMA_JSON = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/schema.json"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_SCHEMATA_SCHEMA_ID_FIELDS = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/schemata/{schemaId}/fields"
    WORKSPACE_ID_OR_ALIAS_PROJECTS_PROJECT_ID_OR_ALIAS_SCHEMATA_SCHEMA_ID_FIELDS_FIELD_ID_OR_KEY = "/{workspaceIdOrAlias}/projects/{projectIdOrAlias}/schemata/{schemaId}/fields/{fieldIdOrKey}"
