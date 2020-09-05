import {ActionReducer} from "@ngrx/store";

export function logger(reducer: ActionReducer<any>): ActionReducer<any> {
  return function(state, action) {
    console.log(action.type, { action: action, state: state});
    return reducer(state, action);
  };
}
