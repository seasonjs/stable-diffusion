# stable-diffusion

pure go ( cgo free ) for stable-diffusion and support cross-platform.

[![Go Reference](https://pkg.go.dev/badge/github.com/seasonjs/stable-diffusion.svg)](https://pkg.go.dev/github.com/seasonjs/stable-diffusion)

sd.go is a wrapper around [stable-diffusion.cpp](https://github.com/leejet/stable-diffusion.cpp), which is an adaption
of ggml.cpp.

<p align="center">
  <img src="./assets/img.png" width="256x">
</p>

## Installation

```bash
go get github.com/seasonjs/stable-diffusion
```

## AutoModel Compatibility

See `deps` folder for dylib compatibility, push request is welcome.

| platform | x32         | x64                     | arm         | cuda           |
|----------|-------------|-------------------------|-------------|----------------|
| windows  | not support | support avx/avx2/avx512 | not support | support cuda12 |
| linux    | not support | support                 | not support |                |
| darwin   | not support | support                 | support     |                |

## AutoModel Dynamic Libraries Disclaimer

#### The Source of dynamic Libraries
These dynamic libraries come from [stable-diffusion.cpp-build release](https://github.com/seasonjs/stable-diffusion.cpp-build/releases), The dynamic library version can be obtained by viewing [stable-diffusion.version file](./deps/stable-diffusion.version)
Anyone can check the consistency of the file by checksum ( MD5 ).

#### The Security Of Dynamic Libraries
All I can say is that the creation of the dynamic library is public and does not contain any subjective malicious logic.
If you are worried about the security of the dynamic library during the use process, you can build it yourself.

**I and any author related to dynamic libraries do not assume any problems, responsibilities or legal liability during use.**

## Usage

This `stable-diffusion` golang library provide two api `Predict` and `ImagePredict`.

Usually you can use `NewAutoModel`, so you don't need to load the dynamic library.

You can find a complete example in [examples](./exmaples) folder.

Here is a simple example:

```go
package main

import (
	"github.com/seasonjs/hf-hub/api"
	sd "github.com/seasonjs/stable-diffusion"
	"io"
	"os"
)

func main() {
	options := sd.DefaultOptions

	model, err := sd.NewAutoModel(options)
	if err != nil {
		print(err.Error())
		return
	}
	defer model.Close()

	hapi, err := api.NewApi()
	if err != nil {
		print(err.Error())
		return
	}

	modelPath, err := hapi.Model("justinpinkney/miniSD").Get("miniSD.ckpt")
	if err != nil {
		print(err.Error())
		return
	}

	err = model.LoadFromFile(modelPath)
	if err != nil {
		print(err.Error())
		return
	}
	var writers []io.Writer
	filenames := []string{
		"../assets/love_cat0.png",
	}
	for _, filename := range filenames {
		file, err := os.Create(filename)
		if err != nil {
			print(err.Error())
			return
		}
		defer file.Close()
		writers = append(writers, file)
	}

	err = model.Predict("british short hair cat, high quality", sd.DefaultFullParams, writers)
	if err != nil {
		print(err.Error())
	}
}
```

## Packaging

To ship a working program that includes this AI, you will need to include the following files:

* libstable-diffusion.dylib / libstable-diffusion.so / stable-diffusion.dll (buildin)
* the model file
* the tokenizer file (buildin)

## Low level API

This package also provide low level Api which is same
as [stable-diffusion-cpp](https://github.com/leejet/stable-diffusion.cpp).
See detail at [stable-diffusion-doc](https://pkg.go.dev/github.com/seasonjs/stable-diffusion).

## Thanks

* [stable-diffusion-cpp](https://github.com/leejet/stable-diffusion.cpp)
* [ggml.cpp](https://github.com/leejet/ggml.cpp)
* [purego](https://github.com/ebitengine/purego)

## Successful Examples
<span>
  <img src="./assets/love_cat0.png" width="128x">
</span>
<span>
  <img src="./assets/love_cat1.png" width="128x">
</span>
<span>
  <img src="./assets/love_cat2.png" width="128x">
</span>
<span>
  <img src="./assets/love_cat3.png" width="128x">
</span>
<span>
  <img src="./assets/love_cat4.png" width="128x">
</span>
<span>
  <img src="./assets/love_cat5.png" width="128x">
</span>

[//]: # (<span>)

[//]: # (  <img src="./assets/love_cat6.png" width="128x">)

[//]: # (</span>)

## License

Copyright (c) seasonjs. All rights reserved.
Licensed under the MIT License. See License.txt in the project root for license information.