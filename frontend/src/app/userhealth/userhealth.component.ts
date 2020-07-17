import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { UserloginService } from "../userlogin.service";
import { Router } from '@angular/router';
@Component({
  selector: 'app-userhealth',
  templateUrl: './userhealth.component.html',
  styleUrls: ['./userhealth.component.css']
})
export class UserhealthComponent implements OnInit {

  constructor(private _api: DataService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {

  }
  byid = false;
  users: any;
  item: any[];
  u_id: any;
  u_age: any;
  u_temp: any;
  u_sour: boolean;
  u_history: string;


  getall() {           // for endpoint 
    this._api.getUserHealth().subscribe(res => {
      this.users = res["data"];
      console.log("Response", res);
      console.log(this.users)
    });
  }
  getid: any;
  id(event: any) {
    this.getid = event.target.value;
  }

  getbyid() {
    this.byid = true;
    console.log(this.getid);
    this._api.getHealthById(this.getid).subscribe(res => {
      this.item = res["data"]
      this.u_id = this.item["id"]
      this.u_age = this.item["age"]
      this.u_temp = this.item["temp"]
      this.u_sour = this.item["sour"]
      this.u_history = this.item["history"]
      if (this.u_temp > 100) {
        alert("High risk. Contact student");
      }
      console.log("Response", res);
      console.log("Items ", this.item);
      // console.log(this.users)


    });
  }
  Cancel() {
    this.router.navigate(["admin"])
  }
}

