import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { UserloginService } from "../userlogin.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-help',
  templateUrl: './help.component.html',
  styleUrls: ['./help.component.css']
})
export class HelpComponent implements OnInit {

  message: boolean;

  constructor(private _api: DataService, private value: UserloginService, private router: Router) { }

  ngOnInit(): void {

  }
  byid = false;
  users: any;
  item: any[];
  u_id: number;
  u_query: string;


  getall() {           // for endpoint 
    this._api.getUserQueries().subscribe(res => {
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
    this._api.getQueryById(this.getid).subscribe(res => {
      this.item = res["data"]
      this.u_id = this.item["id"]
      this.u_query = this.item["query"]

      console.log("Response", res);
      console.log("Items ", this.item);
      // console.log(this.users)


    });
  }
  Cancel() {
    this.router.navigate(["admin"])
  }
}

