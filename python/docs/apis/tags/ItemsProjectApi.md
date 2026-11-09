<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.items_project_api.ItemsProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**items_as_csv**](#items_as_csv) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items.csv | Returns a CSV that has a list of items as features.

# **items_as_csv**
<a id="items_as_csv"></a>
> file_type items_as_csv(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_key)

Returns a CSV that has a list of items as features.

Returns a CSV that has a list of items as features.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import items_project_api
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
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    query_params = {
    }
    try:
        # Returns a CSV that has a list of items as features.
        api_response = api_instance.items_as_csv(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsProjectApi->items_as_csv: %s\n" % e)

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
        # Returns a CSV that has a list of items as features.
        api_response = api_instance.items_as_csv(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ItemsProjectApi->items_as_csv: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('text/csv', ) | Tells the server the content type(s) that are accepted by the client
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
200 | [ApiResponseFor200](#items_as_csv.ApiResponseFor200) | A string in CSV format
400 | [ApiResponseFor400](#items_as_csv.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#items_as_csv.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#items_as_csv.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#items_as_csv.ApiResponseFor500) | Internal server error

#### items_as_csv.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyTextCsv, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyTextCsv

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
bytes, io.FileIO, io.BufferedReader,  | bytes, FileIO,  |  | 

#### items_as_csv.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_csv.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_csv.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### items_as_csv.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

