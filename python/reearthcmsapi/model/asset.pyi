# coding: utf-8

"""
    ReEarth-CMS Integration API

    ReEarth-CMS Integration API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from reearthcmsapi import schemas  # noqa: F401


class Asset(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "createdAt",
            "id",
            "projectId",
            "url",
            "updatedAt",
        }
        
        class properties:
            id = schemas.StrSchema
            projectId = schemas.StrSchema
            url = schemas.StrSchema
            createdAt = schemas.DateTimeSchema
            updatedAt = schemas.DateTimeSchema
            name = schemas.StrSchema
            contentType = schemas.StrSchema
            
            
            class previewType(
                schemas.EnumBase,
                schemas.StrSchema
            ):
                
                @schemas.classproperty
                def IMAGE(cls):
                    return cls("image")
                
                @schemas.classproperty
                def IMAGE_SVG(cls):
                    return cls("image_svg")
                
                @schemas.classproperty
                def GEO(cls):
                    return cls("geo")
                
                @schemas.classproperty
                def GEO_3D_TILES(cls):
                    return cls("geo_3d_Tiles")
                
                @schemas.classproperty
                def GEO_MVT(cls):
                    return cls("geo_mvt")
                
                @schemas.classproperty
                def MODEL_3D(cls):
                    return cls("model_3d")
                
                @schemas.classproperty
                def UNKNOWN(cls):
                    return cls("unknown")
            totalSize = schemas.NumberSchema
            
            
            class archiveExtractionStatus(
                schemas.EnumBase,
                schemas.StrSchema
            ):
                
                @schemas.classproperty
                def PENDING(cls):
                    return cls("pending")
                
                @schemas.classproperty
                def IN_PROGRESS(cls):
                    return cls("in_progress")
                
                @schemas.classproperty
                def DONE(cls):
                    return cls("done")
                
                @schemas.classproperty
                def FAILED(cls):
                    return cls("failed")
        
            @staticmethod
            def file() -> typing.Type['File']:
                return File
            __annotations__ = {
                "id": id,
                "projectId": projectId,
                "url": url,
                "createdAt": createdAt,
                "updatedAt": updatedAt,
                "name": name,
                "contentType": contentType,
                "previewType": previewType,
                "totalSize": totalSize,
                "archiveExtractionStatus": archiveExtractionStatus,
                "file": file,
            }
    
    createdAt: MetaOapg.properties.createdAt
    id: MetaOapg.properties.id
    projectId: MetaOapg.properties.projectId
    url: MetaOapg.properties.url
    updatedAt: MetaOapg.properties.updatedAt
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["projectId"]) -> MetaOapg.properties.projectId: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["url"]) -> MetaOapg.properties.url: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["createdAt"]) -> MetaOapg.properties.createdAt: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["updatedAt"]) -> MetaOapg.properties.updatedAt: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["contentType"]) -> MetaOapg.properties.contentType: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["previewType"]) -> MetaOapg.properties.previewType: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["totalSize"]) -> MetaOapg.properties.totalSize: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["archiveExtractionStatus"]) -> MetaOapg.properties.archiveExtractionStatus: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file"]) -> 'File': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["id", "projectId", "url", "createdAt", "updatedAt", "name", "contentType", "previewType", "totalSize", "archiveExtractionStatus", "file", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["projectId"]) -> MetaOapg.properties.projectId: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["url"]) -> MetaOapg.properties.url: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["createdAt"]) -> MetaOapg.properties.createdAt: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["updatedAt"]) -> MetaOapg.properties.updatedAt: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> typing.Union[MetaOapg.properties.name, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["contentType"]) -> typing.Union[MetaOapg.properties.contentType, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["previewType"]) -> typing.Union[MetaOapg.properties.previewType, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["totalSize"]) -> typing.Union[MetaOapg.properties.totalSize, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["archiveExtractionStatus"]) -> typing.Union[MetaOapg.properties.archiveExtractionStatus, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file"]) -> typing.Union['File', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id", "projectId", "url", "createdAt", "updatedAt", "name", "contentType", "previewType", "totalSize", "archiveExtractionStatus", "file", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        createdAt: typing.Union[MetaOapg.properties.createdAt, str, datetime, ],
        id: typing.Union[MetaOapg.properties.id, str, ],
        projectId: typing.Union[MetaOapg.properties.projectId, str, ],
        url: typing.Union[MetaOapg.properties.url, str, ],
        updatedAt: typing.Union[MetaOapg.properties.updatedAt, str, datetime, ],
        name: typing.Union[MetaOapg.properties.name, str, schemas.Unset] = schemas.unset,
        contentType: typing.Union[MetaOapg.properties.contentType, str, schemas.Unset] = schemas.unset,
        previewType: typing.Union[MetaOapg.properties.previewType, str, schemas.Unset] = schemas.unset,
        totalSize: typing.Union[MetaOapg.properties.totalSize, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        archiveExtractionStatus: typing.Union[MetaOapg.properties.archiveExtractionStatus, str, schemas.Unset] = schemas.unset,
        file: typing.Union['File', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Asset':
        return super().__new__(
            cls,
            *_args,
            createdAt=createdAt,
            id=id,
            projectId=projectId,
            url=url,
            updatedAt=updatedAt,
            name=name,
            contentType=contentType,
            previewType=previewType,
            totalSize=totalSize,
            archiveExtractionStatus=archiveExtractionStatus,
            file=file,
            _configuration=_configuration,
            **kwargs,
        )

from reearthcmsapi.model.file import File