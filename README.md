<div align="center">

# Arona Protos

[![Global Version](https://img.shields.io/badge/dynamic/yaml?url=https%3A%2F%2Fraw.githubusercontent.com%2Farisu-archive%2Farona-protos%2Fmaster%2Fversion.txt&query=%24&prefix=v&style=for-the-badge&logo=nexon&label=Global&color=0099ff)](version.txt) [![Latest Release](https://img.shields.io/github/v/release/arisu-archive/arona-protos?style=for-the-badge&logo=github&label=Release)](https://github.com/arisu-archive/arona-protos/releases)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/arisu-archive/arona-protos/protos) [![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)

**Generated Go models for the Global Blue Archive network protocol.**

[Features](#features) • [Installation](#installation) • [Quick start](#quick-start) • [Data model](#data-model) • [Updating generated sources](#updating-generated-sources) • [Documentation](#documentation) • [Testing](#testing) • [Related projects](#related-projects) • [License](#license)

</div>

---

## Features

- Typed request and response models for the Global game service.
- Protocol and web API error enums with generated string conversion and validation helpers.
- Shared packet, session, MX time, and Unity value types.
- FlatBuffers-backed fields through [`arona-flatbuffers`](https://github.com/arisu-archive/arona-flatbuffers).
- Automated regeneration as new game client versions are published.

## Installation

Install the latest release with Go modules:

```bash
go get github.com/arisu-archive/arona-protos@latest
```

Then import the generated package:

```go
import "github.com/arisu-archive/arona-protos/protos"
```

## Quick start

Create a typed request and serialize it with Go's standard JSON package:

```go
request := protos.AccountAuthRequest{
	RequestPacket: protos.RequestPacket{
		BasePacket: protos.BasePacket{
			Protocol: protos.Protocol_Account_Auth,
		},
	},
	DeviceModel: "example-device",
}

payload, err := json.Marshal(request)
if err != nil {
	panic(err)
}

fmt.Println(request.Protocol.String(), string(payload))
```

The package supplies data models rather than transport, authentication, or encryption. Use [go-arona](https://github.com/arisu-archive/go-arona) when you need a higher-level client.

## Data model

Requests embed `RequestPacket`, responses embed `ResponsePacket`, and both build on `BasePacket`. The package also includes protocol identifiers, error codes, game database records, custom JSON types such as `MxTime`, and references to generated FlatBuffers data where the game protocol requires them.

Despite the repository name, these are ordinary Go structs generated from the game's decompiled C# network models; the package does not depend on the Google Protocol Buffers runtime.

## Updating generated sources

The [generation workflow](.github/workflows/generate.yml) downloads the decompiled `com.nexon.bluearchive` client, converts the `MX.NetworkProtocol` namespace with [`bluearchive-togo`](https://github.com/arisu-archive/bluearchive-togo), and regenerates enum helpers. [`version.txt`](version.txt) records the source game client version.

Maintainers can run the C# conversion locally with Bash, `curl`, `unzip`, Go, and `golangci-lint` installed:

```bash
./scripts/generate.sh com.nexon.bluearchive <client-version>
```

The script converts the decompiled models and records the client version. Complete the workflow by regenerating the enum helpers and tidying the module:

```bash
go run github.com/dmarkham/enumer@latest -type=Protocol -trimprefix=Protocol_ ./protos/Protocol.go
go run github.com/dmarkham/enumer@latest -type=WebAPIErrorCode -trimprefix=WebAPIErrorCode_ ./protos/WebAPIErrorCode.go
go mod tidy
golangci-lint run --fix ./...
```

Generated APIs track the game client and can change between module releases. Pin a module version when reproducible builds matter.

## Documentation

Browse the complete package reference on [pkg.go.dev](https://pkg.go.dev/github.com/arisu-archive/arona-protos/protos). The generated API reference is the source of truth for available packets, fields, and enum values.

## Testing

Compile and test the module with:

```bash
go test ./...
```

Run the configured lint suite with:

```bash
golangci-lint run ./...
```

## Related projects

- [`go-arona`](https://github.com/arisu-archive/go-arona) provides a higher-level Global game API client.
- [`arona-flatbuffers`](https://github.com/arisu-archive/arona-flatbuffers) contains the generated FlatBuffers models referenced by this module.
- [`bluearchive-togo`](https://github.com/arisu-archive/bluearchive-togo) converts decompiled Blue Archive C# models into Go.

## License

Arona Protos is available under the [MIT License](LICENSE).
