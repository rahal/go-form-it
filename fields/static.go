package fields

import (
	"github.com/kirves/revel-forms/common"
)

type StaticType struct {
	Field
}

func StaticField(name, content string) *StaticType {
	ret := &StaticType{
		FieldWithType(name, formcommon.STATIC),
	}
	ret.SetText(content)
	return ret
}
