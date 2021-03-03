kou164nkn/cvimg
==================

This repository is the cli used in converting image file.

Development
==================

Build a binary:

```
$ make build
```

Usage
==================

Convert from JPG image under current directory to PNG image. Default source file format is `jpg` and destination format is `png`

```
$ ./cvimg
```

Specify image file format. You can specify `jpg`, `jpeg`, `png`, `gif` formats.

```
$ ./cvimg -src=gif -dst=jpeg
```

Specify a file hierarchy for find to traverse. You can specify absolute path and relative path.

```
$ ./cvimg -dir=../
```