# Epoch CLI

A small CLI utility for working with epoch timestamps

## Usage

### Display current epoch millisecond timestamp
**Command**

```
epoch
```

**Output**

```
Epoch millis:	1468162248546
UTC:	Sun Jul 10 2016 14:50:48.546
Local:	Sun Jul 10 2016 07:50:48.546
```

### Input epoch millisecond timestamp
**Command**

```
epoch 1000000000
```

**Output**

```
Epoch millis:	1000000000
UTC:	Mon Jan 12 1970 13:46:40.000
Local:	Mon Jan 12 1970 05:46:40.000
```

### Input human readable date
**Command**

```
epoch -d '01/12/1970 13:46:40'
```

**Output**

```
Epoch millis:	1000000000
UTC:	Mon Jan 12 1970 13:46:40.000
Local:	Mon Jan 12 1970 13:46:40.000
```

### Strip human readable date

**Command**

```
epoch -s
```

**Output**

```
1468162984088
```

### Display 12 hour clock instead of 24 hour clock

**Command**

```
epoch -12
```

**Output**

```
Epoch millis:	1468163459218
UTC:	Sun Jul 10 2016 03:10:59.218 PM
Local:	Sun Jul 10 2016 08:10:59.218 AM
```

## Installation
Run the following to install in your $GOBIN directory

```
go get github.com/fallenarchon/epoch
```

## License