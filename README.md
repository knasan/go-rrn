# RRN - Recursive Rename (rewritten)

## Description

[de] RRN - Recursive Rename ist ein kleines Kommandozeilen-Tool, dass Verzeichnisse und Dateien durchsucht,
und durch das angegebene Zeichen ersetzt.

[en] RRN - Recursive Rename is a small command line tool that searches directories and files and replaces them
with the given character.

## Dependencies

go and make for building from source, upx for binary compression when using my Makefile.

## Build from source

```shell
git clone https://github.com/knasan/go-rrn.git
cd go-rrn
make
```

[de] Um alle Leerzeichen durch einen unterstrich zu ersetzten reicht ein einfacher Aufruf.

[en] To replace all spaces with an underscore, simply call.

`rrn -s ' ' -r '_' -d testdir`

[de] Ich Empfehle davor ein Test druchzuführen (-D), um zu sehen was rrn machen würde übergebe den Schalter (-v).

[en] I recommend taking a test (-D) to see what rrn would do by giving the switch (-v).

## TODO

- [ ] files and/or directories can optionally be excluded
- [x] more than one directory depth (and with limit)
- [x] interactive - Request confirmation before attempting to rename any directory or file (since version 0.0.4)
