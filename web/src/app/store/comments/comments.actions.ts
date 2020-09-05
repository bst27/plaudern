import {createAction, props} from "@ngrx/store";
import {Comment} from "../../models/comment";

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

export const commentsLoadedSuccess = createAction(
  '[Comment API] Comments Loaded Success',
  props<{ comments: Comment[] }>()
);

