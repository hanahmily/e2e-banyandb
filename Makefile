# Miniature stand-in for banyandb's own Makefile. Only the four targets the
# tdd-bydb workflow's hooks actually invoke exist here — see AGENTS.md in
# hanahmily/vajra for how they were measured. Everything else in the real
# project's `pre-push` (check-req, generate-test-cases, lint, vuln-check) is
# deliberately absent: this repo exists to prove routing/hooks/paths, not to
# replicate banyandb's build.

.PHONY: generate build check pre-push

# Real banyandb's `generate` turns .proto into .pb.go. This repo ships no
# .proto, so there is nothing to generate — but the target must exist and
# exit 0, because pre_push/green_gate/coverage_report all call it
# unconditionally on a resumed run's cold worktree.
generate:
	@echo "generate: no-op (this stand-in ships no .proto)"

build:
	go build ./...

# Real banyandb's `check` ends by failing if the tree is dirty; that is the
# one behaviour of the real gate this stand-in keeps on purpose, since
# tdd-bydb's pre_push hook depends on it (throwaway commit + git reset --soft).
check:
	go vet ./...
	@if [ -n "$$(git status -s)" ]; then \
		echo "check: working tree is not clean:"; \
		git status -s; \
		exit 1; \
	fi

pre-push:
	go build ./...
	go vet ./...
