# go-cooperhewitt-shoebox

Tools for archiving your Cooper Hewitt shoebox.

## Caveats

This is wet paint. It's probably too soon for you. Things this does not do yet:

* Archiving of anything but objects that have been collected
* Display any custom metadata (title, notes)
* Display any object/person/video metadata (besides title and accession number)
* Honour / display privacy settings
* Templates for generating HTML files
* Responsive hoohah
* Employ any kind of sane and modular code structure

## Setup

```
$> make build
```

## Tools

### shoebox

```
$> ./bin/shoebox -token <COOPERHEWITT_API_TOKEN> -shoebox <PATH_TO_LOCAL_SHOEBOX>
```

## See also

* https://collection.cooperhewitt.org/api/
* http://github.com/cooperhewitt/go-cooperhewitt-api
