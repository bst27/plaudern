import { Component, OnInit } from '@angular/core';
import { Comment } from "../../models/comment";

@Component({
  selector: 'app-comments',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.css']
})
export class CommentsComponent implements OnInit {

  comments: Comment[] = [];

  constructor() { }

  ngOnInit(): void {
    let data = [
      {
        Author: 'John',
        Created: '2020-08-31T21:54:21+02:00',
        Id: 'ba2d2e6a-2299-40fb-94d7-2738a76333c4',
        Message: 'Hi <br>there',
        ThreadId: 'localhost8083/basic/basic-example.html',
      },
      {
        Author: 'Max',
        Created: '2020-08-31T21:57:21+02:00',
        Id: '11111-2299-40fb-94d7-2732222223c4',
        Message: 'Thanks John!',
        ThreadId: 'localhost8083/basic/basic-example.html',
      },
    ];

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
  }

}
