# e2e-banyandb

A miniature stand-in for `github.com/apache/skywalking-banyandb`, used by
[vajra](https://github.com/hanahmily/vajra)'s `tdd-bydb` end-to-end suite.

The suite exists to verify workflow routing, hooks, and agent file I/O — not
to re-run a third party's real build on every case. This repo declares the
same module path and exposes the same four `make` targets banyandb's own
`Makefile` does that vajra's hooks actually invoke (`generate`, `build`,
`check`, `pre-push`), each doing the minimum real work needed to make those
targets meaningful (a real `go build`/`go vet`, and `check`'s dirty-tree
assertion), so the gate stays real while the tree stays tiny.

This is not a fork of banyandb and carries none of its code.
