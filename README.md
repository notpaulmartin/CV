### 👋 Hi, I'm Paul and [this](https://raw.githubusercontent.com/notpaulmartin/CV/main/cv_out.pdf) is my CV.

Written in Markdown.  
Compiles to PDF with my own [Markdown parser](https://github.com/notpaulmartin/mdParser).

**Why Markdown?**  
It is easy to read and very flexible.

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
