# reearthcmsapi.model.condition.Condition

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[and](#and)** | list, tuple,  | tuple,  |  | [optional] 
**[basic](#basic)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[bool](#bool)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[multiple](#multiple)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[nullable](#nullable)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[number](#number)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[or](#or)** | list, tuple,  | tuple,  |  | [optional] 
**[string](#string)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**[time](#time)** | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# and

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Condition**](Condition.md) | [**Condition**](Condition.md) | [**Condition**](Condition.md) |  | 

# basic

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | [optional] 
**operator** | str,  | str,  |  | [optional] must be one of ["equals", "notEquals", ] 
**value** | dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader,  | frozendict.frozendict, str, decimal.Decimal, BoolClass, NoneClass, tuple, bytes, FileIO |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# bool

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**value** | bool,  | BoolClass,  |  | 
**operator** | str,  | str,  |  | must be one of ["equals", "notEquals", ] 
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# multiple

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[value](#value)** | list, tuple,  | tuple,  |  | 
**operator** | str,  | str,  |  | must be one of ["includesAny", "notIncludesAny", "includesAll", "notIncludesAll", ] 
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# value

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader,  | frozendict.frozendict, str, decimal.Decimal, BoolClass, NoneClass, tuple, bytes, FileIO |  | 

# nullable

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | [optional] 
**operator** | str,  | str,  |  | [optional] must be one of ["empty", "notEmpty", ] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# number

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**value** | decimal.Decimal, int, float,  | decimal.Decimal,  |  | 
**operator** | str,  | str,  |  | must be one of ["greaterThan", "lessThan", "greaterThanOrEqualTo", "lessThanOrEqualTo", ] 
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# or

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**Condition**](Condition.md) | [**Condition**](Condition.md) | [**Condition**](Condition.md) |  | 

# string

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**value** | str,  | str,  |  | 
**operator** | str,  | str,  |  | must be one of ["contains", "notContains", "startsWith", "endsWith", "notStartsWith", "notEndsWith", ] 
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# time

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**value** | str, datetime,  | str,  |  | value must conform to RFC-3339 date-time
**operator** | str,  | str,  |  | must be one of ["before", "after", "beforeOrOn", "afterOrOn", "ofThisWeek", "ofThisMonth", "ofThisYear", ] 
**fieldId** | [**FieldSelector**](FieldSelector.md) | [**FieldSelector**](FieldSelector.md) |  | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

