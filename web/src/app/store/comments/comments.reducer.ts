import { createReducer, on } from '@ngrx/store';
import {
  commentApproved,
  commentRevoked,
  commentsLoadedSuccess,
  primaryButtonClicked,
  secondaryButtonClicked
} from './comments.actions';
import {Comment} from '../../models/comment';

export type CommentsState = {
  comments: Comment[],
};

export const initState: CommentsState = {
  comments: [],
};

const reducer = createReducer(
  initState,
  on(primaryButtonClicked, (state, props) => {
    return state;
  }),
  on(secondaryButtonClicked,  (state, props) => {
    return state;
  }),
  on(commentsLoadedSuccess, (state, props) => {
    return {...state, comments: props.comments};
  }),
  on(commentApproved, commentRevoked, (state, props) => {
    return {...state, comments: state.comments.map(comment => {
      if (comment.Id === props.comment.Id) {
        comment = props.comment;
      }

      return comment;
    })};
  }),
);

export function commentsReducer(state, action) {
  return reducer(state, action);
}
