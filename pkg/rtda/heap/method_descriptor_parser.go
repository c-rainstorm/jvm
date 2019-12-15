package heap

// https://github.com/zxh0/jvm.go/blob/129b147ebc/rtda/heap/descriptor.go
// https://github.com/zxh0/jvm.go/blob/129b147ebc/rtda/heap/descriptor.go
// https://github.com/zxh0/jvm.go/blob/129b147ebc/rtda/heap/descriptor.go
type TypeDescriptor string

func (td TypeDescriptor) IsBaseType() bool     { return len(td) == 1 }
func (td TypeDescriptor) IsVoidType() bool     { return td == "V" }
func (td TypeDescriptor) IsObjectType() bool   { return td[0] == 'L' }
func (td TypeDescriptor) IsArrayType() bool    { return td[0] == '[' }
func (td TypeDescriptor) IsLongOrDouble() bool { return td == "J" || td == "D" }

type MethodDescriptor struct {
	ParameterTypes []TypeDescriptor
	ReturnType     TypeDescriptor
}

func (md MethodDescriptor) getParamCount() uint {
	return uint(len(md.ParameterTypes))
}

func (md MethodDescriptor) getParamSlotCount() uint {
	slotCount := md.getParamCount()
	for _, paramType := range md.ParameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}

type MethodDescriptorParser struct {
	d string
}

func calcParamSlotCount(descriptor string) uint {
	return parseMethodDescriptor(descriptor).getParamSlotCount()
}

func parseMethodDescriptor(descriptor string) MethodDescriptor {
	parser := &MethodDescriptorParser{descriptor}
	return parser.parse()
}

func (parser *MethodDescriptorParser) parse() MethodDescriptor {
	if paramTypes, ok := parser.parseParamTypes(); ok {
		if returnType, ok := parser.parseReturnType(); ok {
			return MethodDescriptor{
				ParameterTypes: paramTypes,
				ReturnType:     returnType,
			}
		}
	}
	panic("invalid descriptor: " + parser.d) // TODO
}

func (parser *MethodDescriptorParser) parseReturnType() (TypeDescriptor, bool) {
	if t, ok := parser.parseFieldType(); ok {
		return t, len(parser.d) == 0
	}
	return "V", parser.d == "V"
}

func (parser *MethodDescriptorParser) parseParamTypes() ([]TypeDescriptor, bool) {
	if len(parser.d) == 0 && parser.d[0] != '(' {
		return nil, false
	}
	parser.d = parser.d[1:]

	var ts []TypeDescriptor = nil
	for {
		if t, ok := parser.parseFieldType(); ok {
			ts = append(ts, t)
		} else {
			break
		}
	}

	if len(parser.d) == 0 && parser.d[0] != ')' {
		return nil, false
	}
	parser.d = parser.d[1:]
	return ts, true
}

func (parser *MethodDescriptorParser) parseFieldType() (TypeDescriptor, bool) {
	if len(parser.d) > 0 {
		switch parser.d[0] {
		case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
			t := parser.d[0:1]
			parser.d = parser.d[1:]
			return TypeDescriptor(t), true
		case 'L':
			return parser.parseObjectType()
		case '[':
			return parser.parseArrayType()
		}
	}
	return "", false
}

func (parser *MethodDescriptorParser) parseObjectType() (TypeDescriptor, bool) {
	for i := 0; i < len(parser.d); i++ {
		if parser.d[i] == ';' {
			t := parser.d[:i+1]
			parser.d = parser.d[i+1:]
			return TypeDescriptor(t), true
		}
	}
	return "", false
}

func (parser *MethodDescriptorParser) parseArrayType() (TypeDescriptor, bool) {
	parser.d = parser.d[1:]
	t, ok := parser.parseFieldType()
	return "[" + t, ok
}
