# eyamluate (v0.0.0)

## eyamluate

### Description

eyamluate command line interface

### Syntax

```shell
eyamluate  [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Show help message.  

### Subcommands

* eval:  
  Evaluates a yaml expression.  

* validate:  
  Validates a yaml file.  

* version:  
  Shows the version of the eyamluate command.  


## eyamluate eval

### Description

Evaluates a yaml expression.

### Syntax

```shell
eyamluate eval [<option>]...
```

### Options

* `-format=<string>`, `-f=<string>`  (default=`"yaml"`):  
  Output format. One of yaml or json.  

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Show help message.  

* `-input-path=<string>`, `-i=<string>`  (default=`""`):  
  Input yaml file path. stdin is used if not provided.  

* `-output-path=<string>`, `-o=<string>`  (default=`""`):  
  Output file path. stdout is used if not provided.  

* `-pretty[=<boolean>]`, `-p[=<boolean>]`  (default=`false`):  
  Pretty print the output.  


## eyamluate validate

### Description

Validates a yaml file.

### Syntax

```shell
eyamluate validate [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Show help message.  

* `-input-path=<string>`, `-i=<string>`  (default=`""`):  
  Input yaml file path. stdin is used if not provided.  

* `-output-path=<string>`, `-o=<string>`  (default=`""`):  
  Output file path.s stdout is used if not provided.  


## eyamluate version

### Description

Shows the version of the eyamluate command.

### Syntax

```shell
eyamluate version
```


