# Technical Documentation of "Sort Files with EXIF data" cli program

Problem: Once, I got the corrupted sd card in my camera, I was able to recover all the data, but all the photos were randomly "scattered" in one directory and with random names.

Solution: I decided to write a program that would sort these photos into directories separated by date.

Example:
```
1ABCD.png		23_01_2020.png
2BCDE.png	-> 	02_06_2023.png
3CDEF.png 		13_08_2024.png
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

`sortex -c` 

`sortex --internal-dirs`
- select the path where files will be sorted ✅

`sortex -s "/home/sort/here"`

`sortex --save "/home/sort/here"`

## Also
- set dir date format ❌

`-sdf "mm:dd:yyyy"`

`--set-dir-format`

- should files without exif data be transferred? (`true` by default) ✅

`-m`

`--move-no-data false`

- restore a backup ✅

`-b "/fullpath/to/a/backup/file"`

`--restore-backup "/fullpath/to/a/backup/file"`