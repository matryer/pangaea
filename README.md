pangaea
=======

**Powerful pre-processor for text files - powered by JavaScript.**

Pangaea works on any kind of text file including HTML pages, source code files, documents, etc.  The easy-to-use command line tool lets you effortlessly bake Pangaea into your build process, providing you with an extremely powerful text pre-processor wherever you need it.

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

The `pangaea` command reads from the standard in pipe, and writes to the standard out pipe.

```
cat ./example/source.txt | pangaea >> ./example/output.txt
```

Or you can use it in an interactive mode:

```
pangaea
```

  * See some [examples](https://github.com/stretchr/pangaea/tree/master/examples) of how you can use Pangaea.

#### Parameters

Pangaea supports parameters that are made available to your scripts via the `$params` global object.  You can specify parameters using the `-params` flag and a URL encoded query string:

    pangaea -params="name=Mat&age=30"

This will make the `$params["name"]` and `$params["age"]` variables available to your scripts.

  * All parameter values are strings, if you want to use other types you can cast them in your code.

### Download latest

Check out [releases](https://github.com/stretchr/pangaea/releases), or download the source and build from source using `go install`.

### Rules

  * Keep `<script>` and `</script>` tags on their own lines
  * Please report any [issues](https://github.com/stretchr/pangaea/issues)
  * There are no rules

## Development

To build and run the command inside the `cmd` directory, do:

```
clear; go build -o pangaea; cat ./example/source.txt | ./pangaea >> ./example/output.txt
```
