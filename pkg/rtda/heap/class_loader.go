package heap

import (
	"fmt"

	"jvm/pkg/classfile"
	"jvm/pkg/classpath"
	"jvm/pkg/global"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	loadedClass map[string]*ClassObject
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	classloader := &ClassLoader{
		cp:          cp,
		loadedClass: make(map[string]*ClassObject),
	}
	classloader.loadBaseClass()
	classloader.loadPrimitiveClasses()
	return classloader
}

func (this *ClassLoader) LoadClass(classname string) *ClassObject {
	if class, ok := this.loadedClass[classname]; ok {
		return class
	}

	if global.Verbose {
		log.Infof("loadClass start: %v", classname)
	}

	var class *ClassObject
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

func (this *ClassLoader) loadNonArrayClass(classname string) *ClassObject {
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

func (this *ClassLoader) defineClass(classname string) *ClassObject {
	data, _ := this.readClass(classname)

	classFile := classfile.Parse(data)

	class := this.newClass(classFile)

	this.resolveSuperClass(class)

	this.resolveInterfaces(class)

	return class
}

func (this *ClassLoader) newClass(classfile *classfile.ClassFile) *ClassObject {
	var class *ClassObject
	if classClass, ok := this.loadedClass[global.JavaLangClass]; ok {
		class = classClass.NewObject().(*ClassObject)
	} else {
		class = &ClassObject{}
	}

	class.accessFlags = classfile.AccessFlag()
	class.name = classfile.ClassName()
	class.superClassName = classfile.SuperClassName()
	class.interfaceNames = classfile.InterfaceNames()
	class.newConstantPool(classfile.ConstantPool())
	class.newFields(classfile.Fields())
	class.newMethods(classfile.Methods())
	class.sourceFile = classfile.SourceFile()
	class.classLoader = this
	return class
}

func (this *ClassLoader) resolveSuperClass(class *ClassObject) {
	if class.name != global.JavaLangObject {
		class.superClass = this.LoadClass(class.superClassName)
	}
}

func (this *ClassLoader) resolveInterfaces(class *ClassObject) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*ClassObject, interfaceCount)
		for i := 0; i < interfaceCount; i++ {
			class.interfaces[i] = this.LoadClass(class.interfaceNames[i])
		}
	}
}

func (this *ClassLoader) validate(class *ClassObject) {
	// todo
}

func (this *ClassLoader) prepare(class *ClassObject) {
	// 计算实例字段槽号
	class.calInstanceFieldSlotIds()
	// 计算类字段槽号
	class.calStaticFieldSlotIds()
	// 初始化类变量
	class.initStaticFields()
}

func (this *ClassLoader) symRefProcess(class *ClassObject) {
	// do nothing
}

func (this *ClassLoader) init(class *ClassObject) {
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

func (this *ClassLoader) LoadPrimitiveArrayClass(aType uint8) *ClassObject {
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

func (this *ClassLoader) loadArrayClass(classname string) *ClassObject {
	arrayClass := this.loadedClass[global.JavaLangClass].NewObject().(*ClassObject)

	arrayClass.accessFlags = ACC_PUBLIC
	arrayClass.name = classname
	arrayClass.superClassName = global.JavaLangObject
	arrayClass.superClass = this.LoadClass(global.JavaLangObject)
	arrayClass.classLoader = this
	arrayClass.initStarted = true
	arrayClass.interfaceNames = []string{global.JavaLangCloneable, global.JavaIOSerializable}
	arrayClass.interfaces = []*ClassObject{
		this.LoadClass(global.JavaLangCloneable),
		this.LoadClass(global.JavaIOSerializable),
	}

	return arrayClass
}

func (this *ClassLoader) loadBaseClass() {
	// 触发 java.lang.Class、java.lang.BaseObject 类加载，此时  java.lang.Class 还未被加载成功，
	// 所以当前类的 *ClassObject 字段都是 nil，
	this.LoadClass(global.JavaLangClass)
	// 获取非空的 java.lang.Class
	classClass := this.loadedClass[global.JavaLangClass]

	// 逐一赋值
	for _, classObject := range this.loadedClass {
		classObject.NormalObject = &NormalObject{
			BaseObject: BaseObject{class: classClass},
			slots:      newSlots(classClass.instanceSlotCount),
		}
	}
}

func (this *ClassLoader) loadPrimitiveClasses() {
	for name, _ := range primitiveTypes {
		this.loadedClass[name] = this.loadPrimitiveClass(name)
	}
}

func (this *ClassLoader) loadPrimitiveClass(name string) *ClassObject {
	primitiveClass := this.loadedClass[global.JavaLangClass].NewObject().(*ClassObject)

	primitiveClass.accessFlags = ACC_PUBLIC
	primitiveClass.name = name
	primitiveClass.initStarted = true
	primitiveClass.classLoader = this
	return primitiveClass
}
