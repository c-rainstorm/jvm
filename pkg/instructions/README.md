# JVM Instructions

JVM 指令主要分为以下几大类，具体细节看每个小结的简介

目前所实现的指令，其定义见： [The Java® Virtual Machine Specification](https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf)

JVM 指令只用一个字节来表示，所以叫 字节码，JVM 最多能定义 256 个指令，目前定义了 `0x0[nop] - 0xca[breakpoint]、0xfe[impdep1]、0xff[impdep2]` 205个。

实现时每个指令都会按如下格式进行描述，以 `nop` 指令为例。

```
// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// nop   Page.547
//
// 格式: nop
// 字节: 0x0
// 操作: do nothing
// 描述: do nothing
```

## Constants

## Loads

## Stores

## Stack

## Math

## Conversions

## Comparisons

## References

## Control

## Extended

## Reserved