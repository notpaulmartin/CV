# Paul's CV

Written in Markdown, compiled to HTML using my own [Markdown parser](https://github.com/notpaulmartin/mdParser) and then saved to PDF.

## Compiling the CV
The content is in `cv.md` and the styling is in `styling.scss`

1. Build executable:
    ```shell script
    go build -o buildCV cv
    ```
2. Run executable to compile `cv.md` to `cv_out.html` and `cv_out.pdf`:
    ```shell script
   ./buildCV
    ```
