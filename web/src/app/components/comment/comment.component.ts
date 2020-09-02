import {Component, Input, OnInit} from '@angular/core';
import { Comment } from "../../models/comment";
import {Router} from "@angular/router";

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.css']
})
export class CommentComponent implements OnInit {

  @Input() comment: Comment

  constructor(
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  onClick() {
    this.router.navigate(['/comment', this.comment.Id], );
  }
}
