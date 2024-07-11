# Book to Kindle

Book to Kindle is a simple command-line tool written in Go to send books to your Kindle device via email. It allows you to easily email books stored on your computer to your Kindle's email address.

## Prerequisites

Before using Book to Kindle, ensure you have the following:

- Go installed on your machine
- An active internet connection
- A Kindle device registered with your Amazon account
- The email address associated with your Kindle device (ending in @kindle.com)

## Installation

1. **Clone this repository**
```bash
git clone git@github.com:davidspader/book-to-kindle.git
```

2. **Navigate to the project directory**
```bash
cd book-to-kindle
```

3. **Install dependencies**
```bash
go mod tidy
```


4. **Set up the required environment variables by creating a .env file in the project root directory using the .env-template file, with the following content:**

- EMAIL=your_email@gmail.com
- PASSWORD=your_email_app_password
- TO_EMAIL=your_kindle_email@kindle.com
- BOOKS_DIR=/path/to/your/books/directory

Replace `your_email@gmail.com` with your Gmail address, `your_email_app_password` generate a new application-specific password in your Google account, and `your_kindle_email@kindle.com` with your Kindle email address. Additionally, set the `BOOKS_DIR` variable to the path where your books are stored.

## Usage

To send a book to your Kindle, run the following command:
```bash
go run main.go "book_name.pdf"
```

Replace `book_name.pdf` with the name of the book file you want to send. Make sure to enclose the file name in double quotes if it contains spaces.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.