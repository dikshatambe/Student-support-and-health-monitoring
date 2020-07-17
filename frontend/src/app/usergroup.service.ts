import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams, HttpHeaders } from "@angular/common/http";

import { throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UsergroupService {

  private REST_API_SERVER = "webapi/v1/";

  constructor(private httpClient: HttpClient) { }

  handleError(error: HttpErrorResponse) {
    let errorMessage = 'Unknown error!';
    if (error.error instanceof ErrorEvent) {
      // Client-side errors
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side errors
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    window.alert(errorMessage);
    return throwError(errorMessage);
  }
  public sendGetRequest() {
    return this.httpClient.get(this.REST_API_SERVER).pipe(retry(3), catchError(this.handleError));
}
public getById(id : number) {
    return this.httpClient.get(this.REST_API_SERVER +"/"+ id);
    }

public insertGroup(data) {
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.post(this.REST_API_SERVER, data, options).pipe(catchError(this.handleError));
  }
  createGroup(grupdata: { group_name: string; group_info: string }) {

    // return this.http.post(`${this.apiURL}group`,grupdata)
    return this.httpClient.post(this.REST_API_SERVER + 'group', JSON.stringify(grupdata))

  }
  getGroup() {
    return this.httpClient.get(this.REST_API_SERVER + 'group')
  }
  
  public addUserGroup(group_id, user_id) {
    let userGroup = { "group_id": parseInt(group_id), "user_id": parseInt(user_id) }
    return this.httpClient.post(this.REST_API_SERVER + 'usergroup', JSON.stringify(userGroup))
  }
  }
	  
