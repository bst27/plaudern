import { Injectable } from '@angular/core';
import {Observable, of} from 'rxjs';
import {Comment} from '../models/comment';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';
import {CommentResponse} from '../models/comment-response';

@Injectable({
  providedIn: 'root'
})
export class CommentsService {

  constructor(
    private http: HttpClient
  ) { }

  getComments(): Observable<Comment[]> {
    return this.http.get<CommentResponse>(
      '/manage/comment'
    ).pipe(map(resp => resp.Comments));
  }
}
