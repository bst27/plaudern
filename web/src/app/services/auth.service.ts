import { Injectable } from '@angular/core';
import {BehaviorSubject, Observable, of} from 'rxjs';
import {delay} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private readonly loggedIn: BehaviorSubject<boolean>;

  constructor() {
    this.loggedIn = new BehaviorSubject<boolean>(false);
  }

  isLoggedIn(): Observable<boolean> {
    return this.loggedIn; // TODO: Handle authentication
  }

  login(password: string): Observable<void> {
    setTimeout(() => { this.loggedIn.next(true); }, 3000); // TODO: Handle authentication
    return of(undefined).pipe(delay(3000));
  }

  logout(): Observable<void> {
    setTimeout(() => { this.loggedIn.next(false); }, 3000); // TODO: Handle authentication
    return of(undefined).pipe(delay(3000));
  }
}
