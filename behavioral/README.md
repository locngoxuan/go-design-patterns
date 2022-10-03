# Behavioral Design Patterns

> Behavioral design patterns are design patterns that identify common communication patterns among objects. By doing so, these patterns increase flexibility in carrying out communication.

In these design patterns, the interaction between the objects should be in such a way that they can easily talk to each other and still should be loosely coupled.

That means the implementation and the client should be loosely coupled in order to avoid hard coding and dependencies.

There are 10 types of behavioral design patterns:
- [Chain Of Responsibility](./chain/README.md) consists of a source of command objects and a series of processing objects. It lets you pass requests along a chain of processing objects. Upon receiving a request, each processing object decides either to process the request or to pass it to the next one in the chain. This pattern solves:
  - Coupling the sender of a request to its receiver should be avoided.
  - It should be possible that more than one receiver can handle a request.
- [Command](./command/README.md) encapsulate a request under an object as a command and pass it to invoker object. Invoker object looks for the appropriate object which can handle this command and pass the command to the corresponding object and that object executes the command. It is also known as Action or Transaction. This pattern can solves these problems:
  - Coupling the invoker of a request to a particular request should be avoided. That is, hard-wired requests should be avoided.
  - It should be possible to configure an object (that invokes a request) with a request.
- [Iterator](./iterator/README.md) is a design pattern in which an iterator is used to traverse a container and access the container’s elements. The Iterator pattern is also known as Cursor. The problems the iterator pattern can solves are:
  - The elements of an aggregate object should be accessed and traversed without exposing its representation (data structures).
  - New traversal operations should be defined for an aggregate object without changing its interface.
- [Mediator](./mediator/README.md) defines an object that encapsulates how a set of objects interact. This pattern is considered to be a behavioral pattern due to the way it can alter the program’s running behavior. It solves:
  - Tight coupling between a set of interacting objects should be avoided.
  - It should be possible to change the interaction between a set of objects independently.
- [Memento](./mediator/README.md) lets you save and restore the previous state of an object without revealing the details of its implementation. The memento pattern can solves problems like:
  - The internal state of an object should be saved externally so that the object can be restored to this state later.
  - The object’s encapsulation must not be violated.
- [Observer](./observer/README.md) is a software design pattern in which an object, named the subject, maintains a list of its dependents, called observers, and notifies them automatically of any state changes, usually by calling one of their methods. The observer pattern is also known as Dependents or Publish-Subscribe. It addresses following problems:
  - A one-to-many dependency between objects should be defined without making the objects tightly coupled.
  - It should be ensured that when one object changes state, an open-ended number of dependent objects are updated automatically.
  - It should be possible that one object can notify an open-ended number of other objects.
- [State](./state/README.md) allows an object to alter its behavior when its internal state changes. This pattern is close to the concept of finite-state machines. The state pattern can be interpreted as a strategy pattern, which is able to switch a strategy through invocations of methods defined in the pattern’s interface. This pattern solves problems like:
  - An object should change its behavior when its internal state changes.
  - State-specific behavior should be defined independently. That is, adding new states should not affect the behavior of existing states.
- [Strategy](./strategy/README.md) defines a family of functionality, encapsulate each one, and make them interchangeable. It also enables selecting an algorithm at runtime. Instead of implementing a single algorithm directly, code receives run-time instructions as to which in a family of algorithms to use.

This pattern is set to solves problem that might (or is foreseen they might) be implemented or solved by different strategies and that possess a clearly defined interface for such cases. Each strategy is perfectly valid on its own with some of the strategies being preferable in certain situations that allow the application to switch between them during runtime.
- [Template Method](./template/README.md) just define the skeleton of a function in an operation, deferring some steps to its subclasses. It is used for following reasons:
  - Let subclasses implement varying behavior (through method overriding)
  - Avoid duplication in the code, the general workflow structure is implemented once in the abstract class’s algorithm, and necessary variations are implemented in the subclasses.
  - Control at what points sub-classing is allowed. As opposed to a simple polymorphic override, where the base method would be entirely rewritten allowing radical change to the workflow, only the specific details of the workflow are allowed to change.
- [Visistor](./visitor/README.md) is a way of separating an algorithm from an object structure on which it operates. In this pattern, element object has to accept the visitor object so that visitor object handles the operation on the element object. Visitor pattern is used to solves problem that should be possible to define a new operation for (some) classes of an object structure without changing the classes.