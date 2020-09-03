import {Component, Input, OnInit} from '@angular/core';
import {Comment} from "../../models/comment";
import {ActivatedRoute} from "@angular/router";
import {CommentsService} from "../../services/comments.service";

@Component({
  selector: 'app-comment-details-card',
  templateUrl: './comment-details-card.component.html',
  styleUrls: ['./comment-details-card.component.css']
})
export class CommentDetailsCardComponent implements OnInit {

  @Input() comment: Comment

  constructor(
    private route: ActivatedRoute,
    private commentsService: CommentsService
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      if (params.has('commentId')) {
        this.commentsService.getComments().subscribe(comments => {
          this.comment = comments.find(comment => comment.Id === params.get('commentId'))
        });
      }
    })
  }

  onPrimary() {
    console.log('!');
  }

  onSecondary() {
    console.log('?');
  }

}
