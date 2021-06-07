# IMP 即时通信协议

## 概述

`IMP` 是基于 `TCP` 协议实现的即时通讯协议, 参考了 [MSN](https://tools.ietf.org/pdf/draft-movva-msn-messenger-protocol-00.pdf)
和 [SLCK](https://ieftimov.com/post/understanding-bytes-golang-build-tcp-protocol/). 此协议仅用于计算机网络的学习, 不具备生产可用性.

## 支持的命令

| ID         | 发送者        | 描述                     |
| ---------  | --------     |  --------------         |
| `REG`      | 客户端        | 客户端注册                |
| `CHNS`     | 客户端        | 列出所有的频道             |
| `JOIN`     | 客户端        | 加入某个频道(没有则创建)    |
| `LEAVE`    | 客户端        | 离开某个频道              |
| `MSG`      | 客户端和服务端  | 发送或接收频道的消息       |
| `OK`       | 客户端和服务端  | 命令确认                 |
| `ERR`      | 客户端和服务端  | 错误                    |

接下来依次解释每个命令:

### REG

When a client connects to a server, they can register as a client using the
`REG` command. It takes an identifier as an argument, which is the client's username.

Syntax:

```text
REG <handle>
```

where:

* `handle`: name of the user

### JOIN

When a client connects to a server, they can join a channel using the `JOIN`
command. It takes an identifier as an argument, which is the channel ID.

Syntax:

```text
JOIN <channel-id>
```

where:

* `channel-id`: ID of the channel

### LEAVE

Once a user has joined a channel, they can leave the channel using the `LEAVE`
command, with the channel ID as argument.

Syntax:

```text
LEAVE <channel-id>
```

where:

* `channel-id`: ID of the channel

**Example 1:** to leave the `#general` channel, the client can send:

```text
LEAVE #general
```

### MSG

To send a message to a channel or a user, the client can use the `MSG` command, with the channel or user identifier as
argument, followed with the body length and the body itself.

Syntax:

```text
MSG <entity-type> <entity-id> <length>\r\n[payload]
```

where:

* `entity-id`: the ID of the channel or user
* `length`: payload length
* `payload`: the message body

**Example 1:** send a `Hello everyone!` message to the `#general` channel:

```text
MSG CHN general 16\r\nHello everyone!
```

**Example 2:** send a `Hello!` message to `@jane`:

```text
MSG USER jane 4\r\nHey!
```

### CHNS

To list all available channels, the client can send the `CHNS` message. The server will reply with the list of available
channels.

Syntax:

```text
CHNS
```

### USERS

To list all users in a channel, the client can send the `USERS` message. The server will reply with the list of available users.

Syntax:

```text
USERS <channel-id>
```

### OK/ERR

When the server receives a command, it can reply with `OK` or `ERR`.

`OK` does not have any text after that, think of it as an `HTTP 204`.

`ERR <error-message>` is the format of the errors returned by the server to the client. No protocol errors result in the
server closing the connection. That means that although an `ERR` has been returned, the server is still maintaining the
connection with the client.

**Example 1:** Protocol error due to bad username selected during registration:

```
ERR Username must begin with @
```

**Example 2:** Protocol error due to bad channel ID sent with `JOIN`:

```
ERR Channel ID must begin with #
```
