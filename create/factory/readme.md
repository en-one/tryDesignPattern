## 工厂模式

### 1. 三组工厂模式
- 在工厂模式中中，<font color=Yellow>创建对象不会对客户端暴漏创建逻辑，</font>并且时通过一个共同的接口来
指向新创建的对象，实现了创建者和调用者的分离
  

###2. 实现
- 简单工厂：<font color=Cyan>当我们以NewXXX的方式创建对象/接口的时候，当返回为接口的时候</font>
- 工厂方法：<font color=Cyan>核心的工厂类不再负责所有的产品的创建，而是将具体创建的工作交给子类去做</font>
- 抽象工厂：<font color=Cyan>工厂的工厂，抽象工厂可以创建具体工厂，由具体工厂来产生具体产品。</font>


###3. 如何判断是否使用工厂模式
- <font color=Cyan>封装变化：</font> 封装未工厂类后，当创建逻辑变化时，对调用者透明
- <font color=Cyan>代码复用：</font> 创建代码抽离到独立的工厂类可以复用
- <font color=Cyan>隔离复杂性：</font>封装复杂的创建逻辑
- <font color=Cyan>控制复杂度：</font>将创建代码抽离出来，让园本的函数或类指责更单一


###4. 代码解析
> 【小张家水果】店，有Apple 和 Banana 两种水果
> 
>  【<font color=Yellow>简单工厂</font>】
>  
>  这两种水果都满足Fruit机器的生产要求、现在小张开了一个工厂，叫NewFruit（外人叫他简单工厂），
>  
>  每次生产的时候，你给他类型，由工厂决定生产哪个
> 
> 【<font color=Yellow>工厂方法</font>】
> 
> 当生意越做越大，有专门的苹果工厂生产苹果，香蕉工厂生产香蕉，每个工厂有一个总工厂FruitFactory
> 
> 每次生产的时候，你给他类型，由工厂决定哪个工厂去生产

###4. 应用 DI容器（Dependency Injection Container）
- <font color=Yellow>与单例模式</font>
者使用 