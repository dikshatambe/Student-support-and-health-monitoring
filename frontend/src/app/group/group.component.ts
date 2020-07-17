import { Component, OnInit ,Input} from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { GroupService } from '../group.service';
import { HttpClient, HttpErrorResponse, HttpParams, HttpHeaders } from "@angular/common/http";
import { Router } from "@angular/router";
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-group',
  templateUrl: './group.component.html',
  styleUrls: ['./group.component.css']
})
export class GroupComponent implements OnInit {

constructor(
	private router: Router,
    public authser:GroupService,
    private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    ) { }
    registerForm: FormGroup;
  group_name = new FormControl('', [Validators.required, Validators.maxLength(100)]);
  group_info = new FormControl('', [Validators.required, Validators.maxLength(100)]);
  

  ngOnInit(): void {
  this.createFormValidations();
  }
openSnackBar(message,action) {
    this._snackBar.open(message, action, {
      duration: 6000,
      verticalPosition: 'top'
    });
  }
  createFormValidations() {
    this.registerForm = this.formBuilder.group({

      group_name: this.group_name,
      group_info: this.group_info
    }
    );
  }

  onSubmit() {
    let userData = {
      "group_name": this.registerForm.value.group_name,
      "group_info": this.registerForm.value.group_info
         };
    if (this.registerForm.invalid) {

      return;
    }
    console.log(userData)
    this.authser.insertGroup(userData).subscribe(data => {
      this.openSnackBar("New Record Added Successfully"," ")
    })
    this.registerForm.reset();
  }



	  }
  
//  onSubmitCreate(){
  //this.authser.createGroup(this.grupdata).subscribe((data: {}) => {
   
    //console.log(data)
 // })
//}
//}
