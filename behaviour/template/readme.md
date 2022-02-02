## 模板方法 template

### 1. 定义
- 定义一个操作中的算法骨架<font color=yellow>（n个业务逻辑）</font>，并将某些步骤推迟到
子类中实现。可以让子类在不改变算法整体结构情况下，重新定义算法中的某些步骤。

###2. 实现
-  <font color=#008000>子类使用组合的方式获取父类的算法骨架</font>
-  <font color=#008000>子类使用接口实现的方式，重写父类的抽象方法</font>

###3. 应用场景
- 实现一些操作时，整体步骤很固定，但是其中一小部分需要改变
- 数据库的访问封装，servlet中关于doPost方法的调用

###4. 代码解析 
- <font color=ORANGE>CookStep 结构体</font>实现了 Cook 接口的做饭方法
- 存在一个 <font color=ORANGE>DoCook 封装函数</font>，依次调用做饭方法
- 子类 YouMaiCai 和 PaiGu 结构体，组合了CookStep，也可以有Cook 的做饭方法
- 因为上一步的组合，继承Cook接口，重写了其中的某一方法
