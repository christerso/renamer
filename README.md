# Recursive Image Renamer
This program is a simple command-line tool written in Go that renames image files in a specified directory and its subdirectories. The new file names contain information such as the directory name, date, image dimensions, aspect ratio, and a CRC32 checksum. This is useful for organizing and identifying image files more efficiently.

## Format
The new name format for image files is as follows:

currentdirectoryname-date-widthxheight-aspectratio-crc.xxx

where:

currentdirectoryname is the name of the directory containing the image file
date is the current date in the *YYYYMMDD* format
width and height are the image dimensions
aspectratio is the image aspect ratio with underscores instead of colons (e.g., 16_9 for 16:9)
crc is the CRC32 checksum of the image file
xxx is the original file extension
Usage
To use this program, first compile it with the following command:

```
go build main.go
Then, run the compiled executable with the -dir flag to specify the directory you wish to process:
```

```
./renamer -dir /path/to/directory
```
By default, the program processes the current directory if no -dir flag is provided.

Example
Suppose you have an image file named example.jpg in a directory called images. After running the program, the image file would be renamed to something like:

images-20230410-1920x1080-16_9-5a8f6d23.jpg

Dependencies
This program uses the Go standard library and does not require any additional dependencies.

License
This program is available under the MIT License.