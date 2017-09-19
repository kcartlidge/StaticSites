# Static Sites

Serve any number of static sites under either HTTP or HTTPS using a single port.

* Serve one or more sites on a single port.
* Zero-configuration HTTPS.
* Production ready; no proxy needed.
* Run locally under HTTP for testing.

---

## If you just want to run your sites

Pre-built executables for Mac, Linux and Windows are available in the ```cmd/builds``` folder.

* Drop the version for your OS into a folder.
* Add your static sites, one subfolder each.
* Add a ```sites.ini``` file - see the sample in the ```sites``` root folder.
* Run ```StaticSites -sites <sites.ini> -port <port> -local <hostname>``` for serving locally.
* Run ```StaticSites -sites <sites.ini> -port 443``` for production serving.
* The default is to look for ```sites.ini``` in the current folder.

---

## If you want to work on the source

These instructions assume Mac OS, but will work on Linux and Windows by switching the path slashes (if needed) and choosing the relevant make script.

### Create a single-platform binary and execute it

``` sh
cd StaticSites/cmd
go build -o builds/macos/StaticSites && ./builds/macos/StaticSites -sites ../sites/sites.ini -port 3000
```

### Create binaries for all platforms

``` sh
cd StaticSites/cmd
./make/macos.sh
```
