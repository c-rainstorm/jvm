# 自己动手写 JVM

## 项目规范

1. 项目结构规范遵循 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
2. 日志输出统一使用 [logrus](https://github.com/sirupsen/logrus)

## 实现计划

- [x] 确定项目结构
- [x] 命令行参数解析
- [x] 实现类路径，获取指定类的 `class` 字节流
- [x] 实现 `class` 字节流的解析
- [ ] 实现线程私有运行时数据区
- [ ] 实现解释器及部分JVM指令
- [ ] 实现类加载器、方法区、部分引用类类指令
- [ ] 实现方法调用和返回指令
- [ ] 实现数组相关指令和字符串池
- [ ] 实现本地方法,如 `Objec.hashCode()`
- [ ] 实现 `athrow` 指令进行异常处理
- [ ] 实现 `System.out.println()`，并成功输出 `Hello world!`

## 下载

```bash
go get github.com/c-rainstorm/jvm
```

## 协议声明

MIT

## 参考

- 《自己动手写Java虚拟机》
- 《深入理解Java虚拟机》
