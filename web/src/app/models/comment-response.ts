import { Comment } from "./comment";

// This interface describes a comment response as received from the backend
export interface CommentResponse {
  Comments: Comment[];
}
