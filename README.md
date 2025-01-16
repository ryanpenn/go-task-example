# Go Task Example

This is a very simple [Go web app](https://github.com/go-task/examples) that you can use as an example on how to setup
Task as part of the build pipeline of a Go app.

Features:

- Build everything with a single command;
- It's fast since it just run the necessary steps;
- Allow live-reloading using the `--watch` flag.
- Custom commands can be added to the `Taskfile.yml` file.

How to use:

- This requires [Go][go], [esbuild][esbuild] to be installed;
- [Install Task][installtask];
- Run `task install` to install the tools required to build this app;
- Run `task` to build and run the web app;
- Optionally you can use `task --watch` (or `task -w`) to re-compile and re-run
  the app on any file change.
- Open http://localhost:8383

[go]: https://golang.org/
[installtask]: https://taskfile.dev/installation/
[doc]: https://taskfile.dev/
[doc-zh]: https://task-zh.readthedocs.io/
[esbuild]: https://esbuild.github.io/
[setup-task]: https://github.com/arduino/setup-task
