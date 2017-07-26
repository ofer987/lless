`lless` is the colorizing of `less`. It works similar to `less` but displays content with syntax highlighting.

## Supported Languages

* JavaScript
* Java
* Ruby
* Python
* Go
* C
* JSON

## Installation

### OSX

```
$ brew install lless
```

### Standalone

`lless` can be easily installed as an executable.
Download the latest [compiled binaries](https://github.com/jingweno/lless/releases) and put it in your executable path.

### From source

Prerequisites:
- [Git](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://golang.org/doc/install) 1.4+

```
$ go get -u github.com/jingweno/lless
```

## Usage

```
$ lless FILE1
$ lless FILE1 FILE2 ... --html # output in HTML
$ lless --bg=dark FILE ... # dark background
$ lless # read from standard input
$ curl https://raw.githubusercontent.com/jingweno/lless/master/main.go | lless
```

It's recommended to alias `lless` to `less`:

```
alias less=lless
```

You can always invoke `less` after aliasing `lless` by typing `\less`.

## License

[MIT](https://github.com/jingweno/lless/blob/master/LICENSE)

## Acknowledgements

This project was based off [jingweno's ccat](https://github.com/jingweno/ccat)

## Credits

Thanks to [Sourcegraph](https://github.com/sourcegraph) who built [this](https://github.com/sourcegraph/syntaxhighlight) awesome syntax-highlighting package.
