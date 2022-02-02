## 外观模式 facade

### 1. 定义
- 也叫门面模式
- 隐藏系统的复杂性，将几个细粒度的接口包装下，成一个接口
- 为客户端提供一组访问系统的统一接口

###2. 实现
-  <font color=#008000>向现有的系统添加一个接口，使用这个接口实现系统的复杂性</font>

###3. 应用场景
- 解决易用性问题：用来封装系统的底层实现，隐藏系统的复杂性；例如linxu提供的编程接口
- 解决性能问题： <font color=#008000>减少网络请求性</font>
- 解决分布式事务问题：一次调用就可以在同一个进程的事务中解决

###4. 代码解析 
- 网站之前有登录（Login）和注册接口（Register） 
- 现在，简化流程，直接提供验证码的登录/注册功能
- 使用门面模式，添加一个接口（IUserFacade），实现了登录（Login）和注册接口（Register），
同时，封装了接口（LoginOrRegister），隐藏了系统复杂性