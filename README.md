# Design Patterns

> In software engineering, a **design pattern** is a general repeatable solution to a commonly occurring problem in software design. A design pattern isn't a finished design that can be transformed directly into code. It is a description or template for how to solve a problem that can be used in many different situations.

## Why would we use design patterns?
The truth is that you might manage to work as a programmer for many years without knowing about any pattern. A lot of people do just that. So why would you spend time learning them?

Design patterns are used because they make your job easier. Design patterns let you write better code more quickly. However, patterns are not a panacea. Of the five phases of software development, design patterns do almost nothing in the analysis, testing, or documentation phases. Design patterns, as the name implies, have their biggest impact in the design phase of a project.

To be more precise:
- Design patterns help you analyze the more abstract areas of a program by providing concrete, well-tested solutions.
- Design patterns help you write code faster by providing a clearer picture of how you are implementing the design.
- Design patterns encourage code reuse and accommodate change by supplying well-tested mechanisms for delegation[1] and composition[2], and other non-inheritance based reuse techniques.
- Design patterns encourage more legible and maintainable code by following well-understood paths.
- Design patterns increasingly provide a common language and jargon for programmers.

## Classification of patterns

All patterns can be categorized by their intent, or purpose. This series covers three main groups of patterns:

- [**Creational patterns**](./creational/README.md) provide object creation mechanisms that increase flexibility and reuse of existing code. 
- [**Structural patterns**](./structural/README.md) explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.
- [**Behavioral patterns**](./behavioral/README.md) take care of effective communication and the assignment of responsibilities between objects.

## What does this series consist of?
Most patterns are described very formally so people can reproduce them in many contexts. Here are the sections that are usually present in a pattern description:

- Intent of the pattern briefly describes both the problem and the solution.
- Structure of classes shows each part of the pattern and how they are related.
- Code example in one of the popular programming languages makes it easier to grasp the idea behind the pattern.