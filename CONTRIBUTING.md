# Contributing to Go

Go is an open source project.

It is the work of hundreds of contributors. We appreciate your help!

## Filing issues

When [filing an issue](https://golang.org/issue/new), make sure to answer these five questions:

1.  What version of Go are you using (`go version`)?
2.  What operating system and processor architecture are you using?
3.  What did you do?
4.  What did you expect to see?
5.  What did you see instead?

General questions should go to the [golang-nuts mailing list](https://groups.google.com/group/golang-nuts) instead of the issue tracker.
The gophers there will answer or ask you to file an issue if you've tripped over a bug.

## Verifying changes during development

In order to verify changes to the slides or code examples while developing
locally, run the tour from the repository root with your local toolchain:

    $ go run .

The content, static files and templates are embedded in the binary, so the
installed version can also be run from anywhere:

    $ go install github.com/martinkunc/go-tour-cz@latest
    $ go-tour-cz

## Running the App Engine version locally

To view the App Engine version of the slides while developing locally, run:

    $ GAE_ENV=standard go run .

and visit http://localhost:8080 (set `PORT` to use a different port).

The App Engine version runs code examples against the service at play.golang.org.
To verify changes to the code examples you must use your local toolchain to compile
and run the tour locally with `go run .`.

## Contributing code

Please read the [Contribution Guidelines](https://golang.org/doc/contribute.html)
before sending patches.

Unless otherwise noted, the Go source files are distributed under
the BSD-style license found in the LICENSE file.
