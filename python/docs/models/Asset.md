# reearthcmsapi.model.asset.Asset

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**createdAt** | str, datetime,  | str,  |  | value must conform to RFC-3339 date-time
**id** | str,  | str,  |  | 
**projectId** | str,  | str,  |  | 
**url** | str,  | str,  |  | 
**updatedAt** | str, datetime,  | str,  |  | value must conform to RFC-3339 date-time
**name** | str,  | str,  |  | [optional] 
**contentType** | str,  | str,  |  | [optional] 
**previewType** | str,  | str,  |  | [optional] must be one of ["image", "image_svg", "geo", "geo_3d_Tiles", "geo_mvt", "model_3d", "unknown", ] 
**totalSize** | decimal.Decimal, int, float,  | decimal.Decimal,  |  | [optional] 
**archiveExtractionStatus** | str,  | str,  |  | [optional] must be one of ["pending", "in_progress", "done", "failed", ] 
**file** | [**File**](File.md) | [**File**](File.md) |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

