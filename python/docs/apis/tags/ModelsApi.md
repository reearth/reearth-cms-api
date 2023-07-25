<a id="__pageTop"></a>
# reearthcmsapi.apis.tags.models_api.ModelsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**model_get**](#model_get) | **get** /models/{modelId} | Returns a model.
[**model_get_with_project**](#model_get_with_project) | **get** /projects/{projectIdOrAlias}/models/{modelIdOrKey} | Returns a model.

# **model_get**
<a id="model_get"></a>
> Model model_get(model_id)

Returns a model.

Returns a model.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import models_api
from reearthcmsapi.model.model import Model
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
    api_instance = models_api.ModelsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'modelId': "modelId_example",
    }
    try:
        # Returns a model.
        api_response = api_instance.model_get(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ModelsApi->model_get: %s\n" % e)
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
modelId | ModelIdSchema | | 

# ModelIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#model_get.ApiResponseFor200) | A JSON array of user names
400 | [ApiResponseFor400](#model_get.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#model_get.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#model_get.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#model_get.ApiResponseFor500) | Internal server error

#### model_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Model**](../../models/Model.md) |  | 


#### model_get.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **model_get_with_project**
<a id="model_get_with_project"></a>
> Model model_get_with_project(project_id_or_aliasmodel_id_or_key)

Returns a model.

Returns a model.

### Example

* Bearer Authentication (bearerAuth):
```python
import reearthcmsapi
from reearthcmsapi.apis.tags import models_api
from reearthcmsapi.model.model import Model
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
    api_instance = models_api.ModelsApi(api_client)

    # example passing only required values which don't have defaults set
    path_params = {
        'projectIdOrAlias': "projectIdOrAlias_example",
        'modelIdOrKey': "modelIdOrKey_example",
    }
    try:
        # Returns a model.
        api_response = api_instance.model_get_with_project(
            path_params=path_params,
        )
        pprint(api_response)
    except reearthcmsapi.ApiException as e:
        print("Exception when calling ModelsApi->model_get_with_project: %s\n" % e)
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
200 | [ApiResponseFor200](#model_get_with_project.ApiResponseFor200) | A JSON array of user names
400 | [ApiResponseFor400](#model_get_with_project.ApiResponseFor400) | Invalid request parameter value
401 | [ApiResponseFor401](#model_get_with_project.ApiResponseFor401) | Access token is missing or invalid
404 | [ApiResponseFor404](#model_get_with_project.ApiResponseFor404) | Not found
500 | [ApiResponseFor500](#model_get_with_project.ApiResponseFor500) | Internal server error

#### model_get_with_project.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Model**](../../models/Model.md) |  | 


#### model_get_with_project.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get_with_project.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get_with_project.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### model_get_with_project.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

