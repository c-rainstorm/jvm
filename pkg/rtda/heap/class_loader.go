package heap

import (
	"strings"

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

	if this.isArrayClass(classname) {
		panic("array class can't load for now. " + classname)
	}

	if global.Verbose {
		log.Infof("loadClass start: %v", classname)
	}

	class := this.loadNonArrayClass(classname)

	if global.Verbose {
		log.Infof("loadClass done: %v", classname)
	}

	this.loadedClass[classname] = class

	return class
}

func (this *ClassLoader) isArrayClass(classname string) bool {
	return strings.Contains(classname, "[")
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
