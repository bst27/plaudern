import { Component, OnInit } from '@angular/core';
import { Comment } from "../../models/comment";
import {CommentsService} from "../../services/comments.service";

@Component({
  selector: 'app-comments',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.css']
})
export class CommentsComponent implements OnInit {

  comments: Comment[] = [];

  constructor(
    private commentsService: CommentsService
  ) { }

  ngOnInit(): void {
    this.commentsService.getComments().subscribe(data => {
      this.comments = data.sort((a, b) => {
        console.log(Date.parse(a.Created), Date.parse(b.Created));
        if (Date.parse(a.Created) < Date.parse(b.Created)) {
          return -1;
        }

        if (Date.parse(a.Created) > Date.parse(b.Created)) {
          return 1;
        }

        return 0;
      }).reverse();
    });
  }

}
