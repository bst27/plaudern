import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {CommentsComponent} from './components/comments/comments.component';
import {CommentDetailsCardComponent} from './components/comment-details-card/comment-details-card.component';
import {AuthGuard} from './services/auth.guard';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/comments' },
  { path: 'comments', component: CommentsComponent, canActivate: [AuthGuard] },
  { path: 'comments/:commentId', component: CommentDetailsCardComponent },
  { path: '**', redirectTo: '/comments' }// Fallback: redirect to homepage if route does not exist
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
