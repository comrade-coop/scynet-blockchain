# ScyNet

This is the source code for [ScyNet](http://www.scynet.ai/), a decentralized network for creating and training autonomous AI agents.

## Running

* Initialize submodules
  ```
  git submodule update --init --recursive
  ```
* Install [Go](https://golang.org/doc/install#install) and [`protoc-gen-go`](https://developers.google.com/protocol-buffers/docs/reference/go-generated#invocation)
  ```
  # Ubuntu
  sudo apt install golang
  go get -u github.com/golang/protobuf/protoc-gen-go

  # Arch
  pacman -S go
  go get -u github.com/golang/protobuf/protoc-gen-go
  # or, yay -S protobuf-go # from AUR

  # Windows / chocolatey
  choco install golang
  go get -u github.com/golang/protobuf/protoc-gen-go
  ```
* Run `scripts/build.sh` to build all the command-line tools.
* Run the resulting executables in the `build/` folder.

## Development

* Install [goimports](https://godoc.org/golang.org/x/tools/cmd/goimportsll)
  ```
  go get -u golang.org/x/tools/cmd/goimports
  # or, yay -S goimports-git # from AUR
  ```
* Run `scripts/format.sh` to format the code properly.
