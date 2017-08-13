### 使用redis管理url的问题：
> redis-go通讯协议不明，不能解决open too many files 问题。

1. 测试 循环存储100000条信息到redis，查看内存和cpu情况。
2. 每一个 getUrl -> getHtml -> SaveInfo -> Done 流程不明确，重做。
