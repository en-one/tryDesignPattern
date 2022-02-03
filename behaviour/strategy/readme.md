## 策略模式 strategy

### 1. 定义
- 定义一系列算法，将每个算法分别封装起来，让他们可以互相替换

###2. 实现
- <font color=Cyan>多态</font>
- <font color=Cyan>包含一个策略接口和一组实现这个接口的策略类</font>
- <font color=Cyan>所有的策略都需要实现策略的接口，用以达到解耦的目的</font>


###3. 应用场景
- 使用的时候，也可以通过<font color=#008000>配置文件 </font>方式指定
- 使用的时候，一般通过配置文件 <font color=#008000>类型type</font>来判断哪个策略
- 创建的时候可以采用工程模式进行创建

###4. 代码解析 
- 存在策略接口 StorageStrategy
- 存在一组实现接口的策略类 fileStorage encryptFileStorage
- 将策略类放在一个公共map里
- 使用的时候根据type判断使用哪个策略

###5. 解耦比较
- <font color=Yellow>工厂</font>：解耦对象的创建和使用
- <font color=Yellow>观察者</font>：解耦观察者和被观察者
- <font color=Yellow>策略</font>：解耦的时 策略的定义，创建和使用