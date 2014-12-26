tp
==

Extremely Lightweight Template Processor (for command line tools)

# WHAT IS THIS?

There are many occasions where you are writing some sort of shell script and
you just want to write that extremely simple template... You fetch a JSON
file from somewhere, and just pipe it through a template processor and
get the result. `tp` is that.

`tp` is minimalistic: It receives variables in JSON format from standard input,
a template file from the command line, and spews out the result to standard
output. The template is expected to be in Go's text/template format.

That's all. If you look at the source code, it's dead simple.

# EXAMPLE

So suppose in an adhoc shell script I'd like create markdown file containing
my github repos, and links to them. First I create my template at `repos.md.tmpl`:

```text
# My Repositories!

{{range .}}* [{{ .name }}]({{ .html_url }})
{{end}}
```

Then I call:

```bash
curl https://api.github.com/users/lestrrat/repos | tp repos.md.tmpl > repos.md
```

And voila, I get a nice markdown file!

![](http://peco.github.io/images/tp-demo.gif)
