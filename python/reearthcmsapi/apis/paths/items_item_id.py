from reearthcmsapi.paths.items_item_id.get import ApiForget
from reearthcmsapi.paths.items_item_id.delete import ApiFordelete
from reearthcmsapi.paths.items_item_id.patch import ApiForpatch


class ItemsItemId(
    ApiForget,
    ApiFordelete,
    ApiForpatch,
):
    pass
