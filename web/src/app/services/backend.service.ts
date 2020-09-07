import { Injectable } from '@angular/core';
import {Observable, of} from "rxjs";
import {ApproveCommentResponse, PutCommentResponseInterface} from "../models/approve-comment-response";
import {HttpClient} from "@angular/common/http";
import {Comment} from "../models/comment";

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor(
    private http: HttpClient
  ) { }

  approveComment(comment: Comment): Observable<PutCommentResponseInterface> {
    const formData = new FormData();
    formData.set('status', 'published');

    return this.http.put<PutCommentResponseInterface>(
      '/manage/comment/' + comment.Id,
      formData
    );
  }

  revokeComment(comment: Comment): Observable<PutCommentResponseInterface> {
    const formData = new FormData();
    formData.set('status', 'created');

    return this.http.put<PutCommentResponseInterface>(
      '/manage/comment/' + comment.Id,
      formData
    );
  }
}
