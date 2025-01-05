from ..models.models import Book, Customer, Transaction
from ..errors.error import *

books = []

def add_book(title, author, price, quantity):
    price, quantity = validateInput(price, quantity)
    if price is None and quantity is None:
        raise InvalidInputException()
    
    for book in books:
        if book.title.lower() == title.lower():
            raise DuplicateBookException()
        
    books.append(Book(title, author, price, quantity))
    return "Book added successfully!"

def view_books():
    if len(books)==0:
        print("No Books Available!\n")
    return [book.view_book_details() for book in books]

def search_book(query):
    result = [book.view_book_details() for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    if result:
        return result
    else:
        raise BookNotFoundException()

customers = []

def add_customer(name, email, phone):
    if not name or not email or not phone:
        raise AllFieldsRequiredException()
    
    for customer in customers:
        if customer.email.lower() == email.lower():
            raise DuplicateCustomerException()
        
    customers.append(Customer(name, email, phone))
    return "Customer added successfully!"

def view_customers():
    if len(customers)==0:
        print("No Customers Available!\n")
    return [customer.view_customer_details() for customer in customers]

sales = []

def sell_book(customer_name, email, phone, book_title, quantity):
    book = next((b for b in books if b.title.lower() == book_title.lower()), None)
    temp, quantity = validateInput(1,quantity)
    if quantity is None:
        raise InvalidInputException()
    if not book:
        raise BookNotFoundException()
    if book.quantity < quantity:
        raise OutOfStockException(book.quantity)
    
    customer = next((c for c in customers if c.email == email), None)
    if not customer:
        add_customer(customer_name, email, phone)
    
    book.quantity -= quantity
    sales.append(Transaction(customer_name, email, phone, book_title, quantity))
    return f"Sale successful! Remaining quantity: {book.quantity}"

def view_sales():
    if len(sales)==0:
        print("No Sales Available!\n")
    return [sale.view_transaction_details() for sale in sales]
