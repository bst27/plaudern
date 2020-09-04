import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {CommentsComponent} from "./components/comments/comments.component";
import {CommentDetailsCardComponent} from "./components/comment-details-card/comment-details-card.component";

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/comments' },
  { path: 'comments', component: CommentsComponent },
  { path: 'comments/:commentId', component: CommentDetailsCardComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
