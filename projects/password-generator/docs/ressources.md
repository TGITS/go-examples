# Ressources

- [Golang tests in sub-directory](https://stackoverflow.com/questions/19200235/golang-tests-in-sub-directory)
- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [BubbleTea Go Doc](https://pkg.go.dev/charm.land/bubbletea/v2)
- [The Elm Architecture](https://guide.elm-lang.org/architecture/)
- [Migration Guide towards BubbleTea V2](https://github.com/charmbracelet/bubbletea/blob/main/UPGRADE_GUIDE_V2.md)

## The Elm Architecture

From [The Elm Architecture](https://guide.elm-lang.org/architecture/)

The Elm Architecture is a pattern for architecting interactive programs, like webapps and games.

This architecture seems to emerge naturally in Elm. Rather than someone inventing it, early Elm programmers kept discovering the same basic patterns in their code. It was kind of spooky to see people ending up with well-architected code without planning ahead!

So The Elm Architecture is easy in Elm, but it is useful in any front-end project. In fact, projects like Redux have been inspired by The Elm Architecture, so you may have already seen derivatives of this pattern. Point is, even if you ultimately cannot use Elm at work yet, you will get a lot out of using Elm and internalizing this pattern.

The Basic Pattern
Elm programs always look something like this:

The Elm program produces HTML to show on screen, and then the computer sends back messages of what is going on. "They clicked a button!"

What happens within the Elm program though? It always breaks into three parts:

- Model — the state of your application
- View — a way to turn your state into HTML
- Update — a way to update your state based on messages

These three concepts are the core of The Elm Architecture.

## BubbleTea Basics

Bubble Tea programs are comprised of a model that describes the application state and three simple methods on that model:

- `Init`, a function that returns an initial command for the application to run.
- `Update`, a function that handles incoming events and updates the model accordingly.
- `View,` a function that renders the UI based on the data in the model.