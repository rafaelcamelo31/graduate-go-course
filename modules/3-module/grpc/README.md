# gRPC

gRPC is a communication protocol framework, developed by Google. It efficiently connects with various backend service, mobile and browser (still maturing).

It is language agnostic, fast, light and gives an advantage over standard http communication.

## Remote Procedure Call (RPC)

A Remote Procedure Call (RPC) is a software communication protocol that one program uses to request a service from another program located on a different computer and network, without having to understand the network's details. Specifically, RPC is used to call other processes on remote systems as if the process were a local system. A procedure call is also sometimes known as a function call or a subroutine call. [RPC](https://www.techtarget.com/searchapparchitecture/definition/Remote-Procedure-Call-RPC)

## Protocol Buffers

Protocol buffers are Google’s language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages. [protobuf](https://protobuf.dev/)

## Protobuf vs JSON

Protobuf

- Binary
- Relatively smaller size
- Guarantees type-safety
- Prevents schema-violations
- Gives you simple accessors
- Fast serialization/deserialization
- Backward compatibility

JSON

- Plain text
- Human readable/editable
- Can be parsed without knowing schema in advance
- Excellent browser support
- Less verbose than XML

## HTTP/2

- Datas are binary and not text like in HTTP 1.1
- Uses same TCP connection to receive and send data between client and server (Multiplex)
- Server Push
- Compressed headers
- Less network resource consumption
- Faster processing
