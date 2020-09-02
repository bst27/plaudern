import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {CommentsService} from "../../services/comments.service";
import {Comment} from "../../models/comment";
import {first} from "rxjs/operators";

@Component({
  selector: 'app-comment-details',
  templateUrl: './comment-details.component.html',
  styleUrls: ['./comment-details.component.css']
})
export class CommentDetailsComponent implements OnInit {

  comment: Comment;

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

}
