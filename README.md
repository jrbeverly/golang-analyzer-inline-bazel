# Bazel Golang Inline Analyzer

Experimenting with having analyzers locally defined to a repository, rather than externally defined.

## Notes

- Requires using `go_tool_library` instead of `go_library` due to a dependency change issue (must also use `go_tool_library` of deps)
- Baked natively into `nogo`, so it can be pretty straightforward to test
- Names of types aren't as simple as `package.Type`, but instead include other components (using `HasSuffix`) (What options are there?)
- Change in 'internal/cobrago/storage.go@ListFilesInStorage' can be removed as a test case for the errors
- The `tools/` directory probably isn't the best path. Want something that we can spin-out/externalize as these evolve with the code

Overall this is a pretty good way to start prototyping mechanisms for code analysis that is right next to the code, which can then later be spun out into their own generic analyzers. As it runs automatically with `nogo`, there is a natural way of enabling an analyzer in a tiered manner (warning => error).

What would be necessary to get these setup in as lightweight as possible way to ensure very simple constraints? E.g. `Don't use 'XYZ' type while in 'ABC' module`. The main goal of having these would be essentially creating 'Tests' for the code to ensure that code is being built in a manner idiomatic to the codebase.
