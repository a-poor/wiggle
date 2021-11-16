# wiggle

_created by Austin Poor_

A quick and dirty CLI for moving your mouse around the screen to keep your computer awake.

Have a long-running program to execute? Downloading a large file? Trying to keep your status active in Slack? `wiggle` is the tool for you!


## Installation

```bash
$ go install github.com/a-poor/wiggle@latest
```


## Dependencies

* `github.com/urfave/cli/v2`: For creating a CLI
* `github.com/go-vgo/robotgo`: For detecting the screen size and moving the mouse
* `github.com/briandowns/spinner`: For displaying a terminal spinner while `wiggle` is running


