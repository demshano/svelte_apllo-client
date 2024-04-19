// resolver.go

package graph

import (
	"context"
	"strconv"
)

type Book struct {
    ID     string
    Title  string
    Author string
}

type Resolver struct {
    books []*Book
}

func NewResolver() *Resolver {
    // Initialize resolver with some sample data
    return &Resolver{
        books: []*Book{
            {ID: "1", Title: "Book 1", Author: "Author 1"},
            {ID: "2", Title: "Book 2", Author: "Author 2"},
        },
    }
}

func (r *Resolver) Query_books(ctx context.Context) ([]*Book, error) {
    // Return list of books
    return r.books, nil
}

func (r *Resolver) Mutation_addBook(ctx context.Context, title string, author string) (*Book, error) {
    // Create a new book
    newBook := &Book{
        ID:     strconv.Itoa(len(r.books) + 1), // Generate unique ID
        Title:  title,
        Author: author,
    }

    // Append the new book to the list of books
    r.books = append(r.books, newBook)

    return newBook, nil
}
