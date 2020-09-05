import { createReducer, on } from "@ngrx/store";
import {commentsLoadedSuccess, primaryButtonClicked, secondaryButtonClicked} from "./comments.actions";
import {Comment} from "../../models/comment";

export type CommentsState = {
  comments: Comment[],
}

export const initState: CommentsState = {
  comments: [],
};

const reducer = createReducer(
  initState,
  on(primaryButtonClicked, (state, props) => {
    console.log('Primary button clicked for ' + props.comment.Id);
    return state;
  }),
  on(secondaryButtonClicked,  (state, props) => {
    console.log('Secondary button clicked for ' + props.comment.Id);
    return state;
  }),
  on(commentsLoadedSuccess, (state, props) => {
    return {...state, comments: props.comments};
  })
);

export function commentsReducer(state, action) {
  return reducer(state, action);
}
