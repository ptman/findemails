# findemails - library for extracting emails from text

I didn't find a library (or snippet, really) for this when I went looking, so
here is my solution.

Remember the [Go proverb](https://go-proverbs.github.io/):

> A little copying is better than a little dependency.

I'm aware of how hard it is to validate email addresses against RFC5322 using
[regular expressions](http://emailregex.com/). This library is focused on
practical cases, not 100% spec implementation. Please include real-world
examples of problems with issues and PRs.
