<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.groups_api.GroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**group_create**](#group_create) | **post** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups | Create a new group in a project.
[**group_delete**](#group_delete) | **delete** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups/{groupIdOrKey} | Delete a group within a project.
[**group_filter**](#group_filter) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups | Returns a list of groups in a project.
[**group_get**](#group_get) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups/{groupIdOrKey} | Retrieve a group within a project.
[**group_update**](#group_update) | **patch** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/groups/{groupIdOrKey} | Update a group&#x27;s details within a project.

# **group_create**
<a id="group_create"></a>
> Group group_create(workspace_id_or_aliasproject_id_or_aliasany_type)

Create a new group in a project.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import groups_api
from reearthcmsapi.model.group import Group
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
    api_instance = groups_api.GroupsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
    }
    query_params = {
    }
    body = dict(
        description="description_example",
        key="key_example",
        name="name_example",
    )
    try:
        # Create a new group in a project.
        api_response = api_instance.group_create(
            path_params=path_params,
            query_params=query_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_create: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
    }
    query_params = {
        'page': 1,
        'perPage': 50,
        'sort': "createdAt",
        'dir': "desc",
    }
    body = dict(
        description="description_example",
        key="key_example",
        name="name_example",
    )
    try:
        # Create a new group in a project.
        api_response = api_instance.group_create(
            path_params=path_params,
            query_params=query_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_create: %s\n" % e)
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
**name** | str,  | str,  |  | 
**key** | str,  | str,  |  | 
**description** | str,  | str,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
page | PageSchema | | optional
perPage | PerPageSchema | | optional
sort | SortSchema | | optional
dir | DirSchema | | optional


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

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 

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

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
201 | [ApiResponseFor201](#group_create.ApiResponseFor201) | The created group.
400 | [ApiResponseFor400](#group_create.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#group_create.ApiResponseFor401) | Access token is missing or invalid
500 | [ApiResponseFor500](#group_create.ApiResponseFor500) | Internal server error

#### group_create.ApiResponseFor201
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor201ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor201ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Group**](../../models/Group.md) |  | 


#### group_create.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_create.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_create.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **group_delete**
<a id="group_delete"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} group_delete(workspace_id_or_aliasproject_id_or_aliasgroup_id_or_key)

Delete a group within a project.

Deletes a group within a specified project.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import groups_api
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
    api_instance = groups_api.GroupsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'groupIdOrKey': "groupIdOrKey_example",
    }
    try:
        # Delete a group within a project.
        api_response = api_instance.group_delete(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_delete: %s\n" % e)
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
groupIdOrKey | GroupIdOrKeySchema | | 

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

# GroupIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#group_delete.ApiResponseFor200) | Group successfully deleted.
400 | [ApiResponseFor400](#group_delete.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#group_delete.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#group_delete.ApiResponseFor404) | Group not found
500 | [ApiResponseFor500](#group_delete.ApiResponseFor500) | Internal server error

#### group_delete.ApiResponseFor200
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

#### group_delete.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_delete.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_delete.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_delete.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **group_filter**
<a id="group_filter"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} group_filter(workspace_id_or_aliasproject_id_or_alias)

Returns a list of groups in a project.

Returns a list of groups belonging to the specified project.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import groups_api
from reearthcmsapi.model.group import Group
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
    api_instance = groups_api.GroupsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
    }
    query_params = {
    }
    try:
        # Returns a list of groups in a project.
        api_response = api_instance.group_filter(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_filter: %s\n" % e)

    # example passing only optional values
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
    }
    query_params = {
        'page': 1,
        'perPage': 50,
        'sort': "createdAt",
        'dir': "desc",
    }
    try:
        # Returns a list of groups in a project.
        api_response = api_instance.group_filter(
            path_params=path_params,
            query_params=query_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_filter: %s\n" % e)
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
page | PageSchema | | optional
perPage | PerPageSchema | | optional
sort | SortSchema | | optional
dir | DirSchema | | optional


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

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 

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

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#group_filter.ApiResponseFor200) | A JSON object containing a list of groups
400 | [ApiResponseFor400](#group_filter.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#group_filter.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#group_filter.ApiResponseFor404) | Project not found
500 | [ApiResponseFor500](#group_filter.ApiResponseFor500) | Internal server error

#### group_filter.ApiResponseFor200
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
**[groups](#groups)** | list, tuple,  | tuple,  |  | [optional] 
**page** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**perPage** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**totalCount** | decimal.Decimal, int,  | decimal.Decimal,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# groups

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Group**]({{complexTypePrefix}}Group.md) | [**Group**]({{complexTypePrefix}}Group.md) | [**Group**]({{complexTypePrefix}}Group.md) |  | 

#### group_filter.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_filter.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_filter.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_filter.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **group_get**
<a id="group_get"></a>
> Group group_get(workspace_id_or_aliasproject_id_or_aliasgroup_id_or_key)

Retrieve a group within a project.

Returns details of a specific group within the specified project.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import groups_api
from reearthcmsapi.model.group import Group
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
    api_instance = groups_api.GroupsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'groupIdOrKey': "groupIdOrKey_example",
    }
    try:
        # Retrieve a group within a project.
        api_response = api_instance.group_get(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_get: %s\n" % e)
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
groupIdOrKey | GroupIdOrKeySchema | | 

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

# GroupIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#group_get.ApiResponseFor200) | Group details.
400 | [ApiResponseFor400](#group_get.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#group_get.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#group_get.ApiResponseFor404) | Group not found
500 | [ApiResponseFor500](#group_get.ApiResponseFor500) | Internal server error

#### group_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Group**](../../models/Group.md) |  | 


#### group_get.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_get.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_get.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **group_update**
<a id="group_update"></a>
> Group group_update(workspace_id_or_aliasproject_id_or_aliasgroup_id_or_keyany_type)

Update a group's details within a project.

Updates the name, key, or description of a group within a specific project.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import groups_api
from reearthcmsapi.model.group import Group
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
    api_instance = groups_api.GroupsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'groupIdOrKey': "groupIdOrKey_example",
    }
    body = dict(
        description="description_example",
        key="key_example",
        name="name_example",
    )
    try:
        # Update a group's details within a project.
        api_response = api_instance.group_update(
            path_params=path_params,
            body=body,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling GroupsApi->group_update: %s\n" % e)
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
**description** | str,  | str,  |  | [optional] 
**key** | str,  | str,  |  | [optional] 
**name** | str,  | str,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
workspaceIdOrAlias | WorkspaceIdOrAliasSchema | | 
projectIdOrAlias | ProjectIdOrAliasSchema | | 
groupIdOrKey | GroupIdOrKeySchema | | 

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

# GroupIdOrKeySchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#group_update.ApiResponseFor200) | The updated group.
400 | [ApiResponseFor400](#group_update.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#group_update.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#group_update.ApiResponseFor404) | Group not found
500 | [ApiResponseFor500](#group_update.ApiResponseFor500) | Internal server error

#### group_update.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Group**](../../models/Group.md) |  | 


#### group_update.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_update.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_update.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### group_update.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

