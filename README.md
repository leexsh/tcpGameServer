# 基于TCP的轻量级游戏服务端框架

> 结构参考：./mySource/struct.xmind


- 解决TCP粘包的问题：

    采用[head][body][head][body][head][body]...的形式

    [head]中包含DataLen(消息长度)和ID(消息ID)；
    [body]中包含的是具体的tcp数据。
    所以整个结构就是[Len][ID][Data][Len][ID][Data].....的形式
    ![图片,](./mySource/tcp_pict1.jpg)
<center>解决粘包示意图 (form 飞书文档)</center>

