# Creational Design Patterns

> **Creational design patterns** are design patterns that deal with object creation mechanisms, trying to create objects in a manner suitable to the situation. The basic form of object creation could result in design problems or in added complexity to the design. Creational design patterns solve this problem by somehow controlling this object creation.

Creational design patterns are composed of two dominant ideas. One is encapsulating knowledge about which concrete classes the system uses. Another is hiding how instances of these concrete classes are created and combined.

Creational design patterns are further categorized into object-creational patterns and class-creational patterns, where object-creational patterns deal with object creation and class-creational patterns deal with class-instantiation. In greater details, object-creational patterns defer part of its object creation to another object, while class-creational patterns defer its object creation to subclasses.

Five well-known design patterns that are parts of creational patterns are:

- [Factory Method](./factory/README.md) uses factory methods to deal with the problem of creating objects without having to specify the exact class of the object that will be created. It solves problems like: 
  - How can an object be created so that subclasses can redefine which class to instantiate?
  - How can a class defer instantiation to subclasses?
- [Abstract Factory](./abstract-factory/README.md) provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes. It solves problems like:
  - How can an application be independent of how its objects are created?
  - How can a class be independent of how the objects it requires are created?
  - How can families of related or dependent objects be created?
- [Builder](./builder/README.md) intent to separate the construction of a complex object from its representation. It lets you construct complex objects step by step, and then it allows you to produce different types and representations of an object using the same construction code. This pattern solves problems like:
  - How can a class (the same construction process) create different representations of a complex object?
  - How can a class that includes creating a complex object be simplified?
- [Prototype](./prototype/README.md) is used when the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects. In simple word, it lets you copy existing objects without making your code dependent on their classes. It solves problems like:
  - How can objects be created so that which objects to create can be specified at run-time?
  - How can dynamically loaded classes be instantiated?
- [Singleton](./singleton/README.md) lets you ensure that a class has only one instance, while providing a global access point to this instance. It solves problems by allowing it to:
  - Ensure that a class only has one instance
  - Easily access the sole instance of a class
  - Control its instantiation
  - Restrict the number of instances
  - Access a global variable