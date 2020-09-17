import { Component, OnInit } from '@angular/core';
import {AuthService} from '../../../services/auth.service';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  isLoggedIn: boolean;
  message: string;

  constructor(
    private authService: AuthService
  ) { }

  ngOnInit(): void {
    this.message = 'Checking status ...';

    this.authService.isLoggedIn().subscribe(loggedIn => {
      this.isLoggedIn = loggedIn;
      this.message = '';
    });
  }

  login(): void {
    this.message = 'Log in ...';

    this.authService.login().subscribe(loggedIn => {
      this.isLoggedIn = loggedIn;
      this.message = '';
    });
  }

  logout(): void {
    this.message = 'Log out ...';

    this.authService.logout().subscribe(loggedIn => {
      this.isLoggedIn = loggedIn;
      this.message = '';
    });
  }

}
