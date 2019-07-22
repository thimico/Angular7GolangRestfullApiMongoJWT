import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

import { ProductService } from '../services/product.service';
import { ToastComponent } from '../shared/toast/toast.component';
import { Product } from '../shared/entities/product.entity';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css']
})
export class ProductComponent implements OnInit {

  product = new Product();
  products: Product[] = [];
  isLoading = true;
  isEditing = false;

  addProductForm: FormGroup;
  Name = new FormControl('', Validators.required);
  Price = new FormControl('', Validators.required);

  constructor(private productService: ProductService,
              private formBuilder: FormBuilder,
              public toast: ToastComponent) { }

  ngOnInit() {
    this.getProducts();
    this.addProductForm = this.formBuilder.group({
      Name: this.Name,
      Price: this.Price
    });
  }

  getProducts() {
    this.productService.getProducts().subscribe(
      data => !!data ? this.products = data : null,
      error => console.log(error),
      () => this.isLoading = false
    );
  }

  addProduct() {
    this.productService.addProduct(this.addProductForm.value).subscribe(
      res => {
        this.products.push(res);
        this.addProductForm.reset();
        this.toast.setMessage('item added successfully.', 'success');
      },
      error => console.log(error)
    );
  }

  enableEditing(product: Product) {
    this.isEditing = true;
    this.product = product;
  }

  cancelEditing() {
    this.isEditing = false;
    this.product = new Product();
    this.toast.setMessage('item editing cancelled.', 'warning');
    // reload the products to reset the editing
    this.getProducts();
  }

  editProduct(product: Product) {
    console.log(product)
    this.productService.editProduct(product).subscribe(
      () => {
        this.isEditing = false;
        this.product = product;
        this.toast.setMessage('item edited successfully.', 'success');
      },
      error => console.log(error)
    );
  }

  deleteProduct(product: Product) {
    console.log(product)
    if (window.confirm('Are you sure you want to permanently delete this item?')) {
      this.productService.deleteProduct(product).subscribe(
        () => {
          const pos = this.products.map(elem => elem.Id).indexOf(product.Id);
          this.products.splice(pos, 1);
          this.toast.setMessage('item deleted successfully.', 'success');
        },
        error => console.log(error)
      );
    }
  }

}
