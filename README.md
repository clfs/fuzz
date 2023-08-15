# fuzz

Fuzz tests for open-source Go libraries, covering a variety of data formats.

Built with [ClusterFuzzLite](https://google.github.io/clusterfuzzlite/).

## Reproduce crashes

1. Download and unzip the GitHub artifact from the failed action.

```text
$ unzip crashes-protobuf_FuzzPrototextUnmarshal.zip 
Archive:  crashes-protobuf_FuzzPrototextUnmarshal.zip
   creating: address/
  inflating: address/crash-584153343e3aba1e140db4c787debd23055c2af8.summary  
  inflating: address/crash-584153343e3aba1e140db4c787debd23055c2af8
```

2. Locate the inputs that cause crashes.

```text
$ rg -uu '^Base64: '
address/crash-584153343e3aba1e140db4c787debd23055c2af8.summary
107:Base64: <redacted>
231:Base64: <redacted>
```

