# PDF Generator

The PDF Generator is a project that allows you to generate PDF files based on HTML templates. It utilizes the Go programming language and relies on the `wkhtmltopdf` tool to convert HTML to PDF. With this tool, you can easily create PDF documents from HTML templates and save the output files in the "tmp" directory.

## Prerequisites

Before using this project, make sure you have the following installed:

1. Go: The Go programming language. You can download it from the official website: [https://golang.org/](https://golang.org/)

2. `wkhtmltopdf`: The binary from [https://wkhtmltopdf.org/](https://wkhtmltopdf.org/). This tool is responsible for converting the HTML template to a PDF document.

## Installation

1. Install Go by following the instructions on the official website.

2. Download the `wkhtmltopdf` binary suitable for your operating system from [https://wkhtmltopdf.org/](https://wkhtmltopdf.org/) and install it. Make sure it's accessible via the command line by adding it to your system's PATH.

## Getting Started

Follow these steps to generate PDFs using the PDF Generator:

1. Clone the repository:

```bash
git clone https://github.com/eduardylopes/pdf-generator.git
cd pdf-generator
```

1. Prepare your HTML templates:

Place your HTML templates inside the "templates" folder. These templates will be used to generate the content of the PDF files.

2. Generate a PDF:

Use the following command to generate a PDF based on your HTML template:

```bash
go run main.go
```

Replace your_template.html with the filename of your HTML template, and the output.pdf will be saved in the "tmp" directory.

1. Customize the output:

You can customize the PDF output further by specifying additional parameters supported by wkhtmltopdf. For example, you can set the page size, margins, header, footer, and more. To learn more about these options, refer to the wkhtmltopdf documentation.
