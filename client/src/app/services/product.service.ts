import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { Product } from '../shared/entities/product.entity';

@Injectable()
export class ProductService {

  private BASE_URL: string = "http://localhost:8000/api/products/"

  constructor(private http: HttpClient) { }

  getProducts(): Observable<Product[]> {
    return this.http.get<Product[]>('/api/products');
  }

  countProducts(): Observable<number> {
    return this.http.get<number>('/api/products/count');
  }

  addProduct(product: Product): Observable<Product> {
    return this.http.post<Product>('/api/products', product);
  }

  getProduct(product: Product): Observable<Product> {
    return this.http.get<Product>(`/api/products/${product.id}`);
  }

  editProduct(product: Product): Observable<any> {
    return this.http.put(`/api/products/${product.id}`, product, { responseType: 'text' });
  }

  deleteProduct(product: Product): Observable<any> {
    return this.http.delete(`/api/products/${product.id}`, { responseType: 'text' });
  }
}
