import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { Book } from '../shared/entities/book.entity';

@Injectable()
export class BookService {

  constructor(private http: HttpClient) { }

  getBooks(): Observable<Book[]> {
    return this.http.get<Book[]>('/api/books');
  }

  countBooks(): Observable<number> {
    return this.http.get<number>('/api/books/count');
  }

  addBook(book: Book): Observable<Book> {
    return this.http.post<Book>('/api/books', book);
  }

  getBook(book: Book): Observable<Book> {
    return this.http.get<Book>(`/api/books/${book.Id}`);
  }

  editBook(book: Book): Observable<any> {
    return this.http.put(`/api/books/${book.Id}`, book, { responseType: 'text' });
  }

  deleteBook(book: Book): Observable<any> {
    return this.http.delete(`/api/books/${book.Id}`, { responseType: 'text' });
  }
}
