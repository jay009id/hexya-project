First Thing First
=================
Make sure to install golang version 1.13. Verify :

```bash
go version
```

Install Hexya
=============

```bash
git clone https://github.com/hexya-erp/hexya
cd hexya
go install
```
Now, hexya will compile and install into `~/go/bin/hexya`.

Make project
============
If you want to make project from scratch, follow this step :

```bash
mkdir hexya-project
cd hexya-project
~/go/bin/hexya module init github.com/jay009id/hexya-project
```

Please note that you may change module `github.com/jay009id/hexya-project` into anything name without need it's already in github.com. It just a name.

Make package
============
Still in `hexya-project`, write syntax below :

```bash
~/go/bin/hexya module new openacademy
```

It create folder `hexya-project/openacademy`. Normally, it will create file `hexya-project/openacademy/000hexya.go`. If it not created, just make it manually.

```go
package openacademy

import (
    "github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "openacademy"

func init() {
    server.RegisterModule(&server.Module{
        Name: MODULE_NAME,
        PreInit: func() {},
        PostInit: func() {},
	})
}
```

Build Package
=============

After write the code, make sure to generate the code to update `pool` folder. Do not touch this folder because it will update everytime we run `~/go/bin/hexya generate .`.

To recognize the `openacademy` package, update `hexya.toml` in `Modules` part. To verify the code, just run :

```bash
~/go/bin/hexya generate .
~/go/bin/hexya updatedb -o
```

If it success, it will create `hexya-project` executable file.

Run It
======

If `hexya-project` created, just run :

```bash
./hexya-project server -o
```

Open the browser, visit `http://localhost:9998`. The port 9998 is set in the file `hexya.toml`. You may change it whatever you want.

