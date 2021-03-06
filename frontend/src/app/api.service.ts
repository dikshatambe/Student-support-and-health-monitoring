import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams, HttpHeaders } from "@angular/common/http";

import { throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})

export class ApiService {

  private REST_API_SERVER = "http://localhost:4200/webapi/v1/group";

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
  public getGroupById(id: number) {
    return this.httpClient.get(this.REST_API_SERVER +"/"+ id);   //endpoint
  }
  
  
  public getBypassword(password) {
    return this.httpClient.get(this.REST_API_SERVER +"/"+ password).pipe(catchError(this.handleError));
  }
  public createGroup(data) {
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.post(this.REST_API_SERVER, data, options).pipe(catchError(this.handleError));
  }
  public deleteGroup(id) {
    console.log("delete Product id : "+id);
    return this.httpClient.delete(this.REST_API_SERVER + "/" + id).pipe(catchError(this.handleError));
  }
  public updateGroup(data, id,firstname) {
    console.log("called update")
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.put(this.REST_API_SERVER + "/" + id, data, options).pipe(catchError(this.handleError));
  }

  public getGroup() {
    return this.httpClient.get(this.REST_API_SERVER).pipe(retry(3), catchError(this.handleError));
  }
}
