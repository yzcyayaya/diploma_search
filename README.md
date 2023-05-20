# diploma_search
毕业设计，基于云原生的高可用可缩容的信息检索系统


## 设想
- 运行在k8s上面，高可用可缩容 [未完成]
- 界面清爽且移动端和pc端界面都具有可操作性 [未完成]
- 兼容google搜索引擎的语法内容 [未完成]
- 关键字必须智能，能容错，如搜索`lova`,显示为`love`[未完成]

### 新增[`mockjs`](http://mockjs.com/examples.html),模拟数据

##### 运行命令，进入biz/data文件夹下。

```shell
npm install 
```

```shell
node dataRules.js --count {number}
```

> number自行修改为数字，如：100。数据量则为100条。



##### 默认生成json

> biz/data/mock_data.json





### 运行

先生成mock数据，因为`biz/data/init.go`会导入数据到`meilsearch`进去.

所以默认需要生成`biz/data/mock_data.json`该文件。