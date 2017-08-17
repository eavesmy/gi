# gi爬虫    
    
### 接口：
    POST "/api/reptile/go" 
    DATA { domin: xxx , main : "tag1&tag2" }

### 数据库：
    couchDB:5984
	redis: 8004

### 未完成部分：
    数据存储

### 问题：
1. 单核 1G 机器压力大。
2. 部分站点    href    属性包含陷阱，redis写入过多，造成内存疯涨。建议数据直接落地。
3. 爬取条件不完善。
