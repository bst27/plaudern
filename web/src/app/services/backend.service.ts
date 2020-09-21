import { Injectable } from '@angular/core';
import {Observable, of} from 'rxjs';
import {ApproveCommentResponse, PutCommentResponseInterface} from '../models/approve-comment-response';
import {HttpClient} from '@angular/common/http';
import {Comment} from '../models/comment';
import {map} from 'rxjs/operators';
import {AuthResponse} from '../models/auth-response';

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

  login(password: string): Observable<boolean> {
    const formData = new FormData();
    formData.set('password', password);

    return this.http.post<AuthResponse>('/manage/login', formData).pipe(map(resp => resp.Authorized)); // TODO: Handle failure
  }

  checkAuth(): Observable<boolean> {
    return this.http.get<AuthResponse>('/manage/auth').pipe(map(resp => resp.Authorized)); // TODO: Handle failure
  }

  logout(): Observable<boolean> {
    return this.http.post<AuthResponse>('/manage/logout', undefined).pipe(map(resp => resp.Authorized)); // TODO: Handle failure
  }
}
