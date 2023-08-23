# Duplicate files finder

Fast and simple command line tool, with caching capabilities, to find and clean duplicated files.

## Usage

1. Search for duplicated files 

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

2. Edit the output fiile and add a d to the braket of the files you want to delete

Open the outpute text file and you will see the list of duplicated files such as:
```
Duplicate 0 Size 732631040 Files 2 MD5 8cb9641aaaaaa0dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[]    '/files/a/filea.avi'
[]    '/files/b/filea.avi'
[]    '/files/a/filec.avi'
```

Add a `[d]` to the files you want to delete, for example:
```
Duplicate 0 Size 732631040 Files 2 MD5 8cb9641aaaaaa0dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[]    '/files/a/filea.avi'
[d]    '/files/b/filea.avi'
[d]    '/files/a/filec.avi'
```

Run the clean command on the file:
```
dupfinder clean duplicates.txt
```

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
