import {Comment} from './comment';

export interface PutCommentResponseInterface {
  Comment: Comment;
}

export class ApproveCommentResponse implements PutCommentResponseInterface {
  Comment: Comment;
}
