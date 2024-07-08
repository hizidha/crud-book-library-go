# CRUD Book Library

## Installation Go

Make sure Go is installed on your computer before running this project.
To install Go, visit the [official Go website](https://golang.org/) and follow the installation instructions.

## Running a Project

Here are the steps to get your Go project up and running:

#### 1. Clone Repository

```bash
git clone https://github.com/hizidha/crud-book-library-go.git
cd crud-book-library-go-master
```

#### 2. Install Dependencies
Initialize Go modules and tidy up dependencies:
```bash
go mod init crud-book-library
go mod tidy
```

#### 3. Complete All Data Requirements in ``.env``
```bash
DB_HOST=localhost
DB_PORT=your_port
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
```

#### 4. Enable PostgreSQL database on your Device
Ensure you have PostgreSQL installed and running. Create a database called books in your database.
To install PostgreSQL, visit the [official Go website](https://www.postgresql.org/download/) and follow the installation instructions.
```sql
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) NOT NULL,
    author VARCHAR(255),
    publisher VARCHAR(255),
    year INTEGER,
    category VARCHAR(100),
    location VARCHAR(100),
    eksemplar INTEGER
);
```

#### 5. Execute the Project
Run the following command to start the server:
```bash
go run .
```

#### 6. Access the Link Shortener
Open your web browser and navigate to:
```bash
http://localhost:8080/
```