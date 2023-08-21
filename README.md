# Duplicate files finder

Fast and simple command line tool, with caching capabilities, to find and clean duplicated files.

## Usage

Finding duplicated files at a path. Results will be saved on `duplicates.txt`:
```
dupfinder scan /path/to/scan
```

Finding duplicated files at multiple paths. Results will be saved on `duplicates.txt`::
```
dupfinder scan /path/to/scan /another/path
```

First execution might take quite some time given potential duplicate files need to be loaded
for comparison. From second execution on, most file signatures will already be cached.

The scan shows an execution log. Output file will have the list of duplicates sorted by
the group size.

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
