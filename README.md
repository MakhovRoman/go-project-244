## Hexlet tests and linter status:
[![Actions Status](https://github.com/MakhovRoman/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/MakhovRoman/go-project-244/actions) [![Actions Status](https://github.com/MakhovRoman/go-project-244/actions/workflows/MY_CI.yml/badge.svg)](https://github.com/MakhovRoman/go-project-244/actions) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=MakhovRoman_go-project-244&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=MakhovRoman_go-project-244) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=MakhovRoman_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=MakhovRoman_go-project-244)

## Demo
[![asciicast](https://asciinema.org/a/OPuBxOgHoSJYBMKk.svg)](https://asciinema.org/a/OPuBxOgHoSJYBMKk?speed=2)

## Installation
```bash
git clone https://github.com/MakhovRoman/go-project-244
cd go-project-244
make build
``` 


## Usage
### Help

```bash
./bin/gendiff --help

NAME:
   gendiff - Compares two configuration files and shows a difference.

USAGE:
   gendiff [global options]

GLOBAL OPTIONS:
   --format string, -f string  output format (default: "stylish")
   --help, -h                  show help
```

### Example
```bash
./bin/gendiff ./testdata/fixture/file1.json ./testdata/fixture/file2.json --format plain
```
### Outputs
- **Stylish (default)**:
    ```
  {
      - follow: false
        host: hexlet.io
      - proxy: 123.234.53.22
      - timeout: 50
      + timeout: 20
      + verbose: true
    }
  ```
  
- **Plain**:
    ```
    Property 'follow' was removed
    Property 'proxy' was removed
    Property 'timeout' was updated. From 50 to 20
    Property 'verbose' was added with value: true
  ```
- **JSON**:
    ```
    {
      "follow": {
        "status": "removed",
        "oldValue": false
      },
      "host": {
        "status": "unchanged",
        "value": "hexlet.io"
      },
      "proxy": {
        "status": "removed",
        "oldValue": "123.234.53.22"
      },
      "timeout": {
        "status": "changed",
        "oldValue": 50,
        "newValue": 20
      },
      "verbose": {
        "status": "added",
        "newValue": true
      }
    }
    ```