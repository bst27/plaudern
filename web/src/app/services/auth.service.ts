import { Injectable } from '@angular/core';
import {Observable, of} from 'rxjs';
import {delay} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  isLoggedIn(): Observable<boolean> {
    return of(true).pipe(delay(3000)); // TODO: Handle authentication
  }
}
