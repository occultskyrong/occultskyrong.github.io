# Mongo

文档型NoSQL数据库

## 1. 目录

按照使用及实现顺序排列

<!-- TOC -->

- [1. 目录](#1-目录)
- [2. ORM](#2-orm)
- [3. config](#3-config)
- [4. find](#4-find)
- [5. index](#5-index)
- [6. aggregate](#6-aggregate)
- [7. model](#7-model)

<!-- /TOC -->

## 2. ORM

> [ORM 实例教程 By 阮一峰](http://www.ruanyifeng.com/blog/2019/02/orm-tutorial.html)
>
> ORM 就是通过实例对象的语法，完成关系型数据库的操作的技术，是"对象-关系映射"（Object/Relational Mapping） 的缩写。

## 3. config & connection & client

配置、连接、客户端

## 7. model/schema & index

数据模型，构建和使用索引

## 4. CRUD

增删查改

## count & pagenation

合计数量，分页

## 6. aggregate

聚合

### Aggregation Pipeline Stages

> In the db.collection.aggregate method and db.aggregate method, pipeline stages appear in an array. Documents pass through the stages in sequence.

聚合管道阶段，聚合管道以数组的方式给出，该阶段，文档按数组下标顺序依次经过各个管道。

### Aggregation Pipeline Operators

> These expression operators are available to construct expressions for use in the aggregation pipeline stages.

聚合管道运算符，在聚合管道各个阶段，运算符以带参数函数的形式处理文档；即运算符组成了各个阶段内具体的处理逻辑

#### `$dateToString`

转换`date`型数据为`string`，需要使用`timezone`指定时区，否则按照`UTC`时间转换

```JavaScript
{ $dateToString: { format: "%Y-%m-%d %H:%M:%S:%L%z", date: "$createdAt", timezone: "+08:00" } }
```
