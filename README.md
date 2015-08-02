# Ace [Serve]

Use Ace to serve up static files.

## Install
`go install github.com/samuelcouch/ace`

As long as `$GOPATH/bin` is in your `$PATH`, `ace` will now be available to you.

## Usage
From the directory where you want to serve up files from simply call
`ace`

### Options
* `-p` the port to serve on (defaults to `3000`)
* `-d` the directory to serve up (defaults to your current directory)
```shell
ace -p 3456 -d ~/some/other/directory
```

