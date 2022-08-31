# ms-go-blog

<b>前言:</b>  
&emsp;使用原生golang开发博客,本项目来自 《bilibili 码神之路》；作为一个学造价的跨行生，这个项目让我第一次体会了一个项目从无到有的全过程，第一次感受了编程的魅力；诚然，有很多东西我现在还是搞不懂，希望能在以后的日子里能够一一弄懂！！！



## 功能

1.支持mkdown写作文章；

2.支持文章自定义分类；

3.支持博客全局搜索功能；

4.接入第三方平台valine,实现博客文章评论功能；

5.使用七牛云的go-sdk实现图片上传功能；

6.登录后台时，使用jwt随机生成的token辅助验证；

## 使用方法
- 第一次打开项目时，需要安装相关依赖。如果是使用goland开发工具，只需根据报错一一导入包即可；直接cmd打开，需通过 (go get方式下载包)
- 注意 dao/mysql.go 中的第22行，需要填入自己mysql数据库的账号和密码
- 账号密码为: admin/admin

首页界面如下所示：
![图片](https://user-images.githubusercontent.com/102449999/187653610-5a5d8eef-0f38-4584-a1d7-4abf5c4e9f1f.png)

```
百度网盘链接：
链接：https://pan.baidu.com/s/1X99xHapLRlZ2C5fznAOsJg 
提取码：e1w5
```

<details>
<summary><font size="20" color="orange">开发中所遇到的问题</font></summary>
<pre><code>
- 问题1：未更改golang语句中的连接数据库密码，导致不能正确连接数据库。
![图片](https://user-images.githubusercontent.com/102449999/184500068-d84b5dde-fbd1-4c6d-bd1b-1d2a9ecd0040.png)
<img src="https://user-images.githubusercontent.com/102449999/184500068-d84b5dde-fbd1-4c6d-bd1b-1d2a9ecd0040.png">
</code></pre>
</details>

