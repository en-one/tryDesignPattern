什么是类图

> 显示了一组类，接口、协作以及它们之间的关系。类图中还可以包含接口、包等元素，也可以包括对象、链等实例

类与类之间的关系

> 根据关系强度区分
>
> 依赖关系->关联关系->聚合关系->复合关系->泛化关系



- 依赖关系

  **耦合度最小，虚线加箭头**
  **依赖时是在类的内部具体方法中使用到另外一个类**

```c++
public class Person {  
    public void Call(){  
        Phone phone = new Phone();  
    }  
}  
public class Phone {  
  
} 
//Person 与Phone 并无关系，由于偶尔的需要，在方法中进行实例化
//持有Phone类的是Person类的一个方法，而不是Person类，这点是最重要的。
```

- 关联关系

  **耦合度增强，实线加箭头**
  **关联则是作为内部属性来用**

```c++
public class Person { 
 //头部是人的一个属性 
 public Head head; 
 
 public Person() { 
 // TODO Auto-generated constructor 5stub 
 } 
} 

public class Head {  
    public Head() {  
        // TODO Auto-generated constructor stub  
    }  
}
//一旦实例化Person 也就实例化了一个Head，依赖性更强
//持有Head类的是Person的一个属性。在new一个Person时同时实例化了一个Head类作为Person的内部属性
```

- 聚合关系（车和轮胎）

  **依赖性更强，空心的菱形加箭头**

  **聚合是作为构造方法的传参**

```java
public class PersonGroup { 
 public Person person; 
 //将person作为构造方法的参数传进去 
 public PersonGroup(Person person) { 
 // TODO Auto-generated constructor stub 
 this.person = person; 
 } 
} 
public class Person {   
    public Person() {  
        // TODO Auto-generated constructor stub  
    }  
}
//Person 是PersonGroup 构造方法的参数，Person 可以脱离PersonGroup而存在
```

- 复合关系（组合，燕子与翅膀）

  实心的菱形加箭头

```java
public class Person { 
 
 public Foot foot; 
 
 public Person() { 
 // TODO Auto-generated constructor stub 
 //在构造方法中实例化 
 foot = new Foot(); 
 } 
} 
public class Foot {  
    public Foot() {  
        // TODO Auto-generated constructor stub  
    }   
}
//Foot 类在Person类的构造方法中才被具体实例化
//一旦Person实例生成，则Foot实例也生成，当Person实例消亡，其Foot实例也消亡
```

- 泛化关系

  **类与类的继承关系  实线加箭头**

  **类与接口实现关系  虚线加箭头**
  
  
  
  
  
  
  - 依赖关系：A使用到了B（成员变量、函数参数、函数返回值、函数内变量）
    - 泛化关系
        - 继承：A继承B
        - 实现：A实现b接口
    - 关联关系：单向，双向的一对一、一对多关系
        - 聚合关系：整体A与部分B，A和B可以分开存在
        - 复合关系：整体A与部分B，A与B必须同时存在