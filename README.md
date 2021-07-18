### 👋 Hi, I'm Paul and this is my CV.

It's written in Markdown, compiled to HTML using my own [Markdown parser](https://github.com/notpaulmartin/mdParser) and then saved to PDF.

**Why Markdown?**  
It is easy to read, yet very flexible.

**Compiling the CV**  
The content is in `cv.md` and the styling is in `styling.scss`.

1. Build executable:
    ```shell script
    go build -o buildCV cv
    ```
2. Run executable to compile `cv.md` to `cv_out.html` and `cv_out.pdf`:
    ```shell script
   ./buildCV
    ```
