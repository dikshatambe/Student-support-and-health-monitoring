import { Component, OnInit } from '@angular/core';

import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { DataService } from '../data.service';
import { Router } from "@angular/router";
import { MatSnackBar } from '@angular/material/snack-bar';
@Component({
  selector: 'app-health',
  templateUrl: './health.component.html',
  styleUrls: ['./health.component.css']
})
export class HealthComponent implements OnInit {


  constructor(
    private _snackBar: MatSnackBar,
    private formBuilder: FormBuilder,
    private dataService: DataService,
    private router: Router,
  ) { }
  registerForm: FormGroup;
  id = new FormControl('', [Validators.required, Validators.maxLength(4)]);
  age = new FormControl('', [Validators.required, Validators.maxLength(3)]);
  temp = new FormControl('', [Validators.required, Validators.maxLength(3)]);
  medical_history = new FormControl('', [Validators.required, Validators.maxLength(200)]);
  sour_throat = new FormControl('', [Validators.required, Validators.maxLength(30)]);


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

      id: this.id,
      age: this.age,
      temp: this.temp,
      medical_history: this.medical_history,
      sour_throat: this.sour_throat

    },

    );
  }



  onSubmit() {
    let userData = {
      "user_id": this.registerForm.value.id,
      "age": this.registerForm.value.age,
      "temp": this.registerForm.value.temp,
      "medical_history": this.registerForm.value.medical_history,
      "sour_throat": this.registerForm.value.sour_throat === "false" ? false:true

    };
    if (this.registerForm.invalid) {

      return;
    }
    console.log(userData)

    if (userData.temp > 100) {
      alert("High risk. Contact helpline for help");
    }

    this.dataService.insertUserHealth(userData).subscribe(data => {
      this.openSnackBar("Submitted Successfully", " ")
    })
    this.registerForm.reset();
  }

  cancel() {
    this.router.navigate(["home"])
  }
}
