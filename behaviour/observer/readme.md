## 观察者模式 observe

### 1. 定义
- 也叫 <font color=Yellow>发布订阅模式</font>，在对象之间定义一个一对多的依赖，当一个对象状态
改变的时候，所有依赖于他的对象都会接收到通知

###2. 实现
-  <font color=Cyan>被观察者 有一个集合 存放 观察者</font>

###3. 应用场景
- 消息队列

###4. 代码解析
- 顾客A 和 顾客B 实现了 <font color=ORANGE>Customer 顾客接口</font>
- <font color=ORANGE>报社结构体</font> 的属性里有一个数组，为<font color=ORANGE>[ ]Customer</font>
- 当顾客关注报社的时候，把顾客存到<font color=ORANGE>报社结构体</font>的<font color=ORANGE>[ ]Customer</font>中
- 当报社有消息时，通知结构体中的顾客
