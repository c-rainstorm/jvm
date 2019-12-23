package heap

import (
	"strings"

	"jvm/pkg/classfile"
	"jvm/pkg/global"
)

type Class struct {
	accessFlags uint16
	// 当前类的名称
	name string
	// 父类名
	superClassName string
	// 父类
	superClass *Class
	// 实现的接口列表
	interfaceNames []string
	// 接口
	interfaces []*Class
	// 运行时常量池
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods []*Method
	// 类加载器
	classLoader *ClassLoader
	// 实例字段所占槽数
	instanceSlotCount uint
	// 静态字段所占槽数
	staticSlotCount uint
	// 静态变量表
	staticVars Slots
	// 类已初始化
	initStarted bool
}

func (this *Class) newConstantPool(cfcp classfile.ConstantPool) {
	cpCount := len(cfcp)
	this.constantPool = &ConstantPool{class: this, consts: make([]Constant, cpCount)}
	for i := 1; i < cpCount; i++ {
		cfConstant := cfcp[i]

		switch cfConstant.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cfConstant.(*classfile.ConstantIntegerInfo)
			this.constantPool.consts[i] = intInfo.Val()
		case *classfile.ConstantFloatInfo:
			floatInfo := cfConstant.(*classfile.ConstantFloatInfo)
			this.constantPool.consts[i] = floatInfo.Val()
		case *classfile.ConstantLongInfo:
			longInfo := cfConstant.(*classfile.ConstantLongInfo)
			this.constantPool.consts[i] = longInfo.Val()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cfConstant.(*classfile.ConstantDoubleInfo)
			this.constantPool.consts[i] = doubleInfo.Val()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cfConstant.(*classfile.ConstantStringInfo)
			this.constantPool.consts[i] = stringInfo.Val()
		case *classfile.ConstantClassInfo:
			classInfo := cfConstant.(*classfile.ConstantClassInfo)
			this.constantPool.consts[i] = newClassSymRef(this.constantPool, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := cfConstant.(*classfile.ConstantFieldRefInfo)
			this.constantPool.consts[i] = newFieldSymRef(this.constantPool, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cfConstant.(*classfile.ConstantMethodRefInfo)
			this.constantPool.consts[i] = newMethodSymRef(this.constantPool, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			methodRefInfo := cfConstant.(*classfile.ConstantInterfaceMethodRefInfo)
			this.constantPool.consts[i] = newInterfaceMethodSymRef(this.constantPool, methodRefInfo)
		}
	}
}

func (this *Class) newFields(cfFields []*classfile.MemberInfo) {
	fields := make([]*Field, len(cfFields))
	for i, field := range cfFields {
		fields[i] = &Field{}
		fields[i].class = this
		fields[i].copy(field)
		constAttr := field.ConstValueAttr()
		if constAttr != nil {
			fields[i].constValueIndex = uint(constAttr.ConstantValueIndex())
		}
	}
	this.fields = fields
}

func (this *Class) newMethods(cfMethods []*classfile.MemberInfo) {
	methods := make([]*Method, len(cfMethods))
	for i, method := range cfMethods {
		codeAttr := method.CodeAttr()
		if codeAttr == nil {
			// native 和 abstract 方法没有 code 属性
			methods[i] = &Method{}
		} else {
			methods[i] = &Method{
				maxStack:  uint(codeAttr.MaxStack()),
				maxLocals: uint(codeAttr.MaxLocals()),
				code:      codeAttr.Code(),
			}
		}

		methods[i].class = this
		methods[i].copy(method)
		methods[i].calArgSlotCount()
	}
	this.methods = methods
}

func (this *Class) GetMainMethod() *Method {
	return this.getStaticMethod(global.Main, global.MainDescriptor)
}

func (this *Class) getStaticMethod(methodName string, methodDescriptor string) *Method {
	for _, m := range this.methods {
		if m.name == methodName && m.descriptor == methodDescriptor {
			return m
		}
	}
	return nil
}

func (this *Class) calInstanceFieldSlotIds() uint {
	slotId := uint(0)
	if this.superClass != nil {
		slotId = this.superClass.calInstanceFieldSlotIds()
	}
	for _, field := range this.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsDoubleOrLong() {
				slotId++
			}
		}
	}
	this.instanceSlotCount = slotId
	return slotId
}

func (this *Class) calStaticFieldSlotIds() {
	slotId := uint(0)
	for _, field := range this.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsDoubleOrLong() {
				slotId++
			}
		}
	}
	this.staticSlotCount = slotId
}

func (this *Class) initStaticFields() {
	this.staticVars = newSlots(this.staticSlotCount)
	for _, field := range this.fields {
		if field.IsStatic() {
			if field.constValueIndex > 0 {
				this.initStaticField(field)
			}
		}
	}
}

func (this *Class) initStaticField(field *Field) {
	// 有常量值
	switch field.descriptor {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		val := this.constantPool.GetConstant(field.constValueIndex).(int32)
		this.staticVars.SetInt(field.slotId, val)
	case global.FdLong:
		val := this.constantPool.GetConstant(field.constValueIndex).(int64)
		this.staticVars.SetLong(field.slotId, val)
	case global.FdFloat:
		val := this.constantPool.GetConstant(field.constValueIndex).(float32)
		this.staticVars.SetFloat(field.slotId, val)
	case global.FdDouble:
		val := this.constantPool.GetConstant(field.constValueIndex).(float64)
		this.staticVars.SetDouble(field.slotId, val)
	case global.FdString:
		panic("// todo")
	}
}

func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}

func (this *Class) hasFlag(flag uint16) bool {
	return this.accessFlags&flag != 0
}

func (this *Class) IsInterface() bool {
	return this.hasFlag(ACC_INTERFACE)
}

func (this *Class) IsAbstract() bool {
	return this.hasFlag(ACC_ABSTRACT)
}

func (this *Class) NewObject() *Object {
	return &Object{class: this,
		data: newSlots(this.instanceSlotCount)}
}

func (this *Class) ClassLoader() *ClassLoader {
	return this.classLoader
}

func (this *Class) isAccessibleTo(accessClass *Class) bool {
	return this.IsPublic() || (accessClass.PackageName() == this.PackageName())
}

func (this *Class) IsPublic() bool {
	return this.hasFlag(ACC_PUBLIC)
}

func (this *Class) PackageName() string {
	if i := strings.LastIndex(this.name, global.Slash); i >= 0 {
		return this.name[:i]
	}
	return global.EmptyString
}

func (this *Class) lookupField(name string, descriptor string) *Field {
	for _, field := range this.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, interfaceImpl := range this.interfaces {
		if field := interfaceImpl.lookupField(name, descriptor); field != nil {
			return field
		}
	}

	if this.superClass != nil {
		return this.superClass.lookupField(name, descriptor)
	}

	return nil
}

func (this *Class) IsSubClassOf(class *Class) bool {
	for crtClass := this.superClass; crtClass != nil; crtClass = crtClass.superClass {
		if crtClass == class {
			return true
		}
	}

	return false
}

func (this *Class) StaticSlots() Slots {
	return this.staticVars
}

func (this *Class) IsImplClassOf(interfaceClass *Class) bool {
	for _, inf := range this.interfaces {
		if inf == interfaceClass {
			return true
		}
	}

	if this.superClass == nil {
		return false
	}

	return this.superClass.IsImplClassOf(interfaceClass)
}

func (this *Class) IsSuper() bool {
	return this.hasFlag(ACC_SUPER)
}

func (this *Class) SuperClass() *Class {
	return this.superClass
}

func (this *Class) Name() string {
	return this.name
}

func (this *Class) InitStarted() bool {
	return this.initStarted
}

func (this *Class) StartInit() {
	this.initStarted = true
}

func (this *Class) GetClinitMethod() *Method {
	return this.getStaticMethod("<clinit>", "()V")
}

func (this *Class) NewArray(count int32) *Object {
	if !this.isArray() {
		panic("Not array class: " + this.name)
	}
	switch string(this.name[1]) {
	case global.FdBoolean:
		return &Object{class: this, data: make([]int8, count)}
	case global.FdByte:
		return &Object{class: this, data: make([]int8, count)}
	case global.FdShort:
		return &Object{class: this, data: make([]int16, count)}
	case global.FdChar:
		return &Object{class: this, data: make([]uint16, count)}
	case global.FdInt:
		return &Object{class: this, data: make([]int32, count)}
	case global.FdLong:
		return &Object{class: this, data: make([]int64, count)}
	case global.FdFloat:
		return &Object{class: this, data: make([]float32, count)}
	case global.FdDouble:
		return &Object{class: this, data: make([]float64, count)}
	default:
		return &Object{class: this, data: make([]*Object, count)}
	}
}

func (this *Class) isArray() bool {
	return this.name[0] == '['
}

func (this *Class) ArrayClass() *Class {
	return this.classLoader.LoadClass(global.FdArray + this.descriptor())
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": global.FdBoolean,
	"byte":    global.FdByte,
	"short":   global.FdShort,
	"char":    global.FdChar,
	"int":     global.FdInt,
	"long":    global.FdLong,
	"float":   global.FdFloat,
	"double":  global.FdDouble,
}

func (this *Class) descriptor() string {
	if this.isArray() {
		return this.name
	}

	if primitiveType, ok := primitiveTypes[this.name]; ok {
		return primitiveType
	}

	return global.FdRef + this.name + global.Semicolon
}

func (this *Class) ElementClass() *Class {
	if !this.isArray() {
		panic("Not Array: " + this.name)
	}

	if this.name[1] == '[' {
		return this.classLoader.LoadClass(this.name[1:])
	} else if this.name[1] == 'L' {
		return this.classLoader.LoadClass(this.name[2 : len(this.name)-1])
	} else {
		name := this.name[1:]
		for primitiveClass, value := range primitiveTypes {
			if name == value {
				return this.classLoader.LoadClass(primitiveClass)
			}
		}
	}

	panic("Invalid descriptor: " + this.name[1:])
}

func (this *Class) IsSubInterfaceOf(target *Class) bool {
	for _, inf := range this.interfaces {
		if inf == target || inf.IsSubInterfaceOf(target) {
			return true
		}
	}

	return true
}

func lookupMethod(kls *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(kls, name, descriptor)

	if method == nil {
		method = lookupMethodInInterfaces(kls.interfaces, name, descriptor)
	}

	return method
}

func lookupMethodInInterfaces(interfaces []*Class, name string, descriptor string) *Method {
	for _, inf := range interfaces {
		method := lookupMethodInInterface(inf, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func lookupMethodInInterface(inf *Class, name string, descriptor string) *Method {
	for _, method := range inf.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterfaces(inf.interfaces, name, descriptor)
}

func LookupMethodInClass(kls *Class, name string, descriptor string) *Method {
	for c := kls; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}
