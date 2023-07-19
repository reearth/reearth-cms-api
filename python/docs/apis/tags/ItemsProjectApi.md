<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.items_project_api.ItemsProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**item_create_with_project**](#item_create_with_project) | **post** /projects/{projectIdOrAlias}/models/{modelIdOrKey}/items | 
[**item_filter_with_project**](#item_filter_with_project) | **get** /projects/{projectIdOrAlias}/models/{modelIdOrKey}/items | Returns a list of items.

# **item_create_with_project**
<a id="item_create_with_project"></a>
> VersionedItem item_create_with_project(project_id_or_aliasmodel_id_or_keyany_type)



### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_project_api
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
    api_instance = items_project_api.ItemsProjectApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    body = dict(
        fields=[
            Field(
                id="id_example",
                type=ValueType("text"),
                value=None,
                key="key_example",
            )
        ],
    )
    try:
        api_response = api_instance.item_create_with_project(
            path_params=path_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsProjectApi->item_create_with_project: %s\n" % e)
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

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

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
200 | [ApiResponseFor200](#item_create_with_project.ApiResponseFor200) | A JSON array of user names
400 | [ApiResponseFor400](#item_create_with_project.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_create_with_project.ApiResponseFor401) | Access token is missing or invalid

#### item_create_with_project.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**VersionedItem**](../../models/VersionedItem.md) |  | 


#### item_create_with_project.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_create_with_project.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **item_filter_with_project**
<a id="item_filter_with_project"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} item_filter_with_project(project_id_or_aliasmodel_id_or_key)

Returns a list of items.

Returns a list of items with filtering and ordering.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_project_api
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
    api_instance = items_project_api.ItemsProjectApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
    }
    try:
        # Returns a list of items.
        api_response = api_instance.item_filter_with_project(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsProjectApi->item_filter_with_project: %s\n" % e)

    # example passing only optional values
    path_params = {
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
    }
    try:
        # Returns a list of items.
        api_response = api_instance.item_filter_with_project(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsProjectApi->item_filter_with_project: %s\n" % e)
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


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
projectIdOrAlias | ProjectIdOrAliasSchema | | 
modelIdOrKey | ModelIdOrKeySchema | | 

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
200 | [ApiResponseFor200](#item_filter_with_project.ApiResponseFor200) | A JSON array of user names
400 | [ApiResponseFor400](#item_filter_with_project.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#item_filter_with_project.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#item_filter_with_project.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#item_filter_with_project.ApiResponseFor500) | Internal server error

#### item_filter_with_project.ApiResponseFor200
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
**totalCount** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**page** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**perPage** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
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

#### item_filter_with_project.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_with_project.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_with_project.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### item_filter_with_project.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

