---
title: "Building"
date: 2022-05-03T00:00:00+00:00
anchor: "building"
weight: 30
---

We use [mise][mise] to manage all required tools and their versions. Install it
by following the [official installation instructions][mise-install], then run
the following commands inside the repository to activate mise and install all
tools defined in `mise.toml`:

{{< highlight txt >}}
mise trust
mise install
{{< / highlight >}}

Since all required commands ar part of our [go-task][gotask] taskfile the
commands you got to execute are quite simple:

{{< highlight txt >}}
git clone https://github.com/cloudpunks/dashboardctl.git
cd dashboardctl

task build
{{< / highlight >}}

Finally you should have the binary within the `bin/` folder now, give it a try
with `./bin/dashboardctl -h` to see all available options.

[mise]: https://mise.jdx.dev/
[mise-install]: https://mise.jdx.dev/getting-started.html
[gotask]: https://taskfile.dev/installation/
