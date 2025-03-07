# watson-starship

A simple plugin for [Starship](https://github.com/starship/starship),
which  shows the elapsed time of the current [watson](https://github.com/jazzband/Watson) project.

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

