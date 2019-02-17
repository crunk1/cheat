# Cheat
CLI tool for viewing, editing, and deleting your cheatsheets.
Cheatsheets are stored in `~/.cheat/`.

## Usage
```
cheat <cheatsheet>
cheat -d <cheatsheet>
cheat -e <cheatsheet>
cheat -h
cheat -l

Cheatsheet manager. Displays, edits, and deletes cheatsheets.

Options:
  -d Deletes a cheatsheet.
  -e Edits a cheatsheet. If the cheatsheet does not exist, it is created.
  -h Displays this help message.
  -l Lists cheatsheets.

Examples:
  % cheat -e grep
  % cheat grep
```

## Building
To create a static, super slimmed down binary (I got it to ~600k),
you can use the following commands:

```
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' .
# or GOOS=darwin if building for mac

# And if you have upx (highly recommended):
upx ./cheat
```
