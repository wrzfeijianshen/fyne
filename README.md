
### fyne2 简单介绍
	是基于 fyne-io/fyne 进行更改的,原因是,在下才疏学浅,更改改变不了内部库,所以才另起篇章进行书写.本分支中的所有的fyne-io/fyne 均替换为了github.com/wrzfeijianshen/fyne2,某些东西找不到是很正常的.




<p align="center">
  <a href="https://pkg.go.dev/github.com/wrzfeijianshen/fyne2?tab=doc" title="Go API Reference" rel="nofollow"><img src="https://img.shields.io/badge/go-documentation-blue.svg?style=flat" alt="Go API Reference"></a>
  <a href="https://github.com/fyne-io/fyne/releases/tag/v1.3.0" title="1.3.0 Release" rel="nofollow"><img src="https://img.shields.io/badge/version-1.3.0-blue.svg?style=flat" alt="1.3.0 release"></a>
  <a href='http://gophers.slack.com/messages/fyne'><img src='https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=blue' alt='Join us on Slack' /></a>
  <br />
  <a href="https://goreportcard.com/report/github.com/wrzfeijianshen/fyne2"><img src="https://goreportcard.com/badge/github.com/wrzfeijianshen/fyne2" alt="Code Status" /></a>
  <a href="https://travis-ci.org/fyne-io/fyne"><img src="https://travis-ci.org/fyne-io/fyne.svg" alt="Build Status" /></a>
  <a href='https://coveralls.io/github/fyne-io/fyne?branch=develop'><img src='https://coveralls.io/repos/github/fyne-io/fyne/badge.svg?branch=develop' alt='Coverage Status' /></a>
  <!--a href='https://sourcegraph.com/github.com/fyne-io/fyne?badge'><img src='https://sourcegraph.com/github.com/fyne-io/fyne/-/badge.svg' alt='Used By' /></a-->
</p>

# About

[Fyne](https://fyne.io) is an easy to use UI toolkit and app API written in Go.
It is designed to build applications that run on desktop and mobile devices with a
single codebase.

Version 1.3 is the current release - it introduced file dialogs, notifications and cross
platform data management as well as the Accordion, TextGrid and SplitContainer widgets.
We are now working towards [2.0](https://github.com/fyne-io/fyne/milestone/6)
which aims to add data bindings, animations and more!

# Prerequisites

To develop apps using Fyne you will need Go version 1.12 or later, a C compiler and your system's development tools.
If you're not sure if that's all installed or you don't know how then check out our
[Getting Started](https://fyne.io/develop/) document.

Using the standard go tools you can install Fyne's core library using:

    $ go get github.com/wrzfeijianshen/fyne2

# Widget demo

To run a showcase of the features of Fyne execute the following:

    $ go get github.com/wrzfeijianshen/fyne2/cmd/fyne_demo/
    $ fyne_demo

And you should see something like this (after you click a few buttons):

<p align="center" markdown="1" style="max-width: 100%">
  <img src="cmd/fyne_settings/data/widgets-dark.png" width="654" height="446" alt="Fyne Hello Light Theme" style="max-width: 100%" />
</p>

Or if you are using the light theme:

<p align="center" markdown="1" style="max-width: 100%">
  <img src="cmd/fyne_settings/data/widgets-light.png" width="654" height="446" alt="Fyne Hello Light Theme" style="max-width: 100%" />
</p>

# Getting Started

Fyne is designed to be really easy to code with.
If you have followed the prerequisite steps above then all you need is a
Go IDE (or a text editor). 

Open a new file and you're ready to write your first app!

```go
package main

import (
	"github.com/wrzfeijianshen/fyne2/widget"
	"github.com/wrzfeijianshen/fyne2/app"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
```

And you can run that simply as:

    go run main.go

It should look like this:

<p align="center" markdown="1">
  <img src="img/hello-normal.png" width="207" height="204" alt="Fyne Hello Dark Theme" />
</p>

> Note that Windows applications load from a command prompt by default, which means if you click an icon you may see a command window.
> To fix this add the parameters `-ldflags -H=windowsgui` to your run or build commands.

# Documentation

More documentation is available at the [Fyne developer website](https://fyne.io/develop/) or on [pkg.go.dev](https://pkg.go.dev/github.com/wrzfeijianshen/fyne2?tab=doc).

# Examples

You can find many example applications in the [examples repository](https://github.com/fyne-io/examples/).
Alternatively a list of applications using fyne can be found at [our website](https://apps.fyne.io/).
