# Pandoc filter for Excel table

## Usage 

Include file as below in the Markdown file:

    ~~~~
    {aetable title="test title" file="test/file.xlsx" sheet="Sheet1"}
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Then use the filter:

    ```
        pandoc -o test.docx mytestc.md --filter filters/aefilters
    ```

## Others

It is a very simple version, I will fix it if I really need. Thanks [Go Pandoc filters](https://github.com/oltolm/go-pandocfilters). 


### Author

[rocksun](https://github.com/rocksun/)