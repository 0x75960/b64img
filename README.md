b64img
=======

commandline tool to generate base64-embeded image tag for markdown(html) documents

usage
------

### embed png

```sh
# convert png
$ b64img /path/to/png
```

```html
<image src="data:image/png;base64,..."/>
```

#### another types

```sh
# convert gif
$ b64img -t image/gif /path/to/gif
```

```html
<image src="data:image/gif;base64,..."/>
```

setup
------

### requirement

* go setup

### instruction

```sh
$ go get -u github.com/0x75960/b64img
```
