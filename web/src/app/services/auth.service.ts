import { Injectable } from '@angular/core';
import {Observable, of} from 'rxjs';
import {delay} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private loggedIn: boolean;

  constructor() {
    this.loggedIn = false;
  }

  isLoggedIn(): Observable<boolean> {
    return of(this.loggedIn).pipe(delay(100)); // TODO: Handle authentication
  }

  login(): Observable<boolean> {
    this.loggedIn = true;
    return of(this.loggedIn).pipe(delay(3000)); // TODO: Handle authentication
  }

  logout(): Observable<boolean> {
    this.loggedIn = false;
    return of(this.loggedIn).pipe(delay(3000)); // TODO: Handle authentication
  }
}
