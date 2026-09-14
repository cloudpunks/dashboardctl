# Dashboardctl

[![General Workflow](https://github.com/cloudpunks/dashboardctl/actions/workflows/general.yml/badge.svg)](https://github.com/cloudpunks/dashboardctl/actions/workflows/general.yml) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/3322632d759b44d398633123e082e223)](https://app.codacy.com/gh/cloudpunks/dashboardctl/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/cloudpunks/dashboardctl.svg)](https://pkg.go.dev/github.com/cloudpunks/dashboardctl) [![Go Report Card](https://goreportcard.com/badge/github.com/cloudpunks/dashboardctl)](https://goreportcard.com/report/github.com/cloudpunks/dashboardctl) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/cloudpunks/dashboardctl)

> [!WARNING]
> **This project is in early development.** The builtin resources are
> **not yet stable** and can introduce **breaking changes** at any time. Pin
> your version, review the changelog before upgrading, and do not use this
> for production workloads until it reaches a stable release.

A commandline client to interact with Curseforge. For now it's mostly used to
fetch mods defined within modpack manifests.

## Install

You can download prebuilt binaries from the [GitHub releases][releases]. If you
prefer to use containers you could use our images published on [GHCR][ghcr]. If
you need further guidance how to install this take a look at our [docs][docs].

## Prerequisites

We use [mise][mise] to manage all required tools and their versions. Install it
by following the [official installation instructions][mise-install], then run
the following commands inside the repository to activate mise and install all
tools defined in `mise.toml`:

```console
mise trust
mise install
```

## Build

Since all required commands ar part of our [go-task][gotask] taskfile the
commands you got to execute are quite simple:

```console
git clone https://github.com/cloudpunks/dashboardctl.git
cd dashboardctl

task build
./bin/dashboardctl -h
```

## Development

To start developing on this project you have to execute only a few commands in
multiple terminal tabs or windows:

```console
task watch
```

After that you can simply execute the tool via `bin/dashboardctl -h`. Generally
it supports hot reloading which means the binary gets automatically recompiled
on code changes.

## Security

If you find a security issue please contact
[info@cloudpunks.de](mailto:info@cloudpunks.de) first.

## Contributing

Generally we are following [conventional commits][commits] when we apply
changes. That way we are able to generate proper changelogs for every release.
Please use always pull requests to integrate new functionalities or to fix
issues.

For the release process we are following [semantic versioning][semver] which
clearly indicates if a new version just resolves bugs, includes new features or
even includes breaking changes.

After installing the tools via `mise install` as described above set up the
pre-commit hooks so they run automatically on every commit:

```console
prek install --hook-type pre-commit --hook-type commit-msg
```

> `prek` is managed by mise and will be available after `mise install`.

If you have changed something on the source you should simply commit following
the mentioned conventions:

```console
git checkout -b feat/new-feature
git add --all
git commit -m 'feat: added awesome new feature'
git push --set-upstream origin feat/new-feature
```

After pushing your changes into the Git repository you should create a pull
request on GitHub. If the pull request have been merged and everything built
fine it will also create automatically a new release at least once a week.

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2026 Cloudpunks GmbG <info@cloudpunks.de>
```

[releases]: https://github.com/cloudpunks/dashboardctl/releases
[ghcr]: https://github.com/cloudpunks/dashboardctl/pkgs/container/dashboardctl
[docs]: https://cloudpunks.github.io/dashboardctl/#getting-started
[gotask]: https://taskfile.dev/installation/
[mise]: https://mise.jdx.dev/
[mise-install]: https://mise.jdx.dev/getting-started.html
[commits]: https://www.conventionalcommits.org/en/v1.0.0/
[semver]: https://semver.org/
