<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.items_api.ItemsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**item_create**](#item_create) | **post** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items | create an item
[**item_delete**](#item_delete) | **delete** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId} | delete an item
[**item_filter**](#item_filter) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items | Returns a list of items.
[**item_filter_post**](#item_filter_post) | **post** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/filter | Returns a list of items with complex filtering.
[**item_get**](#item_get) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId} | Returns an item.
[**item_publish**](#item_publish) | **post** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId}/publish | publish item
[**item_update**](#item_update) | **patch** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items/{itemId} | Update an item.
[**items_as_geo_json**](#items_as_geo_json) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items.geojson | Returns a GeoJSON that has a list of items as features.

# **item_create**
<a id="item_create"></a>
> VersionedItem item_create(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyany_type)

create an item

Create an Item.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.field import Field
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    body = dict(
        fields=[
            Field(
                group="group_example",
                id="id_example",
                key="key_example",
                type=ValueType("text"),
                value=None,
            )
        ],
,
    )
    try:
        # create an item
        api_response = api_instance.item_create(
            path_params=path_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_create: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[fields](#fields)** | list, tuple,  | tuple,  |  | [optional] 
**[metadataFields](#metadataFields)** | list, tuple,  | tuple,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# fields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) |  | 

# metadataFields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) |  | 

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_create.ApiResponseFor200) | A JSON object of the item
400 | [ApiResponseFor400](#item_create.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_create.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_create.ApiResponseFor404) | Not found

#### item_create.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**VersionedItem**](../../models/VersionedItem.md) |  | 


#### item_create.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_create.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_create.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_delete**
<a id="item_delete"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} item_delete(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyitem_id)

delete an item

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    try:
        # delete an item
        api_response = api_instance.item_delete(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_delete: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 
itemId | ItemIdSchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ItemIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_delete.ApiResponseFor200) | delete an item
400 | [ApiResponseFor400](#item_delete.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_delete.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_delete.ApiResponseFor404) | Not found

#### item_delete.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**id** | str,  | str,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

#### item_delete.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_delete.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_delete.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_filter**
<a id="item_filter"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} item_filter(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_key)

Returns a list of items.

Returns a list of items with filtering and ordering.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.asset_embedding import AssetEmbedding
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
    }
    try:
        # Returns a list of items.
        api_response = api_instance.item_filter(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_filter: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
        'sort': "createdAt",
        'dir': "desc",
        'page': 1,
        'perPage': 50,
        'ref': "latest",
        'asset': AssetEmbedding("all"),
        'keyword': "keyword_example",
    }
    try:
        # Returns a list of items.
        api_response = api_instance.item_filter(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_filter: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
sort | SortSchema | | optional
dir | DirSchema | | optional
page | PageSchema | | optional
perPage | PerPageSchema | | optional
ref | RefSchema | | optional
asset | AssetSchema | | optional
keyword | KeywordSchema | | optional


# SortSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["createdAt", "updatedAt", ] if omitted the server will use the default value of "createdAt"

# DirSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["asc", "desc", ] if omitted the server will use the default value of "desc"

# PageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 1

# PerPageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 50

# RefSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["latest", "public", ] if omitted the server will use the default value of "latest"

# AssetSchema
Type | Description  | Notes
------------- | ------------- | -------------
[**AssetEmbedding**](../../models/AssetEmbedding.md) |  | 


# KeywordSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_filter.ApiResponseFor200) | A JSON object containing a list of items
400 | [ApiResponseFor400](#item_filter.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_filter.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_filter.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#item_filter.ApiResponseFor500) | Internal server error

#### item_filter.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[items](#items)** | list, tuple,  | tuple,  |  | [optional] 
**page** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**perPage** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**totalCount** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# items

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) | [**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) | [**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) |  | 

#### item_filter.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_filter_post**
<a id="item_filter_post"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} item_filter_post(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyany_type)

Returns a list of items with complex filtering.

Returns a list of items with complex filtering in request body.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.condition import Condition
from reearthcmsapi.model.asset_embedding import AssetEmbedding
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
    }
    body = dict(
        filter=Condition(
            _and=[
                Condition()
            ],
            basic=dict(
                field_id=FieldSelector(
                    field_id="field_id_example",
                    type="id",
                ),
                operator="equals",
                value=None,
            ),
            _bool=dict(
                field_id=FieldSelector(),
                operator="equals",
                value=True,
            ),
            multiple=dict(
                field_id=FieldSelector(),
                operator="includesAny",
                value=[
                    None
                ],
            ),
            nullable=dict(
                field_id=FieldSelector(),
                operator="empty",
            ),
            number=dict(
                field_id=FieldSelector(),
                operator="greaterThan",
                value=3.14,
            ),
,
            string=dict(
                field_id=FieldSelector(),
                operator="contains",
                value="value_example",
            ),
            time=dict(
                field_id=FieldSelector(),
                operator="before",
                value="1970-01-01T00:00:00.00Z",
            ),
        ),
    )
    try:
        # Returns a list of items with complex filtering.
        api_response = api_instance.item_filter_post(
            path_params=path_params,
            query_params=query_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_filter_post: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
        'sort': "createdAt",
        'dir': "desc",
        'page': 1,
        'perPage': 50,
        'ref': "latest",
        'asset': AssetEmbedding("all"),
        'keyword': "keyword_example",
    }
    body = dict(
        filter=Condition(
            _and=[
                Condition()
            ],
            basic=dict(
                field_id=FieldSelector(
                    field_id="field_id_example",
                    type="id",
                ),
                operator="equals",
                value=None,
            ),
            _bool=dict(
                field_id=FieldSelector(),
                operator="equals",
                value=True,
            ),
            multiple=dict(
                field_id=FieldSelector(),
                operator="includesAny",
                value=[
                    None
                ],
            ),
            nullable=dict(
                field_id=FieldSelector(),
                operator="empty",
            ),
            number=dict(
                field_id=FieldSelector(),
                operator="greaterThan",
                value=3.14,
            ),
,
            string=dict(
                field_id=FieldSelector(),
                operator="contains",
                value="value_example",
            ),
            time=dict(
                field_id=FieldSelector(),
                operator="before",
                value="1970-01-01T00:00:00.00Z",
            ),
        ),
    )
    try:
        # Returns a list of items with complex filtering.
        api_response = api_instance.item_filter_post(
            path_params=path_params,
            query_params=query_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_filter_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**filter** | [**Condition**]({{complexTypePrefix}}Condition.md) | [**Condition**]({{complexTypePrefix}}Condition.md) |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
sort | SortSchema | | optional
dir | DirSchema | | optional
page | PageSchema | | optional
perPage | PerPageSchema | | optional
ref | RefSchema | | optional
asset | AssetSchema | | optional
keyword | KeywordSchema | | optional


# SortSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["createdAt", "updatedAt", ] if omitted the server will use the default value of "createdAt"

# DirSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["asc", "desc", ] if omitted the server will use the default value of "desc"

# PageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 1

# PerPageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 50

# RefSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["latest", "public", ] if omitted the server will use the default value of "latest"

# AssetSchema
Type | Description  | Notes
------------- | ------------- | -------------
[**AssetEmbedding**](../../models/AssetEmbedding.md) |  | 


# KeywordSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_filter_post.ApiResponseFor200) | A JSON object containing a list of items
400 | [ApiResponseFor400](#item_filter_post.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_filter_post.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_filter_post.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#item_filter_post.ApiResponseFor500) | Internal server error

#### item_filter_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[items](#items)** | list, tuple,  | tuple,  |  | [optional] 
**page** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**perPage** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**totalCount** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# items

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) | [**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) | [**VersionedItem**]({{complexTypePrefix}}VersionedItem.md) |  | 

#### item_filter_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_post.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_post.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_get**
<a id="item_get"></a>
> VersionedItem item_get(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyitem_id)

Returns an item.

Returns an item.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.asset_embedding import AssetEmbedding
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    query_params = {
    }
    try:
        # Returns an item.
        api_response = api_instance.item_get(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_get: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    query_params = {
        'ref': "latest",
        'asset': AssetEmbedding("all"),
    }
    try:
        # Returns an item.
        api_response = api_instance.item_get(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_get: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
ref | RefSchema | | optional
asset | AssetSchema | | optional


# RefSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["latest", "public", ] if omitted the server will use the default value of "latest"

# AssetSchema
Type | Description  | Notes
------------- | ------------- | -------------
[**AssetEmbedding**](../../models/AssetEmbedding.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 
itemId | ItemIdSchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ItemIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_get.ApiResponseFor200) | An item
400 | [ApiResponseFor400](#item_get.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_get.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_get.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#item_get.ApiResponseFor500) | Internal server error

#### item_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**VersionedItem**](../../models/VersionedItem.md) |  | 


#### item_get.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_get.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_get.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_publish**
<a id="item_publish"></a>
> VersionedItem item_publish(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyitem_id)

publish item

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.asset_embedding import AssetEmbedding
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    query_params = {
    }
    try:
        # publish item
        api_response = api_instance.item_publish(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_publish: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    query_params = {
        'asset': AssetEmbedding("all"),
    }
    try:
        # publish item
        api_response = api_instance.item_publish(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_publish: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
asset | AssetSchema | | optional


# AssetSchema
Type | Description  | Notes
------------- | ------------- | -------------
[**AssetEmbedding**](../../models/AssetEmbedding.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 
itemId | ItemIdSchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ItemIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_publish.ApiResponseFor200) | the published item
400 | [ApiResponseFor400](#item_publish.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_publish.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_publish.ApiResponseFor404) | Not found

#### item_publish.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**VersionedItem**](../../models/VersionedItem.md) |  | 


#### item_publish.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_publish.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_publish.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_update**
<a id="item_update"></a>
> VersionedItem item_update(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_keyitem_idany_type)

Update an item.

Update an item.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from reearthcmsapi.model.versioned_item import VersionedItem
from reearthcmsapi.model.field import Field
from reearthcmsapi.model.asset_embedding import AssetEmbedding
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
        'itemId': "itemId_example",
    }
    body = dict(
        asset=AssetEmbedding("all"),
        fields=[
            Field(
                group="group_example",
                id="id_example",
                key="key_example",
                type=ValueType("text"),
                value=None,
            )
        ],
,
    )
    try:
        # Update an item.
        api_response = api_instance.item_update(
            path_params=path_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->item_update: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**asset** | [**AssetEmbedding**]({{complexTypePrefix}}AssetEmbedding.md) | [**AssetEmbedding**]({{complexTypePrefix}}AssetEmbedding.md) |  | [optional] 
**[fields](#fields)** | list, tuple,  | tuple,  |  | [optional] 
**[metadataFields](#metadataFields)** | list, tuple,  | tuple,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# fields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) |  | 

# metadataFields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) | [**Field**]({{complexTypePrefix}}Field.md) |  | 

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 
itemId | ItemIdSchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ItemIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#item_update.ApiResponseFor200) | An item
400 | [ApiResponseFor400](#item_update.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_update.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_update.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#item_update.ApiResponseFor500) | Internal server error

#### item_update.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**VersionedItem**](../../models/VersionedItem.md) |  | 


#### item_update.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_update.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_update.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_update.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **items_as_geo_json**
<a id="items_as_geo_json"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} items_as_geo_json(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_key)

Returns a GeoJSON that has a list of items as features.

Returns a GeoJSON that has a list of items as features.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_api
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = reearthcmsapi.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = reearthcmsapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)
# Enter a context with an instance of the API client
with reearthcmsapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = items_api.ItemsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
    }
    try:
        # Returns a GeoJSON that has a list of items as features.
        api_response = api_instance.items_as_geo_json(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->items_as_geo_json: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
        'page': 1,
        'perPage': 50,
        'ref': "latest",
    }
    try:
        # Returns a GeoJSON that has a list of items as features.
        api_response = api_instance.items_as_geo_json(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsApi->items_as_geo_json: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/octet-stream', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
page | PageSchema | | optional
perPage | PerPageSchema | | optional
ref | RefSchema | | optional


# PageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 1

# PerPageSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | if omitted the server will use the default value of 50

# RefSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | must be one of ["latest", "public", ] if omitted the server will use the default value of "latest"

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

# WorkspaceIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ProjectIdOrAliasSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# ModelIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#items_as_geo_json.ApiResponseFor200) | A GeoJSON object
400 | [ApiResponseFor400](#items_as_geo_json.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#items_as_geo_json.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#items_as_geo_json.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#items_as_geo_json.ApiResponseFor500) | Internal server error

#### items_as_geo_json.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationOctetStream, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationOctetStream

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

#### items_as_geo_json.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_geo_json.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_geo_json.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_geo_json.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

