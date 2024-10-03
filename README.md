# 📚 Bookstore API 🛍️

Welcome to the **Bookstore API**! Here, Bookstore API (in GoLang using Gin framework): you can manage your literary desires, from adding, edit, delete, deactivate the books to updating your favorite library. 

```

├── bookstore/
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│   ├── .env
│   ├── books.db
│   ├── main.go
│   ├── models/
│   │   └── book.go
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css
│   ├── db/
│   │   ├── database.go
│   │   └── seed.go
│   ├── templates/
│   │   ├── home.html
│   │   ├── admin.html
│   │   └── layout.html
│   ├── controllers/
│   │   └── bookController.go
│   ├── routes/
│   │   └── routes.go

```


## 🚀 Getting Started

Follow these simple steps to get started:

1. **Clone this repository:**
   ```bash
   git clone https://github.com/yourusername/bookstore.git

2.	Navigate to the directory:


```
cd bookstore

```


3.	Install the required packages:
Make sure you have Go installed. If not, it’s time to go shopping! 🛒


```
go mod tidy

```


4.	Run the server:

```
go run main.go

```

5.	Open your browser and visit http://localhost:8080 to see the magic happen! 🎩✨

## 📚 API Endpoints

### 1. Get All Books 📖

- **Endpoint:** `GET /books`
- **Description:** Retrieves a list of all books. No, we don’t have *Harry Potter*… yet. 🧙‍♂️

---

### 2. Get a Specific Book 📘

- **Endpoint:** `GET /books/:id`
- **Description:** Retrieves a book by its ID. A great way to find that one book you loved but forgot the title of!

---

### 3. Create a New Book 🆕

- **Endpoint:** `POST /books`
- **Description:** Adds a new book to the database. Finally, a chance to write that bestseller you’ve always dreamed of!

---

### 4. Edit an Existing Book ✏️

- **Endpoint:** `PUT /books/:id`
- **Description:** Updates an existing book. Because sometimes, even books need a makeover! 💅

---

### 5. Delete a Book 🗑️

- **Endpoint:** `DELETE /books/:id`
- **Description:** Removes a book from the database. Out with the old, in with the new!

---

## 🔧 Requirements

- **Go:** Version 1.15 or later
- **Database:** A compatible database (like MySQL) because storing books in the cloud is so last year! ☁️📦

---

<img width="505" alt="image" src="https://github.com/user-attachments/assets/53c144d0-af7d-4d6a-88d5-f9fd7a4e8890">

