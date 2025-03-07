# watson-starship

A simple plugin for [Starship](https://github.com/starship/starship),
which  shows the elapsed time of the current [watson](https://github.com/jazzband/Watson) project.

## Why?

Unfortunately `watson status` itself is  just too verbose. I only need
the hours elapsed. Manipulating the output were too slow, it lead to a partially hanging prompt.

I tried several variants, among these:

```toml
# using sed
command = "watson status -e | sed -e 's/ ago//' -e 's/ seconds/s/' -e 's/ minutes/m/' -e 's/ hours/h/' -e 's/just now/0/'"

# use perl to parse the status file directly
command = 'perl -n -e "if (/start.: (\d+)/) { \$diff = (time - \$1) / 3600; printf \"%.02fh\n\", \$diff; }" < ~/.config/watson/state'

# use bash, date and https://github.com/TLINDEN/rpnc direclty on the status file
command = 'if [[ "$(grep start ~/.config/watson/state)" =~ ([0-9]+) ]]; then echo $(("$(date +%s)" - "${BASH_REMATCH[1]}")) 3600 / | rpn; fi'
```

This lille tool here is written  in go, reasonably fast, it parses the
watson JSON  status file directly,  calculates the elapsed  time since
the start of the running project and prints it. Pretty simple.

## Setup

You'll need the Golang toolchain for this (version 1.23+).

To build:

```shell
git clone https://github.com/TLINDEN/watson-starship.git
cd watson-starship
make
make install
```

Then add this to your `~/.config/starship.toml`:

```toml
[custom.watson]
when = "watson status  -e | grep -v 'No project'"
style = "green"
command = 'watson-starship'
format = '\[[$output]($style)\]'
```

And add `${custom.watson}\` to your global `format` setting.

## Getting help

Although I'm happy to hear from watson-starship users in private email, that's the
best way for me to forget to do something.

In order to report a bug,  unexpected behavior, feature requests or to
submit    a    patch,    please    open   an    issue    on    github:
https://github.com/TLINDEN/watson-starship/issues.

## Copyright and license

This software is licensed under the GNU GENERAL PUBLIC LICENSE version 3.

## Authors

T.v.Dein <tom AT vondein DOT org>

## Project homepage

https://github.com/TLINDEN/watson-starship

## Copyright and License

Licensed under the GNU GENERAL PUBLIC LICENSE version 3.

## Author

T.v.Dein <tom AT vondein DOT org>

