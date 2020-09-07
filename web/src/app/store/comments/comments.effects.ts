import {Injectable} from "@angular/core";
import {Actions, createEffect, ofType} from "@ngrx/effects";
import {CommentsService} from "../../services/comments.service";
import {
  commentApproved,
  commentRevoked,
  commentsLoadedSuccess, loadComment,
  loadComments,
  primaryButtonClicked
} from "./comments.actions";
import {map, mergeMap} from "rxjs/operators";
import {BackendService} from "../../services/backend.service";


@Injectable()
export class CommentsEffects {

    loadComments$ = createEffect(() => this.actions$.pipe(
      ofType(loadComments),
      mergeMap(() => this.commentsService.getComments().pipe(
        map(comments => commentsLoadedSuccess({ comments: comments }))
        )
      )
    ))

    loadComment$ = createEffect(() => this.actions$.pipe(
      ofType(loadComment),
      mergeMap((action) => {//TODO: Improve by checking if comment has already been loaded
        return this.commentsService.getComments().pipe(
          map(comments => commentsLoadedSuccess({ comments: comments }))
        );
      }
      )
    ))

    primaryButtonClicked$ = createEffect(() => this.actions$.pipe(
      ofType(primaryButtonClicked),
      mergeMap((action) => {
        if (action.comment.Status === 'created') {//TODO: Move this decisions into component and dispatch corresponding actions
          return this.backend.approveComment(action.comment).pipe(
            map(result => commentApproved({ comment: result.Comment }))//TODO: Handle errors
          );
        } else if (action.comment.Status === 'published') {
          return this.backend.revokeComment(action.comment).pipe(
            map(result => commentRevoked({ comment: result.Comment }))//TODO: Handle errors
          );
        }

      })
    ))

  constructor(
    private actions$: Actions,
    private commentsService: CommentsService,
    private backend: BackendService,
  ) {
  }
}
