# Server
Server for Infinity OJ

## Technology Stack

1. Web Framework: Gin
5. Registry: Consul
6. Codec: Protobuf
7. Tracing & Metrics: Jaeger

## Prerequisites

1. [Jaeger](https://www.jaegertracing.io/)
2. [Consul](https://www.consul.io/)
3. [Protobuf](https://developers.google.com/protocol-buffers)
4. [Wire](https://github.com/google/wire)

## Development

1. Run consul daemon: `consul agent -dev`
2. Run jaeger daemon: `jaeger-all-in-one`
3. Run postgres service.

for the database:
``` postgresql
create type judge_status as enum ('Pending', 'PartiallyCorrect', 'WrongAnswer', 'Accepted', 'SystemError', 'JudgementFailed', 'CompilationError', 'FileError', 'RuntimeError', 'TimeLimitExceeded', 'MemoryLimitExceeded', 'OutputLimitExceeded', 'InvalidInteraction', 'ConfigurationError', 'Canceled');
```

add following rule to host file:
``` plain
<ip to db, 127.0.0.1> db
<ip to jaeger, 127.0.0.1> jaeger-agent
<ip to consul, 127.0.0.1> consul
```

```bash
make build
make run
```



## Production

## Usage

