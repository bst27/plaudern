import { Injectable } from '@angular/core';
import {Observable, of} from "rxjs";
import {Comment} from "../models/comment";

@Injectable({
  providedIn: 'root'
})
export class CommentsService {

  constructor() { }

  getComments(): Observable<Comment[]> {
    return of([
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
    ]);
  }
}
