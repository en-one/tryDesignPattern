## 设计模式原则

```bash
SOLID 建立 稳定 灵活 健壮的设计

S  单一职责原则（single responsibility principle）
   		一个类或接口只承担一个职责
O  开闭原则 （open and close）
		对扩展开发 对修改关闭（设计模式的精神领袖）
L  里氏替换原则 （liskov substitution）
		在继承类的时候，务必重写父类所有方法
L  迪米特法则 （law of Demeter）
		减少对象之间的交互，从而减少类之间的耦合
I  接口隔离原则 （interface segregation）
		不要对外暴漏无意义的接口
D  依赖倒置原则 （dependence inversion）
		高层模块依赖与抽象（而不是底层模块）
		细节依赖于抽象（面向接口编程）