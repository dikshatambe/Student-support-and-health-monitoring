
import { Component, OnInit } from '@angular/core';
import { IsrootService } from "../isroot.service";
import { Router } from "@angular/router";

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {

  user_root: boolean;
  constructor(private _isroot: IsrootService,
    private router: Router, ) { }

  ngOnInit(): void {
    this._isroot.sharedMessage.subscribe(user_root => this.user_root = user_root)
  }


}

