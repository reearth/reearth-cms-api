# reearthcmsapi.model.versioned_item.VersionedItem

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**createdAt** | str, datetime,  | str,  |  | [optional] value must conform to RFC-3339 date-time
**[fields](#fields)** | list, tuple,  | tuple,  |  | [optional] 
**id** | str,  | str,  |  | [optional] 
**[metadataFields](#metadataFields)** | list, tuple,  | tuple,  |  | [optional] 
**modelId** | str,  | str,  |  | [optional] 
**[parents](#parents)** | list, tuple,  | tuple,  |  | [optional] 
**[referencedItems](#referencedItems)** | list, tuple,  | tuple,  |  | [optional] 
**[refs](#refs)** | list, tuple,  | tuple,  |  | [optional] 
**updatedAt** | str, datetime,  | str,  |  | [optional] value must conform to RFC-3339 date-time
**version** | str, uuid.UUID,  | str,  |  | [optional] value must be a uuid
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# fields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**](Field.md) | [**Field**](Field.md) | [**Field**](Field.md) |  | 

# metadataFields

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Field**](Field.md) | [**Field**](Field.md) | [**Field**](Field.md) |  | 

# parents

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str, uuid.UUID,  | str,  |  | value must be a uuid

# referencedItems

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**VersionedItem**](VersionedItem.md) | [**VersionedItem**](VersionedItem.md) | [**VersionedItem**](VersionedItem.md) |  | 

# refs

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str,  | str,  |  | 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

