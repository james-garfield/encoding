# Encoding

I use Calibre often to convert EPUBs/PDFs/etc to MOBI files for my Kindle. I found that sometimes a EPUB can not be converted because of a "spine" not foud error.
Researching more about the issue I found that it was an encoding issue. When the files that make up the epub were not encoded in UTF-8, the EPUB could not be converted.
That's why I created this script.

## Usage

``` go run change_epub.go path/to/epub.epub ```

## Result

The result from running the above script is a UTF-8 encoded epub. Then the spine error goes away!