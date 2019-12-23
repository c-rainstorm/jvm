package heap

import (
	"fmt"

	"jvm/pkg/classfile"
	"jvm/pkg/classpath"
	"jvm/pkg/global"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	loadedClass map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		loadedClass: make(map[string]*Class),
	}
}

func (this *ClassLoader) LoadClass(classname string) *Class {
	if class, ok := this.loadedClass[classname]; ok {
		return class
	}

	if global.Verbose {
		log.Infof("loadClass start: %v", classname)
	}

	var class *Class
	if this.isArrayClass(classname) {
		class = this.loadArrayClass(classname)
	} else {
		class = this.loadNonArrayClass(classname)
	}

	if global.Verbose {
		log.Infof("loadClass done: %v", classname)
	}

	this.loadedClass[classname] = class

	return class
}

func (this *ClassLoader) isArrayClass(classname string) bool {
	return classname[0] == '['
}

func (this *ClassLoader) loadNonArrayClass(classname string) *Class {
	class := this.defineClass(classname)

	this.validate(class)

	this.prepare(class)

	this.symRefProcess(class)

	this.init(class)

	if global.Verbose {
		log.Infof("Loaded %s", classname)
	}

	return class
}

func (this *ClassLoader) readClass(classname string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(classname)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + classname)
	}
	return data, entry
}

func (this *ClassLoader) defineClass(classname string) *Class {
	data, _ := this.readClass(classname)

	classFile := classfile.Parse(data)

	class := this.newClass(classFile)

	this.resolveSuperClass(class)

	this.resolveInterfaces(class)

	return class
}

func (this *ClassLoader) newClass(classfile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classfile.AccessFlag()
	class.name = classfile.ClassName()
	class.superClassName = classfile.SuperClassName()
	class.interfaceNames = classfile.InterfaceNames()
	class.newConstantPool(classfile.ConstantPool())
	class.newFields(classfile.Fields())
	class.newMethods(classfile.Methods())
	class.classLoader = this
	return class
}

func (this *ClassLoader) resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = this.LoadClass(class.superClassName)
	}
}

func (this *ClassLoader) resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i := 0; i < interfaceCount; i++ {
			class.interfaces[i] = this.LoadClass(class.interfaceNames[i])
		}
	}
}

func (this *ClassLoader) validate(class *Class) {
	// todo
}

func (this *ClassLoader) prepare(class *Class) {
	// 计算实例字段槽号
	class.calInstanceFieldSlotIds()
	// 计算类字段槽号
	class.calStaticFieldSlotIds()
	// 初始化类变量
	class.initStaticFields()
}

func (this *ClassLoader) symRefProcess(class *Class) {
	// do nothing
}

func (this *ClassLoader) init(class *Class) {
	// do nothing
}

const (
	Boolean uint8 = 4
	Char    uint8 = 5
	Float   uint8 = 6
	Double  uint8 = 7
	Byte    uint8 = 8
	Short   uint8 = 9
	Int     uint8 = 10
	Long    uint8 = 11
)

func (this *ClassLoader) LoadPrimitiveArrayClass(aType uint8) *Class {
	switch aType {
	case Boolean:
		return this.LoadClass(global.FdArray + global.FdBoolean)
	case Byte:
		return this.LoadClass(global.FdArray + global.FdByte)
	case Char:
		return this.LoadClass(global.FdArray + global.FdChar)
	case Short:
		return this.LoadClass(global.FdArray + global.FdShort)
	case Int:
		return this.LoadClass(global.FdArray + global.FdInt)
	case Long:
		return this.LoadClass(global.FdArray + global.FdLong)
	case Float:
		return this.LoadClass(global.FdArray + global.FdFloat)
	case Double:
		return this.LoadClass(global.FdArray + global.FdDouble)
	default:
		panic(fmt.Sprintf("不合法的 atype: %d", aType))
	}
}

func (this *ClassLoader) loadArrayClass(classname string) *Class {
	return &Class{
		accessFlags:    ACC_PUBLIC,
		name:           classname,
		superClassName: global.JavaLangObject,
		superClass:     this.LoadClass(global.JavaLangObject),
		interfaceNames: []string{global.JavaLangCloneable, global.JavaIOSerializable},
		interfaces: []*Class{
			this.LoadClass(global.JavaLangCloneable),
			this.LoadClass(global.JavaIOSerializable),
		},
		classLoader: this,
		initStarted: true,
	}
}
