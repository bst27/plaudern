import {Component, Input, OnInit} from '@angular/core';
import {Comment} from "../../models/comment";
import {primaryButtonClicked, secondaryButtonClicked} from "../../store/comments/comments.actions";
import {Store} from "@ngrx/store";
import {Router} from "@angular/router";

@Component({
  selector: 'app-comment-card',
  templateUrl: './comment-card.component.html',
  styleUrls: ['./comment-card.component.css']
})
export class CommentCardComponent implements OnInit {

  @Input() comment: Comment

  constructor(
    private store: Store,
    private router: Router
  ) {}

  ngOnInit(): void {
  }

  onPrimary() {
    this.store.dispatch(primaryButtonClicked({ comment: this.comment }));
  }

  onSecondary() {
    this.store.dispatch(secondaryButtonClicked({ comment: this.comment }));
    this.router.navigate(['/comments', this.comment.Id], );
  }


}
