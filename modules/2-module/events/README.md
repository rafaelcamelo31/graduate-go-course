# Events

Events represents a fact or state change in context of system. (Something has happened)

Whenever events happen, you can handle it accordingly.

Example:

```
(Event)
- Insert new User;
  (Handlers)
  - Send an Email;
  - Publish a message to a queue;
  - Notify a user in slack;
  - Insert this User into Salesforce.
```

# How to implement an Event System

```
1st component => Event(Get datas)

2nd component => Operations executed when Event is called

3rd component => Event/Operation Manager
  - Register events and operations
  - Dispatch/Fire an event to execute operations
```
