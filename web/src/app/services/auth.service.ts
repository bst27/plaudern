import { Injectable } from '@angular/core';
import {BehaviorSubject, Observable, of} from 'rxjs';
import {delay, map, tap} from 'rxjs/operators';
import {BackendService} from './backend.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private readonly loggedIn: BehaviorSubject<boolean>;

  constructor(
    private backend: BackendService
  ) {
    this.loggedIn = new BehaviorSubject<boolean>(false);

    this.checkLogin().subscribe(loggedIn => {
      this.loggedIn.next(loggedIn);
    });
  }

  isLoggedIn(): Observable<boolean> {
    return this.loggedIn;
  }

  checkLogin(): Observable<boolean> {
    return this.backend.checkAuth().pipe(
      tap(loggedIn => { this.loggedIn.next(loggedIn); })
    );
  }

  login(password: string): Observable<boolean> {
    return this.backend.login(password).pipe(
      tap(loggedIn => { this.loggedIn.next(loggedIn); }),
    );
  }

  logout(): Observable<boolean> {
    return this.backend.logout().pipe(
      tap(loggedIn => { this.loggedIn.next(loggedIn); }),
    );
  }
}
