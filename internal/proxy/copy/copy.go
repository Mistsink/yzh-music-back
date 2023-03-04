package copy

import (
	"fmt"
	"reflect"
	"strconv"
)

type kind uint

const (
	kindInvalid kind = iota
	kindSimple
	kindStruct
	kindSlice
	kindArray
	kindPointer
)

type simpleKind uint

const (
	notSimpleKind simpleKind = iota
	simpleKindInt
	simpleKindFloat
	simpleKindString
)

type Field struct {
	kind   kind
	simple simpleKind
	field  reflect.Value
}

type TagFieldMap map[string]Field

// CopyByTag 需要确保 tag 是结构体中唯一的，否则很可能会出现值覆盖的情况
func CopyByTag(dst any, src any, tag string) error {
	dV, sV := reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem()

	_copyByTag(dV, sV, tag)

	return nil
}

func _copyByTag(dst reflect.Value, src reflect.Value, tag string) {
	tagFieldMapDst, tagFieldMapSrc := parseTagFieldMap(dst, tag), parseTagFieldMap(src, tag)

	for fieldTag, field := range tagFieldMapDst {
		fieldSrc, ok := tagFieldMapSrc[fieldTag]
		if !ok || fieldSrc.kind != field.kind {
			continue
		}

		switch field.kind {
		case kindSimple:
			setSimpleKindValue(field.field, fieldSrc.field)
		case kindStruct:
			// fmt.Println("tag:", fieldTag, "kind: kindstruct")
		case kindPointer:
			// fmt.Println("tag:", fieldTag, "kind: kindpointer")
		case kindArray:
			fallthrough
		case kindSlice:
			_copySlice(field.field, fieldSrc.field, tag)
		}
	}
}

func _copySlice(dst reflect.Value, src reflect.Value, tag string) {
	// no weather dst nil or not
	dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Len()))

	for i := 0; i < dst.Len(); i++ {
		// copy item
		idxSrc := src.Index(i)

		idxDst := dst.Index(i)
		kindDst, _ := resolveKind(idxDst.Kind())
		kindSrc, _ := resolveKind(idxSrc.Kind())
		if kindDst != kindSrc {
			return
		}

		switch kindDst {
		case kindSimple:
			setSimpleKindValue(idxDst, idxSrc)
		case kindStruct:
			_copyByTag(idxDst, idxSrc, tag)
		case kindPointer:

			fmt.Println("kind: ", idxDst.Kind())
			if idxDst.IsNil() {
				idxDst.Set(reflect.New(idxDst.Type().Elem()))
			}
			_copyByTag(idxDst.Elem(), idxSrc.Elem(), tag)
		case kindArray:
			fallthrough
		case kindSlice:
			_copySlice(idxDst, idxSrc, tag)
		}
	}
}

func setSimpleKindValue(dst reflect.Value, src reflect.Value) {
	_, simpleKindD := resolveKind(dst.Kind())
	_, simpleKindS := resolveKind(src.Kind())

	if simpleKindD == simpleKindS {
		dst.Set(reflect.ValueOf(src.Interface()))
		return
	}

	switch {
	case simpleKindD == simpleKindString && simpleKindS == simpleKindInt:
		srcV := src.Interface().(int)
		dst.SetString(strconv.Itoa(srcV))
	case simpleKindD == simpleKindString && simpleKindS == simpleKindFloat:
		srcV := src.Interface().(float32)
		dst.SetString(strconv.FormatFloat(float64(srcV), 'E', -1, 32))
	case simpleKindD == simpleKindInt && simpleKindS == simpleKindString:
		srcStr := src.Interface().(string)
		if srcV, err := strconv.Atoi(srcStr); err != nil {
			dst.SetInt(-1)
		} else {
			dst.SetInt(int64(srcV))
		}
	case simpleKindD == simpleKindFloat && simpleKindS == simpleKindString:
		srcStr := src.Interface().(string)
		if srcV, err := strconv.ParseFloat(srcStr, 32); err != nil {
			dst.SetFloat(-1)
		} else {
			dst.SetFloat(srcV)
		}
	}
}

func parseTagFieldMap(tar reflect.Value, tag string) TagFieldMap {
	ret := make(TagFieldMap)
	_parseTagFieldMap(tar, &ret, tag)
	return ret
}

func _parseTagFieldMap(tar reflect.Value, tagFieldMap *TagFieldMap, tag string) {
	tarType := tar.Type()
	for i := 0; i < tarType.NumField(); i++ {
		fieldT := tarType.Field(i)
		fieldV := tar.Field(i)
		kind, simpleKind := resolveKind(fieldV.Kind())
		if tag, ok := fieldT.Tag.Lookup(tag); ok {
			(*tagFieldMap)[tag] = Field{
				kind:   kind,
				simple: simpleKind,
				field:  fieldV,
			}
		}

		if kind == kindSimple {
		} else if kind == kindStruct {
			_parseTagFieldMap(fieldV, tagFieldMap, tag)
		} else if kind == kindPointer {

			if fieldV.IsNil() && fieldV.CanSet() {
				fieldV.Set(reflect.New(fieldV.Type().Elem()))
			}

			_parseTagFieldMap(fieldV.Elem(), tagFieldMap, tag)
		}
	}
}

func resolveKind(kind reflect.Kind) (kind, simpleKind) {
	switch kind {
	case reflect.Bool:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return kindSimple, simpleKindInt
	case reflect.Uintptr:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		return kindSimple, simpleKindFloat
	case reflect.String:
		return kindSimple, simpleKindString
	case reflect.Array:
		return kindArray, notSimpleKind
	case reflect.Slice:
		return kindSlice, notSimpleKind
	case reflect.Struct:
		return kindStruct, notSimpleKind
	case reflect.Pointer:
		return kindPointer, notSimpleKind
	default:
		return kindInvalid, notSimpleKind
	}
}
