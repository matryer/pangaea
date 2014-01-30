pangaea
=======

Powerful pre-processor for text files - powered by JavaScript.

### How it works

```html
This is just a normal text file that has scripts embedded:

<script>
  function name(){
    return "Pangaea";
  }
</script>

We can call the scripts to get the name, which is <%= name() %>.

Other JavaScript also works, so we know that 5 + 10 is <%= 5 + 10 %>.
```

The above file, when processed will output as:

```
This is just a normal text file that has scripts embedded:

We can call the scripts to get the name, which is Pangaea.

Other JavaScript also works, so we know that 5 + 10 is 15.
```

### Usage

The `pangaea` command takes a source file, and a destination file.

```
pangaea -s=file.txt.pangaea -o=file.txt
```

```
Usage of ./pangaea:
  -o="pangaea-out.txt": File to output to
  -s="": File to process
```

### Download latest

Or you can check out other [releases](https://github.com/stretchr/pangaea/releases).

  * Version v0.1 for darwin 386 - [version-v0.1-darwin-386.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-darwin-386.tar)
  * Version v0.1 for darwin amd64 - [version-v0.1-darwin-amd64.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-darwin-amd64.tar)
  * Version v0.1 for freebsd 386 - [version-v0.1-freebsd-386.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-freebsd-386.tar)
  * Version v0.1 for freebsd amd64 - [version-v0.1-freebsd-amd64.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-freebsd-amd64.tar)
  * Version v0.1 for freebsd arm - [version-v0.1-freebsd-arm.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-freebsd-arm.tar)
  * Version v0.1 for linux 386 - [version-v0.1-linux-386.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-linux-386.tar)
  * Version v0.1 for linux amd64 - [version-v0.1-linux-amd64.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-linux-amd64.tar)
  * Version v0.1 for linux arm - [version-v0.1-linux-arm.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-linux-arm.tar)
  * Version v0.1 for windows 386 - [version-v0.1-windows-386.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-windows-386.tar)
  * Version v0.1 for windows amd64 - [version-v0.1-windows-amd64.tar](https://github.com/stretchr/pangaea/releases/download/v0.1.0/version-v0.1-windows-amd64.tar)

### Rules

  * Keep `<script>` and `</script>` tags on their own lines

## Development

To build and run the command inside the `cmd` directory, do:

```
clear; go build -o pangaea; ./pangaea -s=./example/source.txt -o=./example/output.txt
```

The `example/output.txt` file will be overwritten.
