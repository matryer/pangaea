pangaea
=======

Powerful pre-processor for text files - powered by JavaScript

## Usage

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