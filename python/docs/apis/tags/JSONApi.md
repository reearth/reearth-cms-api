<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.json_api.JSONApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metadata_schema_by_model_as_json**](#metadata_schema_by_model_as_json) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/metadata_schema.json | Returns a metadata schema as json by project and model ID
[**schema_by_model_as_json**](#schema_by_model_as_json) | **get** /{workspaceIdOrAlias}/projects/{projectIdOrAlias}/models/{modelIdOrKey}/schema.json | Returns a schema as json by project and model ID

# **metadata_schema_by_model_as_json**
<a id="metadata_schema_by_model_as_json"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} metadata_schema_by_model_as_json(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_key)

Returns a metadata schema as json by project and model ID

Returns a metadata schema as json by project and model ID

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import json_api
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
    api_instance = json_api.JSONApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    try:
        # Returns a metadata schema as json by project and model ID
        api_response = api_instance.metadata_schema_by_model_as_json(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling JSONApi->metadata_schema_by_model_as_json: %s\n" % e)
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
200 | [ApiResponseFor200](#metadata_schema_by_model_as_json.ApiResponseFor200) | A JSON object
400 | [ApiResponseFor400](#metadata_schema_by_model_as_json.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#metadata_schema_by_model_as_json.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#metadata_schema_by_model_as_json.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#metadata_schema_by_model_as_json.ApiResponseFor500) | Internal server error

#### metadata_schema_by_model_as_json.ApiResponseFor200
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

#### metadata_schema_by_model_as_json.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### metadata_schema_by_model_as_json.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### metadata_schema_by_model_as_json.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### metadata_schema_by_model_as_json.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **schema_by_model_as_json**
<a id="schema_by_model_as_json"></a>
> {str: (bool, date, datetime, dict, float, int, list, str, none_type)} schema_by_model_as_json(workspace_id_or_aliasproject_id_or_aliasmodel_id_or_key)

Returns a schema as json by project and model ID

Returns a schema as json by project and model ID

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import json_api
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
    api_instance = json_api.JSONApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'workspaceIdOrAlias': "workspaceIdOrAlias_example",
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    try:
        # Returns a schema as json by project and model ID
        api_response = api_instance.schema_by_model_as_json(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling JSONApi->schema_by_model_as_json: %s\n" % e)
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
200 | [ApiResponseFor200](#schema_by_model_as_json.ApiResponseFor200) | A JSON object
400 | [ApiResponseFor400](#schema_by_model_as_json.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#schema_by_model_as_json.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#schema_by_model_as_json.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#schema_by_model_as_json.ApiResponseFor500) | Internal server error

#### schema_by_model_as_json.ApiResponseFor200
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

#### schema_by_model_as_json.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### schema_by_model_as_json.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### schema_by_model_as_json.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### schema_by_model_as_json.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

