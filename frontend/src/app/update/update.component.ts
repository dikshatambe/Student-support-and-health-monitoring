import { Component, OnInit, OnDestroy } from '@angular/core';
import { DataService } from '../data.service';
import { takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router, NavigationEnd } from "@angular/router";
import { HttpErrorResponse } from '@angular/common/http';
import { MatDialog } from '@angular/material/dialog';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent implements OnInit {

  

  constructor(private _snackBar: MatSnackBar,
    private dataService: DataService,
    public dialog: MatDialog,
    private router: Router) { }

    updateForm: FormGroup;

    navigationSubscription;
  users  = [];
    //data =[];
    destroy$: Subject<boolean> = new Subject<boolean>();
    

    openDialog(user,id,firstname,lastname,email,contactNumber) {

      id : user.data['id'],
      firstname="daud",
      lastname= user.data['last_name'],
      email= user.data['email'],
      contactNumber= user.data['contact_number'],
     // user.data['first_name']="madhav"
      //console.log("firstname : "+ user.data['first_name']);
      this.dataService.updateById(user,id,firstname,lastname,email,contactNumber).subscribe(data => {
        this.openSnackBar("Successfully Updated", " 🎉")
        //this.dialogRef.close();
      })

  }

  loadUsers() {
    this.dataService.sendGetRequest().pipe(takeUntil(this.destroy$)).subscribe((data: any[]) => {
      this.users = data;
    })
  }

  

  getById(id : number) {
    if (id < 0) {
      this.openSnackBar("Please Enter User id ","👤")
    }
    else {
      this.dataService.getById(id).subscribe((data: any[]) => {
        console.log(Object.values(data).length)
        var stringData = '[' + JSON.stringify(data) + ']'
        var parseData = JSON.parse(stringData)
        this.users = parseData;
        //console.log(data)
          this.openSnackBar("User Exist","😊")
         // this.loadUsers();
      },
        (err: HttpErrorResponse) => {
          if (err.error instanceof Error) {
            //A client-side or network error occurred.
            console.log('An error occurred:', err.error.message);
          } else {
            this.openSnackBar("User Not Exist", "😞")
           // this.loadUsers();

            //Backend returns unsuccessful response codes such as 404, 500 etc.
            console.log('Backend returned status code: ', err.status);

            console.log('Response body:', err.error);
          }
        }
      )
    }
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 4000,
      verticalPosition: 'top'
    });
  }

  ngOnInit(): void {
  }

  onSubmit() {
    let userData = {
      "first_name": this.updateForm.value.firstname,
      "last_name": this.updateForm.value.lastname,
      "email": this.updateForm.value.email,
      "password": this.updateForm.value.password,
      "contact_number": this.updateForm.value.contactNumber
    };
    console.log("hello"+this.updateForm.value.id)
      

    if (this.updateForm.invalid) {

      return;
    }
    console.log(userData)
    this.dataService.insertUser(userData).subscribe(data => {
      this.openSnackBar("New Record Added Successfully"," ")
    })
    this.updateForm.reset();
  }

}
