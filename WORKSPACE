workspace(name = "com_github_yugui_gazel")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "53d012d785e2d5c693627c27b95c4185bbeb7d36",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")
go_repositories()

git_repository(
    name = "io_bazel_buildifier",
    remote = "https://github.com/bazelbuild/buildifier.git",
    commit = "84cdc95dd453430af1206c1bfc9e4cddb45e7670",
)
