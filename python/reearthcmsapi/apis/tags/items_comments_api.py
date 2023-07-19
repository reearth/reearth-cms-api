# coding: utf-8

"""
    ReEarth-CMS Integration API

    ReEarth-CMS Integration API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""

from reearthcmsapi.paths.items_item_id_comments.post import ItemCommentCreate
from reearthcmsapi.paths.items_item_id_comments_comment_id.delete import ItemCommentDelete
from reearthcmsapi.paths.items_item_id_comments.get import ItemCommentList
from reearthcmsapi.paths.items_item_id_comments_comment_id.patch import ItemCommentUpdate


class ItemsCommentsApi(
    ItemCommentCreate,
    ItemCommentDelete,
    ItemCommentList,
    ItemCommentUpdate,
):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """
    pass
