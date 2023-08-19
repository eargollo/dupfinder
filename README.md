# Duplicate files finder

Fast and simple command line tool, with caching capabilities, to find and clean duplicated files.

## Usage

Finding duplicated files at a path placing result at an output file:
```
dupfinder scan /path/to/scan > duplicates.txt
```

Finding duplicated files at multiple paths:
```
dupfinder scan /path/to/scan /another/path > duplicates.txt
```

First execution might take a lot of time given files might need to be read. From second execution
on, most file signatures will already be cached.

The scan shows an execution log and comparisson is prioritized by file size.

### Managing the cache

Cache summary:
```
dupfinder cache
```

Listing cached items:

```
dupfinder cache --verbose
```

Deleting selected cached items (will delete every item started with the prefixes):
```
dupfinder cache --clean /my/path
```
