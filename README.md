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