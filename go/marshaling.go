package cms

import (
	"reflect"
	"slices"
	"strings"

	"github.com/samber/lo"
)

var marshalCMSType = reflect.TypeOf((*MarshalCMS)(nil)).Elem()

func (d *Item) Unmarshal(i any) {
	if i == nil {
		return
	}

	v := reflect.ValueOf(i)
	if v.IsNil() {
		return
	}

	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, opts, _ := strings.Cut(tag, ",")
		if key == "" || key == "-" {
			continue
		}

		isMetadata := strings.HasSuffix(opts, ",metadata")
		vf := v.FieldByName(f.Name)

		if key == "id" {
			if f.Type.Kind() == reflect.String {
				vf.SetString(d.ID)
			}
			continue
		}

		var itf *Field
		if isMetadata {
			itf = d.MetadataFieldByKey(key)
		} else {
			itf = d.FieldByKey(key)
		}

		if itf != nil && itf.Type == "group" {
			groupIDs := itf.GetValue().Strings()
			if len(groupIDs) == 0 {
				if groupID := itf.GetValue().String(); groupID != nil {
					groupIDs = []string{*groupID}
				}
			}

			groups := make([]*Item, 0, len(groupIDs))
			for _, g := range groupIDs {
				group := d.Group(g)
				groups = append(groups, group)
			}

			if len(groups) == 0 {
				continue
			}

			if f.Type.Kind() == reflect.Slice &&
				(f.Type.Elem().Kind() == reflect.Struct ||
					f.Type.Elem().Kind() == reflect.Ptr &&
						f.Type.Elem().Elem().Kind() == reflect.Struct) {
				s := reflect.MakeSlice(f.Type, 0, len(groups))
				isPointer := f.Type.Elem().Kind() == reflect.Ptr

				for _, g := range groups {
					var rv reflect.Value
					if isPointer {
						rv = reflect.New(f.Type.Elem().Elem())
					} else {
						rv = reflect.New(f.Type.Elem())
					}

					i := rv.Interface()
					g.Unmarshal(i)

					if isPointer {
						s = reflect.Append(s, rv)
					} else {
						s = reflect.Append(s, rv.Elem())
					}

				}

				vf.Set(s)
			} else if f.Type.Kind() == reflect.Struct {
				groups[0].Unmarshal(vf.Addr().Interface())
			} else if f.Type.Kind() == reflect.Pointer && f.Type.Elem().Kind() == reflect.Struct {
				groups[0].Unmarshal(vf.Interface())
			}

			continue
		}

		if itf == nil || itf.Value == nil || !vf.CanSet() {
			continue
		}

		// tag
		if ok := assignIf(vf, func() (Tag, bool) {
			t := TagFrom(itf.Value)
			if t == nil {
				return Tag{}, false
			}
			return *t, true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]Tag, bool) {
			return TagsFrom(itf.Value), true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]*Tag, bool) {
			return lo.ToSlicePtr(TagsFrom(itf.Value)), true
		}); ok {
			continue
		}

		// asset
		if ok := assignIf(vf, func() (Asset, bool) {
			a := AssetFrom(itf.Value)
			if a == nil {
				return Asset{}, false
			}
			return *a, true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]Asset, bool) {
			return AssetsFrom(itf.Value), true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]*Asset, bool) {
			return lo.ToSlicePtr(AssetsFrom(itf.Value)), true
		}); ok {
			continue
		}

		// public asset
		if ok := assignIf(vf, func() (PublicAsset, bool) {
			a := PublicAssetFrom(itf.Value)
			if a == nil {
				return PublicAsset{}, false
			}
			return *a, true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]PublicAsset, bool) {
			return PublicAssetsFrom(itf.Value), true
		}); ok {
			continue
		}

		if ok := assignIf(vf, func() ([]*PublicAsset, bool) {
			return lo.ToSlicePtr(PublicAssetsFrom(itf.Value)), true
		}); ok {
			continue
		}

		// value
		if ok := assignIf(vf, func() (Value, bool) {
			return *NewValue(itf.Value), true
		}); ok {
			continue
		}

		// primitive
		itfv := reflect.ValueOf(itf.Value)
		res, ok := convertPrimitive(itfv, vf.Type())
		if ok && res.IsValid() {
			vf.Set(res)
		}
	}
}

func Marshal(src any, item *Item) {
	if item == nil || src == nil {
		return
	}

	t := reflect.TypeOf(src)
	if t == nil {
		return
	}

	v := reflect.ValueOf(src)
	if t.Kind() == reflect.Pointer {
		if v.IsNil() {
			return
		}
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	ni := Item{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, opts, _ := strings.Cut(tag, ",")

		if key == "" || key == "-" {
			continue
		}

		ty, optsRemaining, _ := strings.Cut(opts, ",")
		optsSplited := strings.Split(optsRemaining, ",")
		omitempty := !slices.Contains(optsSplited, "includezero")
		isMetadata := slices.Contains(optsSplited, "metadata")

		vf := v.FieldByName(f.Name)
		if key == "id" {
			ni.ID, _ = vf.Interface().(string)
			continue
		}

		vft := vf.Type()
		var value any
		if m, ok := vf.Interface().(MarshalCMS); ok {
			value = m.MarshalCMS()
		} else if vft.Kind() == reflect.Slice && vft.Elem().Implements(marshalCMSType) {
			v := make([]any, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				t := vf.Index(i).Interface().(MarshalCMS).MarshalCMS()
				if t != nil {
					v = append(v, t)
				}
			}
			value = v
		} else if vft.Kind() == reflect.Slice && vft.Elem().Kind() == reflect.String && vf.Len() > 0 {
			st := reflect.TypeOf("")
			v := make([]string, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				vfs := vf.Index(i).Convert(st)
				v = append(v, vfs.String())
			}
			value = v
		} else if vft.Kind() == reflect.Slice && vf.Len() > 0 && (vft.Elem().Kind() == reflect.Struct ||
			vft.Elem().Kind() == reflect.Ptr && vft.Elem().Elem().Kind() == reflect.Struct) {
			isPointer := vft.Elem().Kind() == reflect.Ptr

			v := make([]string, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				var in any
				if isPointer {
					in = vf.Index(i).Interface()
				} else {
					in = vf.Index(i).Addr().Interface()
				}

				item := Item{}
				Marshal(in, &item)
				if item.ID == "" {
					continue
				}

				// assign group
				for i := range item.Fields {
					item.Fields[i].Group = item.ID
				}
				for i := range item.MetadataFields {
					item.Fields[i].Group = item.ID
				}

				// merge i to ni
				ni.Fields = append(ni.Fields, item.Fields...)
				ni.MetadataFields = append(ni.MetadataFields, item.MetadataFields...)

				v = append(v, item.ID)
			}

			if len(v) > 0 {
				value = v
				ty = "group"
			}
		} else if vft.Kind() == reflect.Struct || vft.Kind() == reflect.Ptr && vft.Elem().Kind() == reflect.Struct {
			isPointer := vft.Kind() == reflect.Ptr
			var v any
			if isPointer {
				v = vf.Interface()
			} else {
				v = vf.Addr().Interface()
			}

			item := Item{}
			Marshal(v, &item)
			if item.ID == "" {
				continue
			}

			// assign group
			for i := range item.Fields {
				item.Fields[i].Group = item.ID
			}
			for i := range item.MetadataFields {
				item.Fields[i].Group = item.ID
			}

			// merge i to ni
			ni.Fields = append(ni.Fields, item.Fields...)
			ni.MetadataFields = append(ni.MetadataFields, item.MetadataFields...)

			value = item.ID
			ty = "group"
		} else if !omitempty || !vf.IsZero() {
			value = vf.Convert(vft).Interface()
		}

		if value != nil {
			f := &Field{
				Key:   key,
				Type:  ty,
				Value: value,
			}

			if isMetadata {
				ni.MetadataFields = append(ni.MetadataFields, f)
			} else {
				ni.Fields = append(ni.Fields, f)
			}
		}
	}

	*item = ni
}

func convertPrimitive(v reflect.Value, ty reflect.Type) (reflect.Value, bool) {
	if !v.IsValid() {
		return reflect.Value{}, false
	}

	// slices
	if v.Kind() == reflect.Slice && ty.Kind() == reflect.Slice && v.Len() > 0 {
		toType := ty.Elem()
		zero := reflect.Zero(toType)
		s := reflect.MakeSlice(ty, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, ok := convertPrimitive(v.Index(i), toType)
			if ok && res.IsValid() {
				s = reflect.Append(s, res)
			} else {
				s = reflect.Append(s, zero)
			}
		}

		return s, true
	}

	// maps
	if v.Kind() == reflect.Map &&
		ty.Kind() == reflect.Map &&
		!v.IsNil() &&
		v.Type().Key().AssignableTo(ty.Key()) &&
		v.Len() > 0 {
		toType := ty.Elem()
		m := reflect.MakeMap(ty)
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			res, ok := convertPrimitive(val, toType)
			if ok && res.IsValid() {
				m.SetMapIndex(key, res)
			}
		}

		return m, true
	}

	// interface{}
	if v.Kind() == reflect.Interface && v.Type().NumMethod() == 0 {
		v = v.Elem()
	}

	if !v.IsValid() {
		return reflect.Value{}, false
	}

	if v.Type().AssignableTo(ty) {
		return v, true
	}

	if v.CanConvert(ty) {
		return v.Convert(ty), true
	}

	return reflect.Value{}, false
}

func assignIf[T any](vf reflect.Value, conv func() (T, bool)) bool {
	var t T
	if valueType := reflect.TypeOf(&t); vf.Type().AssignableTo(valueType) {
		v, ok := conv()
		if !ok {
			return false
		}
		vf.Set(reflect.ValueOf(lo.ToPtr(v)))
		return true
	} else if valueType := reflect.TypeOf(t); vf.Type().AssignableTo(valueType) {
		v, ok := conv()
		if !ok {
			return false
		}
		vf.Set(reflect.ValueOf(v))
		return true
	}

	return false
}
