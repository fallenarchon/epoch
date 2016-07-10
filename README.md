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
UTC:	Sun Jul 10 2016 02:50:48.546 PM
Local:	Sun Jul 10 2016 07:50:48.546 AM
```

### Input epoch millisecond timestamp
**Command**

```
epoch 1000000000
```

**Output**

```
Epoch millis:	1000000000
UTC:	Mon Jan 12 1970 01:46:40.000 PM
Local:	Mon Jan 12 1970 05:46:40.000 AM
```

### Input human readable date
**Command**

```
epoch -d '01/12/1970 13:46:40'
```

**Output**

```
Epoch millis:	1000000000
UTC:	Mon Jan 12 1970 01:46:40.000 PM
Local:	Mon Jan 12 1970 01:46:40.000 PM
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

## Installation
Run the following to install in your $GOBIN directory

```
go get github.com/fallenarchon/epoch
```

## License