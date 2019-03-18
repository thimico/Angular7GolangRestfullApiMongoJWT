import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

import { BookService } from '../services/book.service';
import { ToastComponent } from '../shared/toast/toast.component';
import { Book } from '../shared/entities/book.entity';

@Component({
  selector: 'app-books',
  templateUrl: './books.component.html',
  styleUrls: ['./books.component.css']
})
export class BooksComponent implements OnInit {

  book = new Book();
  books: Book[] = [];
  isLoading = true;
  isEditing = false;

  addBookForm: FormGroup;
  Isbn = new FormControl('', Validators.required);
  Title = new FormControl('', Validators.required);
  Price = new FormControl('', Validators.required);

  constructor(private bookService: BookService,
              private formBuilder: FormBuilder,
              public toast: ToastComponent) { }

  ngOnInit() {
    this.getBooks();
    this.addBookForm = this.formBuilder.group({
      Isbn: this.Isbn,
      Title: this.Title,
      Price: this.Price
    });
  }

  getBooks() {
    this.bookService.getBooks().subscribe(
      data => !!data ? this.books = data : null,
      error => console.log(error),
      () => this.isLoading = false
    );
  }

  addBook() {
    this.bookService.addBook(this.addBookForm.value).subscribe(
      res => {
        this.books.push(res);
        this.addBookForm.reset();
        this.toast.setMessage('item added successfully.', 'success');
      },
      error => console.log(error)
    );
  }

  enableEditing(book: Book) {
    this.isEditing = true;
    this.book = book;
  }

  cancelEditing() {
    this.isEditing = false;
    this.book = new Book();
    this.toast.setMessage('item editing cancelled.', 'warning');
    // reload the books to reset the editing
    this.getBooks();
  }

  editBook(book: Book) {
    this.bookService.editBook(book).subscribe(
      () => {
        this.isEditing = false;
        this.book = book;
        this.toast.setMessage('item edited successfully.', 'success');
      },
      error => console.log(error)
    );
  }

  deleteBook(book: Book) {
    if (window.confirm('Are you sure you want to permanently delete this item?')) {
      this.bookService.deleteBook(book).subscribe(
        () => {
          const pos = this.books.map(elem => elem.Id).indexOf(book.Id);
          this.books.splice(pos, 1);
          this.toast.setMessage('item deleted successfully.', 'success');
        },
        error => console.log(error)
      );
    }
  }

}
