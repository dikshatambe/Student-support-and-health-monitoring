import { Component, OnInit } from '@angular/core';
import { MustMatch } from './must-match.validator';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { DataService } from '../data.service';
import { Router } from "@angular/router";
import { MatSnackBar } from '@angular/material/snack-bar';
@Component({
  selector: 'app-register-user',
  templateUrl: './register-user.component.html',
  styleUrls: ['./register-user.component.css']
})
export class RegisterUserComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private formBuilder: FormBuilder,
    private dataService: DataService,
    private router: Router,
  ) { }
  registerForm: FormGroup;

  firstname = new FormControl('', [Validators.required, Validators.maxLength(100)]);
  lastname = new FormControl('', [Validators.required, Validators.maxLength(100)]);
  batch = new FormControl('', [Validators.required, Validators.maxLength(4)]);
  email = new FormControl('', [Validators.required, Validators.email, Validators.maxLength(100)]);
  contactNumber = new FormControl('', [Validators.required, Validators.pattern(new RegExp("[0-9 ]{10}"))])
  address = new FormControl('', [Validators.required, Validators.maxLength(50)]);
  password = new FormControl('', [Validators.required, Validators.minLength(6), Validators.maxLength(10)]);
  confirmPassword = new FormControl('', Validators.required);

  ngOnInit() {
    this.createFormValidations();

  }
  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 6000,
      verticalPosition: 'top'
    });
  }
  createFormValidations() {
    this.registerForm = this.formBuilder.group({

      firstname: this.firstname,
      lastname: this.lastname,
      batch: this.batch,
      email: this.email,
      contactNumber: this.contactNumber,
      address: this.address,
      password: this.password,
      confirmPassword: this.confirmPassword
    },
      { validator: MustMatch('password', 'confirmPassword') }
    );
  }

  passwordMatchValidator(group: FormGroup) {
    return group.get('password').value === group.get('confirmPassword').value ? null : { misMatch: true };
  }

  onSubmit() {
    let userData = {

      "first_name": this.registerForm.value.firstname,
      "last_name": this.registerForm.value.lastname,
      "batch": this.registerForm.value.batch_name,
      "email": this.registerForm.value.email,
      "password": this.registerForm.value.password,
      "address": this.registerForm.value.address,
      "contact_number": this.registerForm.value.contactNumber
    };
    if (this.registerForm.invalid) {

      return;
    }
    console.log(userData)
    this.dataService.insertUser(userData).subscribe(data => {
      this.openSnackBar("Registered Successfully", " ")
    })
    this.registerForm.reset();
  }

  cancel() {
    this.router.navigate(["login"])
  }
}
