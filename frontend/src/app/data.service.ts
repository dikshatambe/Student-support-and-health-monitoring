import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams, HttpHeaders } from "@angular/common/http";

import { throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class DataService {

  private REST_API_SERVER = "http://localhost:4200/webapi/v1/user";
  private apiURL: string = 'http://localhost:4200/webapi/v1/file';
  private URL: string = 'http://localhost:4200/webapi/v1/userquery';
  private APIURL: string = 'http://localhost:4200/webapi/v1/userhealth';

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
  public getById(id: number) {
    return this.httpClient.get(this.REST_API_SERVER + "/" + id);
  }
  public getUserById(id: number) {
    return this.httpClient.get(this.REST_API_SERVER + "/" + id);   //endpoint
  }
  public getByName(firstname) {
    return this.httpClient.get(this.REST_API_SERVER + "/" + firstname).pipe(catchError(this.handleError));
  }

  public getBypassword(password) {
    return this.httpClient.get(this.REST_API_SERVER + "/" + password).pipe(catchError(this.handleError));
  }
  public insertUser(data) {
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.post(this.REST_API_SERVER, data, options).pipe(catchError(this.handleError));
  }
  public deleteProduct(id) {
    console.log("delete Product id : " + id);
    return this.httpClient.delete(this.REST_API_SERVER + "/" + id).pipe(catchError(this.handleError));
  }
  public updateById(data, id, firsttname, lastname, email, contactnumber) {
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
  public getUsers() {
    return this.httpClient.get(this.REST_API_SERVER).pipe(retry(3), catchError(this.handleError));
  }

  //services for User Query
  public getQueryById(id: number) {
    let httpHeaders = new HttpHeaders({
      'Access-Control-Allow-Origin': '*'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.get(this.URL + "/" + id, options);   //endpoint
  }

  public insertUserQuery(data) {

    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.post(this.URL, data, options).pipe(catchError(this.handleError));
  }

  public insertUserHealth(data) {
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Cache-Control': 'no-cache'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.post(this.APIURL, data, options).pipe(catchError(this.handleError));
  }
  public sendQueryRequest() {
    return this.httpClient.get(this.URL).pipe(retry(3), catchError(this.handleError));
  }

  public getUserQueries() {
    let httpHeaders = new HttpHeaders({
      'Access-Control-Allow-Origin': '*'
    });
    let options = {
      headers: httpHeaders
    };
    return this.httpClient.get(this.URL).pipe(retry(3), catchError(this.handleError));
  }

  //services for user health
  public getUserHealth() {
    return this.httpClient.get(this.APIURL).pipe(retry(3), catchError(this.handleError));
  }

  public getHealthById(id: number) {
    return this.httpClient.get(this.APIURL + "/" + id);
  }

  public getHealthOfUserById(id: number) {
    return this.httpClient.get("alldetailsofuser" + "/" + id);
  }
}
