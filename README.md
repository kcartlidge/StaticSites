[![buddy pipeline](https://app.buddy.works/kcartlidge/staticsites/pipelines/pipeline/69154/badge.svg?token=d3d7d605b6634d88deb5b9f9ed424e77c1fe725dce86fdd4d145c674e3e3aeba "buddy pipeline")](https://app.buddy.works/kcartlidge/staticsites/pipelines/pipeline/69154)

# Static Sites

Serve any number of static sites under either HTTP or HTTPS using a single port.

* Serve one or more sites on a single port.
* Zero-configuration HTTPS.
* Production ready; no proxy needed.
* Run locally under HTTP for testing.

## Documentation

Some snippets of useful info are available below, but in general you should [go to the Static Sites public website](https://staticsites.io) for fuller details.

## Running the Static Sites public website locally

The static sites public website just mentioned is contained within this repository and can also be run locally as an example:

``` sh
cd cmd
./builds/macos/StaticSites -local staticsites.io -sites ../sites/sites.ini
```

Switch around the build folder and slashes according to whether you are running Mac OS, Linux or Windows.

---

## If you just want to run your own sites ...

Pre-built executables for Mac, Linux and Windows are available in the ```cmd/builds``` folder as well as from the public site. [Visit the site](https://staticsites.io) or run it locally for details.

---

## If you want to work on the source ...

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

If you do a rebuild then ensure you clone the ```cmd/builds``` folder to the ```sites/staticsites.io/builds``` folder so the public site will reflect the changes.
