# Structural Design Patterns

> Structural design patterns are design patterns that ease the design by identifying a simple way to realize relationships among entities. It explains how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.

There are following 7 types of structural design patterns:
- [Adapter](./adapter/README.md) allows the interface of an existing class to be used as another interface. It is often used to make existing classes work with others without modifying their source code. This pattern solves problems like:
  - How can a class be reused that does not have an interface that a client requires?
  - How can classes that have incompatible interfaces work together?
  - How can an alternative interface be provided for a class?
- [Bridge](./bridge/README.md) decouples the functional abstraction from the implementation so that the two can vary independently. It is also known as Handle or Body. Bridge design pattern solves:
  - An abstraction and its implementation should be defined and extended independently from each other.
  - A compile-time binding between an abstraction and its implementation should be avoided so that an implementation can be selected at run-time.
- [Composite](./composite/README.md) is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects. It solves problems like:
  - A part-whole hierarchy should be represented so that clients can treat part and whole objects uniformly.
  - A part-whole hierarchy should be represented as tree structure.
- [Decorator](./decorator/README.md) allows behavior to be added to an individual object, dynamically, without affecting the behavior of other objects from the same class. It can solves:
  - Responsibilities should be added to (and removed from) an object dynamically at run-time.
  - A flexible alternative to sub-classing for extending functionality should be provided.
- [Facade](./facade/README.md) is an object that serves as a front-facing interface masking more complex underlying or structural code. In other word, it provides a unified and simplified interface to a set of interfaces in a subsystem, therefore it hides the complexities of the subsystem from the client. It solves problems like:
  - To make a complex subsystem easier to use, a simple interface should be provided for a set of interfaces in the subsystem.
  - The dependencies on a subsystem should be minimized.
- [Flyweight](./flyweight/README.md) is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object. As a result, flyweight objects can:
  - Store intrinsic state that is invariant, context-independent and shareable (for example, the code of character ‘A’ in a given character set)
  - Provide an interface for passing in extrinsic state that is variant, context-dependent and can’t be shared (for example, the position of character ‘A’ in a text document)
- [Proxy](./proxy/README.md), in its most general form, is a class functioning as an interface to something else. It provides the control for accessing the original object, allowing you to perform something either before or after the request gets through to the original object. Proxy design pattern can solves:
  - The access to an object should be controlled.
  - Additional functionality should be provided when accessing an object.