# Proto

> These definitions come directly from Headscale's repo. Best effort to keep them in sync.

- `v1-0.22.3` From release
- `v1-0.23.0-latest` Latest commit [inside /proto](https://github.com/juanfont/headscale/commits/main/proto): `58bd38a`

Protocol buffers are needed to communicate with Headscale over the GRPC API and used for generating the GRPC-gateway.

## Requirements

### Cli tools

Some command line tools are required to generate the gateway.

- Protoc

## Generate

```sh
protoc -I . \
  --go_out ../gen --go_opt paths=source_relative \
  --go-grpc_out ../gen --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ../gen --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ../gen \
  headscale/<VERSION>/headscale.proto
```
