import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {CommentsComponent} from './components/comments/comments.component';
import {CommentDetailsCardComponent} from './components/comment-details-card/comment-details-card.component';
import {AuthGuard} from './services/auth.guard';
import {LoginComponent} from './components/login/login/login.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/comments' },
  { path: 'login', component: LoginComponent },
  { path: 'comments', component: CommentsComponent, canActivate: [AuthGuard] },
  { path: 'comments/:commentId', component: CommentDetailsCardComponent, canActivate: [AuthGuard] },
  { path: '**', redirectTo: '/comments' }// Fallback: redirect to homepage if route does not exist
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
