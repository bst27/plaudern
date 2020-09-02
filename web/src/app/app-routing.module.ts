import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {CommentDetailsComponent} from "./components/comment-details/comment-details.component";
import {CommentsComponent} from "./components/comments/comments.component";

const routes: Routes = [
  { path: '', component: CommentsComponent },
  { path: 'comment/:commentId', component: CommentDetailsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
