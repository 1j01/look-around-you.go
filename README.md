# Look Around You: Go

Learning [Golang](https://golang.org).  
I thought I'd make a bit of a learning resource of it as I go.

The repo contains snapshots of a single program as I worked on it,
learning about Go features along the way,
sometimes with comments that are removed in the next version.

I started with "Hello world!" and then learned about simple loops, etc.

The end result is not much more complicated than "Hello world!"
(it's maybe in the same tier, still fairly trivial),
but it's a bit more animated, and achieving a specific effect.

## Getting Started

First you have to [download and install Go](https://golang.org/doc/install).

Next you have to learn about Workspaces, which are not one-per-project as you might expect.  
Rather, you're expected to create a single workspace, with a `src` folder,  
with *repositories within the `src` folder.*  
I'm not sure how this plays out for a project where you want Go for parts of it and not for others.

You can clone this repo into the `src` of your Go workspace.  

Next, building and running programs:  

`go build`, `go install`, and `go run` will all compile your program.

`go run` is the simplest to use because you just pass it a `.go` file and it'll compile and run it.  
And then you can hit up+enter in your terminal to rerun it.

[Here's an explanation of `go build` and `go install`](https://stackoverflow.com/a/30612612/2624876).  
To feel the power of native code (if coming from a JavaScript background), you've got to build an executable.  
Go makes this pretty painless, once you've got the workspace stuff sorted out.

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

It's an emulation/recreation of the terminal output of the [Look Around You intro](https://youtu.be/FBaVwwuErmU?t=56),
including typing out the commands letter by letter.

If running in the VS Code terminal, be sure to kill the program promptly haha as it may crash VS Code.

# Releases

Download releases here: https://github.com/1j01/look-around-you.go/releases

Executables were built (cross-compiled!) en-masse with a tool called [gox](https://github.com/mitchellh/gox), and the following bash script/commands:

```bash
go get github.com/mitchellh/gox

cd 1.*/
for i in 1 2 3 4 5 6 7 8 9 10 11
do
cd ../$i.*/
gox -osarch="linux/386 linux/arm linux/s390x linux/mips linux/mipsle darwin/386 freebsd/386 openbsd/386 windows/386 freebsd/arm netbsd/386 netbsd/arm" -output="look-around-you-v$i-{{if eq .OS \"darwin\"}}macOS{{else if eq .OS \"windows\"}}Windows{{else if eq .OS \"linux\"}}Linux{{else if eq .OS \"freebsd\"}}FreeBSD{{else if eq .OS \"netbsd\"}}NetBSD{{else if eq .OS \"openbsd\"}}OpenBSD{{else if eq .OS \"dragonfly\"}}DragonFlyBSD{{else if eq .OS \"plan9\"}}Plan9{{else if eq .OS \"android\"}}Android{{else if eq .OS \"solaris\"}}Solaris{{else if eq .OS \"nacl\"}}NaCl{{else}}{{.OS}}{{end}}-{{if eq .Arch \"386\"}}x86{{else}}{{.Arch}}{{end}}"
done
```

I gave nice names to all the operating systems listed in `gox -osarch-list`, in a complicated [go template](https://gohugo.io/templates/introduction/) which is what the tool supports for formatting.

I considered generating the template using code, since it's really not a good language to author something like that...

```go
names := map[string]string{
	"darwin":    "macOS",
	"linux":     "Linux",
	"windows":   "Windows",
	"freebsd":   "FreeBSD",
	"openbsd":   "OpenBSD",
	"netbsd":    "NetBSD",
	"dragonfly": "DragonFlyBSD",
	"plan9":     "Plan9",
	"android":   "Android",
	"solaris":   "Solaris",
	"nacl":      "NativeClient",
}
```

but I decided it would be easier to just copy and paste and type it out, in that ridiculous blob of template text.

As for what platforms to actually include, I took the `gox -osarch-list` and filtered out 64-bit builds,  
since the program doesn't use hardly any memory at all, and 32-bit should still work on 64-bit operating systems.  
I also filtered out the ones that were not default enabled platforms.  
