# ElasticSearch Memory in One

ElasticSearch 内存管理全解，文档基于 `ElasticSearch v7.5` 版本编写，
其他版本可能略有出入，请自行翻阅对应版本的官方文档

## What

什么是ES的内存管理

### 内存拆分

- `ElasticSearch Heap`
    - 说明: ES堆内存，即ES运行时占用总内存大小
    - 官方文档: <https://www.elastic.co/guide/en/elasticsearch/reference/7.5/modules-indices.html>
    - 查询方式:
        - 最大值: `GET _cat/nodes?v&h=hm`
        - 当前占用值:`GET _cat/nodes?v&h=hc`
    - 分解
        - `Fielddata cache`
            - 官方说明: Set limits on the amount of heap used by the in-memory fielddata cache.
            - 查询方式: `GET _cat/nodes?v&h=fm`
        - `Node query cache`
            - 官方说明: Configure the amount of heap used to cache queries results.
            - 查询方式: `GET _cat/nodes?v&h=qcm`
        - `Indexing buffer`
            - 官方说明: Control the size of the buffer allocated to the indexing process.
            - 查询方式: `GET _cat/nodes?v&h=`
        - `Shard request cache`
            - 官方说明: Control the behaviour of the shard-level request cache.
            - 查询方式: `GET _cat/nodes?v&h=rcm`
- `Lucene`

## Where

内存管理的应用

## When

何时进行内存管理（GC）

## Who

谁有权限控制内存，及进行GC处理

## Why

为何进行GC

## How

如何提高GC的效率，方式内存崩溃，从而引起ES雪崩

## 参考文献

> <https://segmentfault.com/a/1190000018558875>
> <https://elasticsearch.cn/article/32>
> <https://www.elastic.co/guide/cn/elasticsearch/guide/current/heap-sizing.html>
> <https://www.elastic.co/cn/blog/found-understanding-memory-pressure-indicator>
> <https://blog.csdn.net/laoyang360/article/details/79998974>
> <https://www.elastic.co/guide/cn/elasticsearch/guide/current/_limiting_memory_usage.html>
