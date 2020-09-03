import {Component, Input, OnInit} from '@angular/core';
import {Comment} from "../../models/comment";
import {Router} from "@angular/router";

@Component({
  selector: 'app-comment-card',
  templateUrl: './comment-card.component.html',
  styleUrls: ['./comment-card.component.css']
})
export class CommentCardComponent implements OnInit {

  @Input() comment: Comment

  constructor(
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  onPrimary() {

  }

  onSecondary() {
    this.router.navigate(['/comment', this.comment.Id], );
  }


}
