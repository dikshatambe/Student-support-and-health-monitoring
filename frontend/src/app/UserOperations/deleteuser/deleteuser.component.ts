import { Component, OnInit } from '@angular/core';
import { DataService } from '../../data.service';
import { UserloginService } from "../../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-deleteuser',
  templateUrl: './deleteuser.component.html',
  styleUrls: ['./deleteuser.component.css']
})
export class DeleteuserComponent implements OnInit {

  message:boolean;

  constructor(private _api : DataService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {
   
  }
  
  getid: any;
  isdel: Boolean = false;
  Confirmation: String = "User Deleted Successfully.";

  id(event:any){
    this.getid = event.target.value;
  }

  delete(){
    this._api.deleteProduct(this.getid).subscribe(res => {
      console.log(res);      
      this.isdel=true;
   })
  }


}

