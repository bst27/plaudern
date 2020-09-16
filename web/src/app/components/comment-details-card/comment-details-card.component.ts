import {Component, Input, OnInit} from '@angular/core';
import {Comment} from '../../models/comment';
import {ActivatedRoute} from '@angular/router';
import {select, Store} from '@ngrx/store';
import {State} from '../../store/state';
import {loadComment} from '../../store/comments/comments.actions';

@Component({
  selector: 'app-comment-details-card',
  templateUrl: './comment-details-card.component.html',
  styleUrls: ['./comment-details-card.component.css']
})
export class CommentDetailsCardComponent implements OnInit {

  comment: Comment;

  constructor(
    private route: ActivatedRoute,
    private store: Store<State>,
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      if (params.has('commentId')) {
        this.store.dispatch(loadComment({ commentId: params.get('commentId') }));

        this.store.pipe(select((state) => {
          return state.comments.comments.find(comment => comment.Id === params.get('commentId'));
        })).subscribe(comment => {
          this.comment = comment;
        });
      }
    });
  }

}
