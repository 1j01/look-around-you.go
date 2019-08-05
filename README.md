# Look Around You: Go

Learning [Go](https://golang.org).  
I thought I'd make a bit of a learning resource of it as I go.

## Getting Started

First you have to [download and install Go](https://golang.org/doc/install).

Next you have to learn about Workspaces, which are not one-per-project as you might expect.  
Rather, you're expected to create a single workspace, with a `src` folder,  
with *repositories within the `src` folder.*  
I'm not sure how this plays out for a project where you want Go for parts of it and not for others.


Next, building and running programs:  
`go build`, `go install`, and `go run` will all compile your program.  
`go run` is the simplest to use because you just pass it a `.go` file and it'll compile and run it.  
And then you can hit up+enter in your terminal to rerun it.  
[Here's an explanation of `go build` and `go install`](https://stackoverflow.com/a/30612612/2624876)

## Editor Setup

I'm using [Visual Studio Code](https://code.visualstudio.com/) with [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
Hit Install All when it pops up with a notification - this installs some helper Go packages (not other VS Code plugins), into your Go workspace's `bin` folder.

It formats automatically, and gives you syntax error highlighting inline, triggered when you save.
[More info about the plugin](https://code.visualstudio.com/docs/languages/go)

btw, note that autoformatting may be insidious in an editor that doesn't support nonlinear undo history,  
or at least multiple levels of history, i.e. one that's totally time-based (often a totally separate feature)  
(or if you don't know about these options)  
I don't know about any option like this in VS Code  
so I'm liable to lose--  
there's probably a plugin, I bet there's a plugin  
to patch on some editing history data safety  
--to lose any history if I undo and then save (particularly implicitly, with autosave) and it formats, changing the editor contents,  
and erasing all redos, as is the common practice, the status quo.

# The Program

be sure to kill the program promptly haha as it may kill your terminal and crash VS Code or whatever
