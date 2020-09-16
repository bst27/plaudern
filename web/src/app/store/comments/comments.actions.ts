import {createAction, props} from '@ngrx/store';
import {Comment} from '../../models/comment';

export const primaryButtonClicked = createAction(
  '[Comment Card Component] Primary Button Clicked',
  props<{ comment: Comment}>()
);

export const secondaryButtonClicked = createAction(
  '[Comment Card Component] Secondary Button Clicked',
  props<{ comment: Comment}>(),
);

export const loadComments = createAction(
  '[Comment List] Load Comments'
);

export const loadComment = createAction(
  '[Comment Details] Load Comment',
  props<{ commentId: string }>()
);

export const commentsLoadedSuccess = createAction(
  '[Comment API] Comments Loaded Success',
  props<{ comments: Comment[] }>()
);

export const commentApproved = createAction(
  '[Comment API] Comment Approved',
  props<{ comment: Comment }>()
);

export const commentRevoked = createAction(
  '[Comment API] Comment Revoked',
  props<{ comment: Comment }>()
);

