import { Component, OnInit } from '@angular/core';
import { DataService } from '../../data.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-updateuser',
  templateUrl: './updateuser.component.html',
  styleUrls: ['./updateuser.component.css']
})
export class UpdateuserComponent implements OnInit {

  message:boolean;

  constructor(private _api : DataService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
    
    
  }
  updateuser: Object  = {} ;
  isUpdated: Boolean = false;
  Confirmation: String = "User details has been Updated.";

  put = function(data){
    this.updateuser = {
      "id":data.id,
      "first_name":data.first_name,
      "last_name":data.last_name,
      "email":data.email,
      "password":data.password,
      "contact_number":data.contact_number
    }

    this._api.updateByIdss(this.updateuser).subscribe(res => {
       console.log(res);
       this.isUpdated=true;
       
    })
  }


}

