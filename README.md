# Lincon

Lincon is a command-line tool for converting absolute hyperlinks to relative hyperlinks within a file or directory.

# Status

Lincon can be considered early alpha and is not ready for production use.

# Usage

**Convert all files in a folder**

```
lincon -path=FOLDER
```

**Convert an individual file**

To convert an individual file, you must also supply the path to the website's base (root) folder.

```
lincon -path=FILE -base=ROOT
```

## License

Lincon is published under the GNU General Public License, version 2. See [LICENSE](LICENSE) for more information.
