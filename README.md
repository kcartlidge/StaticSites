# Static Sites

Serve any number of static sites under either HTTP or HTTPS using a single port.

* Serve one or more sites on a single port.
* Zero-configuration HTTPS.
* Production ready; no proxy needed.
* Run locally under HTTP for testing.
* A single 9mb binary deploy (+ sites).
* Uses less then 20mb RAM on Debian 4.19.

## Documentation

Some snippets of useful info are available below, but in general you should [see the Static Sites documentation sample website](http://localhost:8000) by running:

``` sh
cd cmd
./builds/macos/StaticSites -verbose -local staticsites.io -sites ../sites -port 8000
```

Switch around the build folder and slashes according to whether you are running Mac OS, Linux or Windows.

---

## You want to run your own sites

Pre-built executables for Mac, Linux and Windows are available in the `cmd/builds` folder as well as from the public site. Run it as per the above. Remember that the executables are small and standalone, so they can be committed into site repos or as part of a monorepo for convenience.

Parameters/Arguments:

- `-sites` folder containing all site folders
- `-port` port to serve sites on (443 for https, default is 8000)
- `-local` optional site to serve as localhost
- `-verbose` whether to show requests as they occur

---

## You want to work on the source

These instructions assume Mac OS, but will work on Linux and Windows by switching the path slashes (if needed) and choosing the relevant make script.

### Create a single-platform binary and execute it

``` sh
cd cmd
go build -o builds/macos/StaticSites && ./builds/macos/StaticSites -verbose -sites ../sites -port 8000
```

### Create binaries for all platforms

``` sh
cd cmd
./make/macos.sh
```

If you do a rebuild then ensure you copy the `cmd/builds` folder to the `sites/staticsites.io/builds` folder so the public site will reflect the changes.
This will happen automatically if you use the commands above.
