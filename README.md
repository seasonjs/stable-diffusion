# stable-diffusion
pure go for stable-diffusion and support cross-platform.

[![Go Reference](https://pkg.go.dev/badge/github.com/seasonjs/stable-diffusion.svg)](https://pkg.go.dev/github.com/seasonjs/stable-diffusion)

sd.go is a wrapper around [stable-diffusion-cpp](https://github.com/leejet/stable-diffusion.cpp), which is an adaption of ggml.cpp.

## Installation

```bash
go get github.com/seasonjs/stable-diffusion
```

## Compatibility

Not complete yet.

## Usage

Not complete yet. See `binding_test.go` for detail.

## Packaging

To ship a working program that includes this AI, you will need to include the following files:

* libstable-diffusion.dylib / libstable-diffusion.so / stable-diffusion.dll
* the model file
* the tokenizer file (buildin)

## Low level API

This package also provide low level Api which is same as [stable-diffusion-cpp](https://github.com/leejet/stable-diffusion.cpp).
See detail at [stable-diffusion-doc](https://pkg.go.dev/github.com/seasonjs/stable-diffusion).

## Thanks

* [stable-diffusion-cpp](https://github.com/leejet/stable-diffusion.cpp)
* [ggml.cpp](https://github.com/saharNooby/ggml.cpp)
* [purego](https://github.com/ebitengine/purego)

## License

Copyright (c) seasonjs. All rights reserved.
Licensed under the MIT License. See License.txt in the project root for license information.