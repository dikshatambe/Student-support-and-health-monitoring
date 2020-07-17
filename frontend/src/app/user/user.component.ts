import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {

  constructor(private formBuilder: FormBuilder,
    private dataService: DataService,
    private _snackBar: MatSnackBar,
    private router: Router, ) { }
  registerForm: FormGroup;
  //fid = new FormControl('', [Validators.required, Validators.pattern(new RegExp("[0-9 ]"))]);
  user_id = new FormControl('', [Validators.required, Validators.maxLength(20)]);
  query = new FormControl('', [Validators.required, Validators.maxLength(200)]);




  ngOnInit(): void { this.createFormValidations(); }
  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 6000,
      verticalPosition: 'top'
    });
  }

  createFormValidations() {
    this.registerForm = this.formBuilder.group({

      user_id: this.user_id,
      query: this.query
    },
    );
  }
  onSubmit() {
    let userData = {

      "user_id": this.registerForm.value.user_id,
      "Query": this.registerForm.value.query
    };
    if (this.registerForm.invalid) {

      return;
    }
   

    this.dataService.insertUserQuery(userData).subscribe(data => {
      this.openSnackBar("Query added Successfully", " ")
    })
    this.registerForm.reset();
  }

  Cancel() {
    this.router.navigate(["home"])
  }




}
