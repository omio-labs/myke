# myke [![Latest Release](https://img.shields.io/github/tag/goeuro/myke.svg)](https://github.com/goeuro/myke/releases/latest) [![Build Status](https://travis-ci.org/goeuro/myke.svg?branch=travis-ci)](https://travis-ci.org/goeuro/myke) [![Go Report Card](https://goreportcard.com/badge/github.com/goeuro/myke)](https://goreportcard.com/report/github.com/goeuro/myke) [![codecov](https://codecov.io/gh/goeuro/myke/branch/master/graph/badge.svg)](https://codecov.io/gh/goeuro/myke)

> myke makes it easy to write development tasks

Development scripts usually begin as a simple shell wrapper with switch cases (`service.sh start|stop|etc`), and then aggregate multiple scripts, add arguments, discovery/listing, environment variable handling, then easy overriding and defaults, and soon manipulating files based on these variables, `sed|awk|envsubst`, then proceed to python/ruby/etc with some real templating, then start adding dependencies, then become projects themselves with a checkout/setup process :trollface:

myke solves all these problems in a single tiny binary, to avoid reinventing the same stuff over and over again.

## Features

* Define tasks in simple `.yml` files
* Tasks execute in a predictable environment irrespective of which folder they are invoked from
* Nice aggregation and discovery with tag-based grouping, suitable for few and many tasks, organizing into subfolders/submodules/repos/projects
* Robust environment handling - Can be defined as keys in the YML or as dotenv files, overridden by dotenv.local files, `PATH` is always prepended, shell always takes precedence
* Built-in templating using golang text/template and 50+ functions provided by [sprig](https://github.com/Masterminds/sprig)
* Mixin ymls to share tasks, envvars, etc
* Runtime arguments like `myke task1 --key1=val1 task2 --key2=val2 ...`
* `before/after/error` hooks to perform cleanups, chains with mixins, etc
* `retry` support with max and delay for failing tasks
* Tiny, cross-platform binaries
* and a lot of small utilities packed in

## Usage

Create `myke.yml` with tasks. For example, running `myke` on this folder prints:

```
  PROJECT  |    TAGS    |             TASKS
+----------+------------+-------------------------------------+
  myke     |            | test
  example  |            | build
  env      |            | env
  tags1    | tagA, tagB | tag
  tags2    | tagB, tagC | tag
  depends  |            | after, before, before_after, itself
  template |            | args, file
  mixin    |            | task2, task3, task1
```

Using the above myke.yml, you can invoke tasks like:

* `myke build` runs build in all projects
* `myke <project>/build` runs build in that specific `<project>`
* `myke <tag>/build` runs build in all projects tagged `<tag>`
* `myke <tagA>/<tagB>/.../build` can match tasks by many tags (AND)
* `myke task1 --key1=val1 task2 --key2=val2 ...` passes arguments to individual tasks

## Installation

* [Grab the latest release](https://github.com/goeuro/myke/releases/latest)

## Examples

Explore the self documenting `examples` folder.

## Task Execution Environment

* tasks always run with `cwd` set to the folder where the task is defined
* `cwd/bin` is always added to `PATH`
* environment variables are loaded from:
  * `env` property in yml
  * dotenv files from `env_files`
  * for every dotenv file, the corresponding dotenv `.local` file is also loaded if present
* same is done for every mixin that the yml uses
  * So, if you mixin `<some-other-folder>/myke.yml`, then that yml's `cwd/bin` is also added to the PATH, that yml's env/env_files/env_files.local are also loaded, and so on
* shell exported environment variables take precedence
* additional variables: `$MYKE_PROJECT`, `$MYKE_TASK`, `$MYKE_CWD` are always set
  * `$myke` is set to full path of myke itself to easily nest myke calls (e.g. `$myke do_something` will become `myke.exe do_something` in windows)
* command is templated using golang text/template and sprig
  * environment and task arguments are passed in as variables
* command is run using `sh -exc` in Linux/OSX, and `cmd.exe /C` in Windows

## FAQs

### How do I share common logic in tasks?

There are multiple ways including:

* Place shared scripts in `bin` folder (remember that `CWD/bin` is always added to the `PATH`). If the scripts are complex, you can write them in any scripting language of your choice
* If multiple projects need to share the same scripts, then use a common mixin folder (remember that for mixin ymls - the same `CWD/bin` is added to PATH, same env files are loaded, etc, refer Task Execution Environment above)

For example,

* `java-mixin`
  * `myke.yml` - project template with tasks
  * `myke.env` - environment vars, can be overridden by extending projects
  * `bin` - gets added to the PATH of extending projects
    * any shared scripts that you want
* `kubernetes-mixin`
  * ...
  * ...

### Why use myke?

Deferring higher order build logic (like reading scm history for changelogs, updating scm tags/branches, generating version numbers, etc) to a meta-build tool (like a task runner or aggregator), restricting build tools to do only simple source builds, and having a shared build vocabulary across projects is a generally good idea. There are millions of such meta-build tools or task aggregators out there, we just wanted something fast, zero-dependency and language-agnostic while still helping us manage multiple components across repositories with ease.

In that sense, `myke` is never a build or deployment tool, its just a task aggregator. Its not designed to be an alternative for source build tools, rather it just augments them. The comparison below is on that same perspective.

* `maven` is a lifecycle reactor and/or project management tool that does a lot of things (compilation/scm/release/lifecycle/build/etc), except its hard to use it as a simple task runner. myke focuses only on the latter
* `bazel` `buck` `pants` `gradle` `...` replace your current buildchain by giving you a totally new DSL to compile your programs (`java_binary`, etc). myke simply acts as a yml-based interface to your existing tools and workflows, thereby not needing to change your project and IDE setup
* `grunt` `gulp` `pyinvoke` `rake` `sake` `thor` `...` myke is zero-dependency, language agnostic, uses simple yml and allows aggregation of tasks through hierarchies, templates and tags
* `make` `scons` `ninja` `...` they are low-level build tools with a crux of file-based dependencies. Most buildchains today are already intelligent enough to process only changed files, so myke completely bypasses file tracking and only focuses on task aggregation and discoverability
* `capistrano` `fabric` `...` myke is not a deployment tool for remote machines, and does not do anything over SSH
* `ansible` `salt` `...` myke is not a configuration management tool, its a task runner
* [`robo`](https://github.com/tj/robo) is the closest relative to myke, you should check it out as well

## Development

Use docker/docker-compose to develop. You don't need to have golang installed.

* `docker-compose build` Builds and runs tests
* `docker-compose up` Produces `bin` folder with executables
* `docker-compose run --rm default /bin/bash` Gives you a terminal inside the container, from where you can run go commands like:
  * `go test ./...` Runs all tests
  * `go run main.go` Compiles and runs main
