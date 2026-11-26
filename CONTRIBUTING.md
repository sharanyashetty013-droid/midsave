# Contributing to midsave

Thanks for wanting to help! Small fixes are just as welcome as big features.

## Setup

You need Go 1.22 or newer.

```
git clone https://github.com/sharanyashetty013-droid/midsave.git
cd midsave
go build -o bin/midsave ./cmd/midsave
go build -o bin/writer ./examples/writer
go build -o bin/badwriter ./examples/badwriter
go build -o bin/check ./examples/check
```

On Windows, add `.exe` to each output name.

## Check that it works

```
bin/midsave --app bin/writer --check bin/check --clean data.txt --runs 5
bin/midsave --app bin/badwriter --check bin/check --clean data.txt --runs 10
```

The first should report `0 of 5 runs lost or broke data`. The second should catch broken runs.

## Making a change

1. Pick an issue, ideally one labelled `good first issue`, and comment that you're taking it.
2. Fork the repo and create a branch: `git checkout -b feat/short-name`
3. Make your change. Keep it small and focused on one thing.
4. Run `go vet ./...` and the two checks above.
5. Commit with a clear message, e.g. `add --json output`.
6. Push and open a pull request. Say what you changed and how you tested it.

CI runs on Linux and Windows for every pull request. It must pass before merging.

## Ideas welcome

Found a bug, or have an idea? Open an issue. You don't need to know how to fix it.

## Be kind

Be respectful and patient. Everyone here is learning.