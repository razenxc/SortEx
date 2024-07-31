# Technical Documentation of "Sort File With EXIF data" cli program

Idea - sort files with EXIF data by dirs with date

Example:
```
1.png		23_01_2020
2.png	-> 	02_06_2023
3.png 		13_08_2024
```

## Program should:
- be cli ✅

`sortex -h` 

`sortex --help`

- sort files in current dir ✅

`sortex -c` 

`sortex --current`

- sort files in determined dir ✅

`sortex -d` 

`sortex --determined`

- sort files from dirs in current dirrectory ❌

`sortex -c -i` 

`sortex --internal-dirs`
- select the path where files will be sorted ✅

`sortex -c -s "/home/sort/here"`

`sortex --current --internal-dirs --save "/home/sort/here"`

## Also
- set dir date format ❌

`-sdf "mm:dd:yyyy"`

`--set-dir-format`

- should files without exif data be transferred? (`true` by default) ✅

`-m`

`--move-no-data false`

- restore a backup

`-b "/path/to/dir/with/a/backup"`

`--restore-backup "/path/to/dir/with/a/backup"`