# go-versioning

This repository demonstrates one way to version go projects.

## Artifacts

* **Go binary**  
    To create a Go binary from this repository's source code, just
    run `make` without arguments in your local copy.
    Every binary built with this command knows its version number. You can
    retrieve it like this:

    ```
    $ ./go-versioning --version
    go-versionining version 0.3.0 build 2106621
    ```

* **Docker image**  
    To create a Docker image with the Go binary in it, run `make docker-image`.
    This will build a Docker image with the output of `git describe --tags --always`
    as tag name. This means that when your repository is currently on the revision
    of tag `v0.3.0`, the Docker image's tag will be `v0.3.0`. If you add one commit
    and build the Docker image again, its tag will be something like `v0.3.0-1-gabc123`.

## Releasing new Versions

Applying the concept of [Semantic Versioning](http://semver.org/), there are three
make targets to create releases:

* `make release-major`
* `make release-minor`
* `make release-patch`

Each of these targets increments the respective version number in the VERSION file,
commits it, and tags the commit with the correct version number.

**Note:**  

While preparing this setup, I created several releases in order to test the setup itself.
Because I went back and forth between commits and edited things in old releases (which is not a good 
practice in general), I added a Make target `retag-releases` that makes sure that all git tags
point to the correct release commits. It's probably not a good idea to do this in a real life
project.
