# myke [![Build Status](https://travis-ci.org/goeuro/myke.svg?branch=travis-ci)](https://travis-ci.org/goeuro/myke) [![Latest Release](https://img.shields.io/github/tag/goeuro/myke.svg)](https://github.com/goeuro/myke/releases/latest)

> A higher order task aggregator with cascading configuration, suitable as a wrapper over existing task runners.

myke allows you to define tasks in `.yml` files and aggregates all of them. This helps you to manage multiple components in multiple projects in multiple repositories.

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

Using the above definition, you can invoke tasks like:

* `myke build` runs build in all projects
* `myke <project>/build` runs build in that specific `<project>`
* `myke <tag>/build` runs build in all projects tagged `<tag>`
* `myke <tagA>/<tagB>/.../build` can match tasks by many tags (AND)

You can pass task parameters like:

* `myke template/arg[from=1,to=2]`
* `myke template/file`

## Features

* Define tasks in language-agnostic `.yml` files
* Environment variables can be defined in the yml or included from dotenv files, and can be overridden from the shell
* Runtime parameters like `myke ...task[key1=val1, key2=val2, ...]`
* One YML can mixin another YML, acquiring all tasks, env, env files, PATH defaults, etc, and can override all of them
* Built-in templating using golang text/template and 50+ functions provided by [sprig](https://github.com/Masterminds/sprig)
* Other commands can run other tasks in `before/after` hooks, and they are chained with mixins

## Installation

* [Grab the latest release](https://github.com/goeuro/myke/releases/latest)

## Examples

Explore the self documenting `examples` folder.

## Task Execution Environment

* `cwd` is set to the YML file base folder
* `cwd/bin` is added to `PATH`
* environment variables are loaded from:
  * `env` property in yml
  * dotenv files from `env_files`
  * for every dotenv file, the corresponding dotenv `.local` file is also loaded
* same is done for every mixin that the yml uses
  * So, if you mixin `<some-other-folder>/myke.yml`, then that yml's `cwd/bin` is also added to the PATH, that yml's env/env_files/env_files.local are also loaded, and so on
* shell exported environment variables take precedence
* additional variables: `MYKE_PROJECT`, `MYKE_TASK`, `MYKE_CWD` are set
* command is templated using golang text/template and sprig
  * environment and task arguments are passed in as variables
* command is run using `sh -exc`

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
* `docker-compose run --rm myke /bin/bash` Gives you a terminal inside the container, from where you can run go commands like:
  * `ginkgo -r` Runs all tests
  * `go run main.go` Compiles and runs main
