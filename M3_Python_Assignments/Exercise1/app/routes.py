from flask import Blueprint, request, jsonify
from app.db import db
from app.models import Book

def error_response(status_code, error, message):
    return jsonify({"error": error, "message": message}), status_code

book_api = Blueprint('api', __name__)

genre_list = ['Fiction', 'Non-Fiction', 'Science Fiction', 'Graphic Novel', 'Comedy']

# Create
@book_api.route('/books', methods=['POST'])
def create_book():
    data = request.json
    if 'title' not in data or 'author' not in data or 'published_year' not in data or 'genre' not in data:
        return error_response(400, "Bad Request", "Missing required fields")
    if data['genre'] not in genre_list:
        return error_response(400, "Bad Request", "Invalid genre")
    if data['published_year'] < 0:
        return error_response(400, "Bad Request", "Invalid published year")
    new_book = Book(title=data['title'], author=data['author'], published_year=data['published_year'], genre=data['genre'])
    db.session.add(new_book)
    db.session.commit()
    return jsonify({"message": "Book created", "book": {"id": new_book.id, "title": new_book.title, "author": new_book.author, "published_year": new_book.published_year, "genre": new_book.genre}}), 201

# Read all
@book_api.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    books_list = [{"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre} for book in books]
    return jsonify(books_list)

# Read one
@book_api.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")
    return jsonify({"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre})

# Update
@book_api.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")
    data = request.json
    if data['genre'] not in genre_list:
        return error_response(400, "Bad Request", "Invalid genre")
    if data['published_year'] < 0:
        return error_response(400, "Bad Request", "Invalid published year")
    book.title = data['title']
    book.author = data['author']
    book.published_year = data['published_year']
    book.genre = data['genre']
    db.session.commit()
    return jsonify({"message": "Book updated", "book": {"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre}})

# Delete
@book_api.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")
    db.session.delete(book)
    db.session.commit()
    return jsonify({"message": "Book deleted"})

@book_api.route('/books/genres', methods=['GET'])
def get_by_genre():
    data = request.json
    if 'genre' not in data:
        return error_response(400, "Bad Request", "Missing required fields")
    if data['genre'] not in genre_list:
        return error_response(400, "Bad Request", "Invalid genre")
    books = Book.query.filter_by(genre=data['genre']).all()
    books_list = [{"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre} for book in books]
    return jsonify(books_list)

@book_api.route('/books/authors', methods=['GET'])
def get_by_author():
    data = request.json
    if 'author' not in data:
        return error_response(400, "Bad Request", "Missing required fields")
    books = Book.query.filter_by(author=data['author']).all()
    books_list = [{"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre} for book in books]
    return jsonify(books_list)