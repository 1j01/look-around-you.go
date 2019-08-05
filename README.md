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

Especially with autoformatting, and especially in combination with autosave,
and given that there's no nonlinear undo history and redos are deleted implicitly upon any changes (as is the status quo)
but just in general,  
DEFINITELY install a plugin to provide local history for files as you edit them,
because the VS Code doesn't provide that built in, and neither does Git.
There's [Local History](https://marketplace.visualstudio.com/items?itemName=xyz.local-history)
and [Checkpoints](https://marketplace.visualstudio.com/items?itemName=micnil.vscode-checkpoints).

So far I've tried Local History so idk which is better.  
It creates a folder `.history` in your project folder. You should add this to your global `.gitignore` (but for convenience, it's included in the local `.gitignore` for this repo)

# The Program

be sure to kill the program promptly haha as it may kill your terminal and crash VS Code or whatever
