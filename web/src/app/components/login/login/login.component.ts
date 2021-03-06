import {Component, Input, OnInit} from '@angular/core';
import {AuthService} from '../../../services/auth.service';
import {MatDialog} from '@angular/material/dialog';
import {LoadingDialogComponent} from '../../loading-dialog/loading-dialog.component';
import {Router} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  isLoggedIn: boolean;
  hidePassword: boolean;
  @Input() password: string;

  constructor(
    private authService: AuthService,
    private loadingDialog: MatDialog,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.hidePassword = true;
    this.password = '';

    this.authService.isLoggedIn().subscribe(loggedIn => {
      this.isLoggedIn = loggedIn;
    });
  }

  login(): void {
    const dialog = this.loadingDialog.open(LoadingDialogComponent);

    this.authService.login(this.password).subscribe(loggedIn => {
      if (loggedIn) {
        this.password = '';

        // Navigate to homepage
        this.router.navigateByUrl('/');
      }

      dialog.close();
    });
  }

  logout(): void {
    const dialog = this.loadingDialog.open(LoadingDialogComponent);

    this.authService.logout().subscribe(() => {
      dialog.close();
    });
  }

}
