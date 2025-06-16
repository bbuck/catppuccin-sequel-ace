# Catppuccin Themes for Sequel Ace

In the `themes/` folder you can find the output for Catppuccin's Latte, Frappe,
Macchiato, and Mocha themes for use by Sequel Ace. Download the files you want
and import them through the preferences settings to use them as your color
schemes.

If you find any issues with the themes please let me know. I welcome improvements.

## What's with all the Go stuff?

I decided to leverage programmatic templates to build the different schemes.
The template.spTheme is the base and defines the Catppuccin color names for
each desired type, and then those are injected by the templating process. I
just so happened to choose Go to template them because Catppuccin has a Go
library that defines all the themes and color codes, Go supports a templating
language as part of it's standard library, and I like Go. If you came here for
the themes you can find the output in `themes/`. If you want to build them
for yourself or build yoru own from other definitions, feel free to use this
in whole or part.
