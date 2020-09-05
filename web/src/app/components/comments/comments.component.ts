import { Component, OnInit } from '@angular/core';
import { Comment } from "../../models/comment";
import {select, Store} from "@ngrx/store";
import {Observable} from "rxjs";
import {loadComments} from "../../store/comments/comments.actions";
import {State} from "../../store/state";

@Component({
  selector: 'app-comments',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.css']
})
export class CommentsComponent implements OnInit {

  comments$: Observable<Comment[]>;

  constructor(
    private store: Store<State>,
  ) {
    this.comments$ = store.pipe(select(state => state.comments.comments));
  }

  ngOnInit(): void {
    this.store.dispatch(loadComments());
    // TODO: Sort comments (newest first)
  }

}
