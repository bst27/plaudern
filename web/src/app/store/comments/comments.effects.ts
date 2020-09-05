import {Injectable} from "@angular/core";
import {Actions, createEffect, ofType} from "@ngrx/effects";
import {CommentsService} from "../../services/comments.service";
import {commentsLoadedSuccess, loadComments} from "./comments.actions";
import {map, mergeMap} from "rxjs/operators";


@Injectable()
export class CommentsEffects {

    loadComments$ = createEffect(() => this.actions$.pipe(
      ofType(loadComments),
      mergeMap(() => this.commentsService.getComments().pipe(
        map(comments => commentsLoadedSuccess({ comments: comments }))
        )
      )
    ))

  constructor(
    private actions$: Actions,
    private commentsService: CommentsService
  ) {
  }
}
