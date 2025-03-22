# SortEx
## Sort your files with EXIF data by date on directories
### [Technical Documentation](TD.md)

```
photo1.jpg              2024_05_14.jpg
photo2.png      to =>   2024_04_26.png
photo3.jpeg             2023_12_05.jpeg
```

## Running
- run it with **go** from terminal
    - <code>go run .</code>
- compile it with **go** from terminal
    - <code>mkdir build</code>
    - <code>go build -o build</code>
- download it with **releases page**
    - [**Here**](https://github.com/razenxc/SortFilesWithEXIFdata/releases)

## Usage 
\* colors means compatible arguments with each other - oranges with oranges, yellows with yellows, red and green only themself
- help
`-h --help`🔴
- sort files in current dir
`-c --current`🟠
- sort files in specific dir
`-d --determined "/home/user/media"`🟡
- select the path where the files will be sorted
`-s --save "/home/user/sorthere` 🟠🟡
- should files without data be transferred? (true by default)
`-m --move-no-data false` 🟠🟡
- undo changes if something went wrong
`-b --restore-backup "/home/user/mydir/31_7_2024-21_30_4-backup.sfbackup"`🟢

## The Screenshot
![image](https://github.com/user-attachments/assets/d698ad99-4d37-4c68-bd0c-0cdd54226c6e)

## Thirdparty
- [goexif](https://github.com/rwcarlsen/goexif)
