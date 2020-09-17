import { Component, OnInit } from '@angular/core';
import {AuthService} from '../../../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  isLoggedIn: boolean;
  message: string;
  readonly: boolean;
  hidePassword: boolean;

  constructor(
    private authService: AuthService
  ) { }

  ngOnInit(): void {
    this.message = 'Checking status ...';
    this.readonly = true;
    this.hidePassword = true;

    this.authService.isLoggedIn().subscribe(loggedIn => {
      this.isLoggedIn = loggedIn;
      this.message = '';
      this.readonly = false;
    });
  }

  login(): void {
    this.message = 'Log in ...';
    this.readonly = true;

    this.authService.login().subscribe(() => {
      this.message = '';
      this.readonly = false;
    });
  }

  logout(): void {
    this.message = 'Log out ...';
    this.readonly = true;

    this.authService.logout().subscribe(() => {
      this.message = '';
      this.readonly = false;
    });
  }

}
